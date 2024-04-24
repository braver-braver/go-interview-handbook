package main

import (
	"errors"
	"fmt"
)

func f1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("defer1 error")
	return
}

func f2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	return
}

func f3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 error")
	return
}

func main() {
	f1()
	f2()
	f3()
}
