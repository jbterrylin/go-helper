package validhelper_test

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	formathelper "github.com/jbterrylin/go-helper/formatHelper"
	pointerhelper "github.com/jbterrylin/go-helper/pointerHelper"
	validhelper "github.com/jbterrylin/go-helper/validHelper"
)

func TestNotEmpty(t *testing.T) {
	result := validhelper.NotEmpty()
	expected := "notEmpty"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRegexpMatch(t *testing.T) {
	rule := `^\d+$`
	result := validhelper.RegexpMatch(rule)
	expected := fmt.Sprintf("%s=%s", validhelper.REGEXP, rule)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLt(t *testing.T) {
	mark := "10"
	result := validhelper.Lt(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.LT, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestLe(t *testing.T) {
	mark := "10"
	result := validhelper.Le(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.LE, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEq(t *testing.T) {
	mark := "10"
	result := validhelper.Eq(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.EQ, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNe(t *testing.T) {
	mark := "10"
	result := validhelper.Ne(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.NE, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGe(t *testing.T) {
	mark := "10"
	result := validhelper.Ge(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.GE, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGt(t *testing.T) {
	mark := "10"
	result := validhelper.Gt(mark)
	expected := fmt.Sprintf("%s=%s", validhelper.GT, mark)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestValidatorVerify(t *testing.T) {
	type User struct {
		Age           int
		Name          string
		Email         string
		HeightPointer *float64
		Height        float64
		TimePointer   *time.Time
		Time          time.Time
	}

	rules := validhelper.Rules{
		"Age":           {validhelper.NotEmpty(), validhelper.Gt("18")},
		"Name":          {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[A-Za-z]+$`)},
		"Email":         {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)},
		"HeightPointer": {validhelper.Ge("150")},
		"Height":        {validhelper.NotEmpty(), validhelper.Ge("150")},
		"TimePointer":   {validhelper.NotEmpty(), validhelper.Ge("15000")},
		"Time":          {validhelper.NotEmpty(), validhelper.Ge("15000")},
	}

	user := User{
		Age:           20,
		Name:          "John",
		Email:         "john.doe@example.com",
		HeightPointer: pointerhelper.Pointer(180.5),
		Height:        180.5,
		TimePointer:   pointerhelper.Pointer(time.Now()),
		Time:          time.Now(),
	}

	validator := validhelper.NewValidator(nil)
	err := validator.Verify(user, rules)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	user.Age = 17
	err = validator.Verify(user, rules)
	if err == nil || err.Error() != "Age len is not valid" {
		t.Errorf("Expected Age len is not valid error, got %v", err)
	}
}

func TestValidatorVerifyNestedStruct(t *testing.T) {
	type Address struct {
		City  string
		State string
	}

	type User struct {
		Age     int
		Name    string
		Email   string
		Address *Address
	}

	rules := validhelper.Rules{
		"Age":           {validhelper.NotEmpty(), validhelper.Gt("18")},
		"Name":          {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[A-Za-z]+$`)},
		"Email":         {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)},
		"Address.City":  {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[A-Za-z]+$`)},
		"Address.State": {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[A-Za-z]+$`)},
	}

	address := &Address{
		City:  "NewYork",
		State: "NY",
	}

	user := User{
		Age:     20,
		Name:    "John",
		Email:   "john.doe@example.com",
		Address: address,
	}

	validator := validhelper.NewValidator(nil)
	err := validator.Verify(user, rules)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	user.Address.City = ""
	err = validator.Verify(user, rules)
	if err == nil || err.Error() != "Address.City is empty" {
		t.Errorf("Expected Address.City is empty error, got %v", err)
	}
}

func TestValidatorVerifyWithCustomRule(t *testing.T) {
	type User struct {
		Age           int
		Name          string
		Email         string
		HeightPointer *float64
		Height        float64
		TimePointer   *time.Time
		Time          time.Time
		SavingAccount float64
		Force2Field   float64
	}

	// custom rule start
	// encourage to make error and custom rule key to variable
	ErrWrongDecimal := errors.New("wrong decimal")
	const FloatForceDecimalConst = "FLOAT_FORCE_DECIMAL"
	// function use for rules
	floatForceDecimal := func(mark int) string {
		return fmt.Sprintf("%s=%d", FloatForceDecimalConst, mark)
	}
	// may copy defaultT and all new for custom error
	newErrTranslate := func(fieldName string, _ reflect.Value, err error, rule string) error {
		switch err {
		case validhelper.ErrIsEmpty:
			return fmt.Errorf("%s %s", fieldName, err.Error())
		case validhelper.ErrRegexp:
			return fmt.Errorf("%s %s", fieldName, err.Error())
		case validhelper.ErrLenNotValid:
			return fmt.Errorf("%s %s", fieldName, err.Error())
		case ErrWrongDecimal:
			return fmt.Errorf("%s %s", fieldName, err.Error())
		}
		return nil
	}
	customRule := map[string]func(st interface{}, targetField reflect.Value, rule string) (err error){
		FloatForceDecimalConst: func(st interface{}, targetField reflect.Value, rule string) (err error) {
			value := targetField.Float()
			decimal, _ := strconv.ParseInt(strings.Split(rule, "=")[1], 10, 64)
			if formathelper.TruncateFloat(value, int(decimal)) != value {
				return ErrWrongDecimal
			}
			return nil
		},
	}
	// custom rule end

	rules := validhelper.Rules{
		"Age":           {validhelper.NotEmpty(), validhelper.Gt("18")},
		"Name":          {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[A-Za-z]+$`)},
		"Email":         {validhelper.NotEmpty(), validhelper.RegexpMatch(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)},
		"HeightPointer": {validhelper.NotEmpty(), validhelper.Ge("150")},
		"Height":        {validhelper.NotEmpty(), validhelper.Ge("150")},
		"TimePointer":   {validhelper.NotEmpty(), validhelper.Ge("15000")},
		"Time":          {validhelper.NotEmpty(), validhelper.Ge("15000")},
		"Force2Field":   {floatForceDecimal(2)},
	}

	user := User{
		Age:           20,
		Name:          "John",
		Email:         "john.doe@example.com",
		HeightPointer: pointerhelper.Pointer(180.5),
		Height:        180.5,
		TimePointer:   pointerhelper.Pointer(time.Now()),
		Time:          time.Now(),
		Force2Field:   180.14,
	}

	validator := validhelper.NewValidator(newErrTranslate)
	validator.AddCustomRule(customRule)

	err := validator.Verify(user, rules)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
