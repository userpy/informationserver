package main

import "fmt"

func TestPanic(varl int) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("function recover")
			result = ""
			err = fmt.Errorf("Context valuet not found")
			return
		}
	}()
	panic("AAAAAA!")
}

func main() {
	TestPanic(1)

}
