package utils

import "fmt"

func CheckNilError(err error) {
	if err != nil {
		fmt.Printf("error due to: %v\n", err)
	}
}

func CheckObjectIsNil(obj any, errMessage string) {
	if obj == nil {
		fmt.Printf("error due to: %v\n", errMessage)
	}
}
