package array

import (
	"fmt"
	"reflect"
)

// Chunk - Creates an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements.
func Chunk(arr interface{}, size int) ([]interface{}, error) {
	var tempArr []interface{}
	var arrVal = reflect.ValueOf(arr)
	var arrKind = arrVal.Kind()
	if arrKind != reflect.Slice && arrKind != reflect.Array {
		return tempArr, fmt.Errorf("The argument arr is not an array or slice.")
	}
	if size < 1 {
		return tempArr, fmt.Errorf("The argument size must grather")
	}
	var arrLength = arrVal.Len()
	if size >= arrLength {
		size = arrLength
	}

	var result []interface{}
	for i := 0; i < arrLength; i++ {
		tempArr = append(tempArr, arrVal.Index(i).Interface())
		if len(tempArr) == size {
			result = append(result, tempArr)
			tempArr = tempArr[:0]
		}
	}
	if len(tempArr) > 0 {
		result = append(result, tempArr)
	}
	return result, nil
}
