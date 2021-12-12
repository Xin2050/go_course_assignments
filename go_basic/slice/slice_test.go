package slice

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Printf("%v\n", s)
	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Printf("%v\n", s)
	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Printf("%v\n", s)
	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Printf("%v\n", s)
	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Printf("%v\n", s)
	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Printf("%v\n", s)
}

func Add(s []int, index int, value int) []int {

	var newSlice []int
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, value)
	return append(newSlice, s[index:]...)

}

func Delete(s []int, index int) []int {
	var newSlice []int
	newSlice = append(newSlice, s[:index]...)
	return append(newSlice, s[index+1:]...)

}
func TestSliceShare(t *testing.T) {
	months := []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	Q2 := months[3:6]

	t.Log(Q2, len(Q2), cap(Q2))
	months[3] = "April"
	t.Log(Q2, len(Q2), cap(Q2))

}
