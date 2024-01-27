package js

import (
	"reflect"
)

// Includes checks if an element is present in a slice.
// It uses deep equality comparison for element matching.
func Includes(slice interface{}, element interface{}) bool {
    sliceValue := reflect.ValueOf(slice)

    for i := 0; i < sliceValue.Len(); i++ {
        if reflect.DeepEqual(sliceValue.Index(i).Interface(), element) {
            return true
        }
    }

    return false
}

// Filter takes a slice and a callback function, and returns a new slice containing
// only the elements for which the callback function returns true.
func Filter(words []string, callback func(string) bool) []string {
    var result []string

    for _, word := range words {
        if callback(word) {
            result = append(result, word)
        }
    }

    return result
}

