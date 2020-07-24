package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Print(ToUpperCase("daeneris", 2))
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(FilterRange(slice, 3, 7))
	name:= []string{"javaScript", "Go", "React"}
	fmt.Print(RemoveItemInArray(name, "a"))

}
// Return same String with upperCase
func ToUpperCase(str string, index int) (string, error) {
	if index > len(str)-1 {
		return "", errors.New("Invalid index")
	}
	word := ""
	for i, value := range str {
		if i == index {
			word += strings.ToUpper(string(value))
			continue
		}
		word += string(value)
	}
	return word, nil
}
// Return Range of Slice
func FilterRange(array []int, min, max int)[]int {
	_range := []int{}
	for _, value := range array {
		if value <= max && value >= min {
			_range = append(_range, value)
		}
	}
	return _range
}
// custom include function for RemoveItemInArray function
func customInclude(item, filter string) bool {
	for _, value:= range item {
		if string(value) == filter {
			return true
		}
	}
	return false
}



// Return filtered array
func RemoveItemInArray(array []string, filter string) []string{
	filtered:= []string{}
	for _, value:= range array {
		condidate:= customInclude(strings.ToLower(value),strings.ToLower(filter))
		if condidate == true {
			continue
		}
		filtered = append(filtered, value)
	}
	return filtered
}
