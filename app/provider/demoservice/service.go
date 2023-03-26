package demoservice

import "fmt"

type DemoService struct {
}

func NewDemo(params ...interface{}) (interface{}, error) {
	return &DemoService{}, nil
}

func (demo *DemoService) SayHello() {
	fmt.Println("Hello !!")
}

func (demo *DemoService) SayByeBye() {
	fmt.Println("Bye Bye !!")
}
