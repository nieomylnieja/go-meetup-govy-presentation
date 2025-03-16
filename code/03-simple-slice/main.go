package main

import (
	"fmt"
	"slices"
	"strings"
)

type multiError []string // CO
// CO
func (m multiError) Error() string { // CO
	return strings.Join(m, "; ") // CO
} // CO
// CO
type Teacher struct {
	Name     string
	Age      int
	Students []Student
}

type Student struct {
	Index string
	Age   int
}

func validateTeacher(teacher Teacher) error {
	var mErr multiError
	if len(teacher.Name) > 20 {
		mErr = append(mErr, "Teacher's name is too long!")
	}
	if !slices.Contains([]string{"John", "Eve"}, teacher.Name) {
		mErr = append(mErr, "we can only allow John and Eve to become teachers")
	}
	if teacher.Age < 20 {
		mErr = append(mErr, "Teacher is too young...")
	}
	return mErr
}

func validateStudent(student Student) error {

}

func main() {
	teacher := Teacher{Name: "Frank", Age: 18}
	err := validateTeacher(teacher)
	fmt.Println("Error:", err)
}
