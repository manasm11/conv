package conv

import (
	"fmt"
)

func ToInt(str string) (o int) {
	_, err := fmt.Sscanf(str, "%d", &o)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert string to int: %v\n", err))
	}
	return o
}

func ToFloat(str string) (o float64) {
	_, err := fmt.Sscanf(str, "%f", &o)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert string to int: %v\n", err))
	}
	return o
}

func ToBool(str string) (o bool) {
	_, err := fmt.Sscanf(str, "%t", &o)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert string to int: %v\n", err))
	}
	return o
}
func ToString(v any) string {
	return fmt.Sprintf("%+#v", v)
}
