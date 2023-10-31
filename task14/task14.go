package main

import (
	"fmt"
	"reflect"
)

func Type(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	default:
		// с помощью библиотеки reflect отбираеются все каналы, а не каналы конкретного типа
		switch reflect.ValueOf(a).Kind() {
		case reflect.Chan:
			fmt.Println("Type is chan")
		default:
			fmt.Println("Type cannot be recognised")
		}
	}
}

func main() {
	a := 0
	Type(a)

	b := "0"
	Type(b)

	c := false
	Type(c)

	d := make(chan int)
	Type(d)

	e := struct{}{}
	Type(e)
}
