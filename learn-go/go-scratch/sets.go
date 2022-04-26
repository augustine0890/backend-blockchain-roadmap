package main

import (
	"fmt"
	"reflect"
)

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(col))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func removeDupInt(a []int) []int {
	allEle := make(map[int]bool)
	list := []int{}
	for _, v := range a {
		if _, ok := allEle[v]; !ok {
			allEle[v] = true
			list = append(list, v)
		}
	}
	return list
}

func main() {
	data := []struct{ in, out []string }{
		{[]string{}, []string{}},
		{[]string{"", "", ""}, []string{""}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "b", "b", "a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "c", "a", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, exp := range data {
		res := UniqStr(exp.in)
		if !reflect.DeepEqual(res, exp.out) { // if inputs don't match
			fmt.Printf("%q didn't match %q\n", res, exp.out)
		}
	}

	studentIDs := []int{11, 33, 22, 1, 33, 12, 25, 11}
	fmt.Printf("Student IDs: %v\n", removeDupInt(studentIDs))
}
