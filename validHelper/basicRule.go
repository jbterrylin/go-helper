package validhelper

import "fmt"

func NotEmpty() string {
	return "notEmpty"
}

func RegexpMatch(rule string) string {
	return fmt.Sprintf("%s=%s", REGEXP, rule)
}

func Lt(mark string) string {
	return fmt.Sprintf("%s=%s", LT, mark)
}

func Le(mark string) string {
	return fmt.Sprintf("%s=%s", LE, mark)
}

func Eq(mark string) string {
	return fmt.Sprintf("%s=%s", EQ, mark)
}

func Ne(mark string) string {
	return fmt.Sprintf("%s=%s", NE, mark)
}

func Ge(mark string) string {
	return fmt.Sprintf("%s=%s", GE, mark)
}

func Gt(mark string) string {
	return fmt.Sprintf("%s=%s", GT, mark)
}
