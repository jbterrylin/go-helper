package validhelper

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rules map[string][]string

type Validator struct {
	t                  func(fieldName string, field reflect.Value, err error, rule string) error
	customRuleExistMap map[string]bool
	customRule         map[string]func(st interface{}, targetField reflect.Value, rule string) (err error)
}

func NewValidator(t func(fieldName string, field reflect.Value, err error, rule string) error) Validator {
	if t == nil {
		t = defaultT
	}
	return Validator{
		t:                  t,
		customRuleExistMap: map[string]bool{},
		customRule:         map[string]func(st interface{}, targetField reflect.Value, rule string) (err error){},
	}
}

// customRule key format must be <NAME>=<VALUE>
// e.x: IS_FLOAT_FORCE_DECIMAL=2
func (a *Validator) AddCustomRule(customRule map[string]func(st interface{}, targetField reflect.Value, rule string) (err error)) Validator {
	for key, val := range customRule {
		a.customRuleExistMap[CUSTOM_RULE+key] = true
		a.customRule[CUSTOM_RULE+key] = val
	}
	return *a
}

func (a *Validator) RemoveCustomRule(customRules []string) Validator {
	if len(customRules) == 0 {
		a.customRuleExistMap = map[string]bool{}
		a.customRule = map[string]func(st interface{}, targetField reflect.Value, rule string) (err error){}
		return *a
	}
	for _, customRule := range customRules {
		delete(a.customRuleExistMap, CUSTOM_RULE+customRule)
		delete(a.customRule, CUSTOM_RULE+customRule)
	}
	return *a
}

func defaultT(fieldName string, _ reflect.Value, err error, rule string) error {
	switch err {
	case ErrIsEmpty:
		return fmt.Errorf("%s %s", fieldName, err.Error())
	case ErrRegexp:
		return fmt.Errorf("%s %s", fieldName, err.Error())
	case ErrLenNotValid:
		return fmt.Errorf("%s %s", fieldName, err.Error())
	}
	return nil
}

func (a *Validator) verifyField(st interface{}, targetField reflect.Value, rules []string, fieldName string) (err error) {
	for _, v := range rules {
		switch {
		case v == NOT_EMPTY:
			if isBlank(targetField) {
				return a.t(fieldName, targetField, ErrIsEmpty, v)
			}
		case strings.Split(v, "=")[0] == REGEXP:
			switch targetField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if !regexpMatch(strings.Split(v, "=")[1], fmt.Sprint(targetField.Interface())) {
					return a.t(fieldName, targetField, ErrRegexp, v)
				}
			default:
				if !regexpMatch(strings.Split(v, "=")[1], targetField.String()) {
					return a.t(fieldName, targetField, ErrRegexp, v)
				}
			}
		case a.customRuleExistMap[CUSTOM_RULE+strings.Split(v, "=")[0]]:
			err := a.customRule[CUSTOM_RULE+strings.Split(v, "=")[0]](st, targetField, v)
			if err != nil {
				return a.t(fieldName, targetField, err, v)
			}
		case compareMap[strings.Split(v, "=")[0]]:
			if targetField.Kind() == reflect.Pointer {
				if targetField.IsNil() {
					return a.t(fieldName, targetField, ErrIsEmpty, v)
				}
				targetField = targetField.Elem()
			}
			if !compareVerify(targetField, v) {
				return a.t(fieldName, targetField, ErrLenNotValid, v)
			}
		}
	}
	return
}

func combineFieldName(parentField, currentField string) string {
	if parentField == "" {
		return currentField
	}
	return parentField + "." + currentField
}

func (a *Validator) Verify(st interface{}, ruleMap Rules) (err error) {
	return a.verify(st, ruleMap, "")
}

func (a *Validator) verify(st interface{}, ruleMap Rules, parentField string) (err error) {
	val := reflect.ValueOf(st)

	if val.Kind() != reflect.Struct {
		return errors.New("expect struct")
	}

	for key, values := range ruleMap {
		if len(values) == 0 {
			continue
		}
		keys := strings.Split(key, ".")
		if len(keys) > 1 {
			derivedStruct := val.FieldByName(keys[0])
			if derivedStruct.Kind() == reflect.Pointer {
				derivedStruct = derivedStruct.Elem()
			}
			if !derivedStruct.IsValid() {
				return a.t(combineFieldName(parentField, keys[0]), derivedStruct, ErrIsEmpty, "")
			}
			err = a.verify(derivedStruct.Interface(), Rules{
				strings.Join(keys[1:], "."): values,
			}, combineFieldName(parentField, keys[0]))
			if err != nil {
				return err
			}
			continue
		}
		err = a.verifyField(st, val.FieldByName(keys[0]), values, combineFieldName(parentField, keys[0]))
		if err != nil {
			return
		}
	}
	return nil
}

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String:
		return compare(len([]rune(value.String())), VerifyStr)
	case reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		// will compare as Int
		if isTimeType(value) {
			return compare(timeToUnixTime(value), VerifyStr)
		} else {
			return false
		}
	}
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func compareHelper[T Comparable](val T, op string, target T) bool {
	switch {
	case op == LT:
		return val < target
	case op == LE:
		return val <= target
	case op == EQ:
		return val == target
	case op == NE:
		return val != target
	case op == GE:
		return val >= target
	case op == GT:
		return val > target
	default:
		return false
	}
}

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		return compareHelper(val.Int(), VerifyStrArr[0], VInt)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		return compareHelper(val.Uint(), VerifyStrArr[0], uint64(VInt))
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		return compareHelper(val.Float(), VerifyStrArr[0], VFloat)
	default:
		return false
	}
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
