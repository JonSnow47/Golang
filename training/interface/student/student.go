/*
 *  Revision History:
 *      Initial: 2018/07/26    Chen Yan Chen
 */

package student

import (
	"fmt"
	"os/exec"
)

type Student struct {
	ID     int
	Name   string
	School string
}

var students map[int]Student

func init() {
	students = make(map[int]Student)
}

func New() *Student {
	id := len(students)
	cmd := exec.Command("openssl rand -hex 10")
	result, err := cmd.Output()
	fmt.Println(result)
	if err != nil {
		return &Student{
			ID:     id,
			Name:   string(result),
			School: "NCEPU",
		}
	}
	return nil
}

func (s *Student) GetName() string {
	return s.Name
}
func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) GetSchool() string {
	return s.School
}
