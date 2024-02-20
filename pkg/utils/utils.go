package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/twinj/uuid"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorGray   = "\033[37m"
	ColorWhite  = "\033[97m"
)

func Colorize(colorCode, text string) string {
	if colorCode == "" || text == "" {
		return ""
	}
	return colorCode + text + ColorReset
}

func UUIDv4() string {
	return uuid.NewV4().String()
}

func GenHexStr(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func PrintStructFieldsAndValues(s interface{}, indent string) {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Printf("%s%v is not a struct\n", indent, v.Type())
		return
	}

	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("%s - %s: ", indent, typeOfS.Field(i).Name)

		if field.Kind() == reflect.Struct {
			fmt.Println()
			PrintStructFieldsAndValues(field.Interface(), indent+"  ")
		} else if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			fmt.Println()
			PrintStructFieldsAndValues(field.Interface(), indent+"  ")
		} else {
			if field.CanInterface() {
				fmt.Println(Colorize(ColorGreen, fmt.Sprint(field.Interface())))
			} else {
				fmt.Println()
			}
		}
	}
}
