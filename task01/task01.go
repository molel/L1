package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greet() {
	fmt.Printf("Hello! My name is %s and I am %d years old.\n", h.Name, h.Age)
}

// встраивание структуры (struct embedding)
type Action struct {
	Human
}

func main() {
	act := Action{Human{
		Name: "John",
		Age:  21,
	}}

	act.Greet()
}
