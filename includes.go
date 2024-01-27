package esc

import (
	"reflect"
)

func Includes(slice interface{}, element interface{}) bool {
	sliceValue := reflect.ValueOf(slice)

	for i := 0; i < sliceValue.Len(); i++ {
		if reflect.DeepEqual(sliceValue.Index(i).Interface(), element) {
			return true
		}
	}

	return false
}
