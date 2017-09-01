package main

import (
	"fmt"
)

//inheritance in golang
type Person struct {
	name string
	age  int
}
type Author struct {
	person         Person
	publishedBooks []string
}

func main() {
	p := Person{"Riya", 35}
	fmt.Println(p.ingrementAge())
	fmt.Println(p.age)
	me := Author{
		person: Person{
			name: "Riya",
			age:  37},
		publishedBooks: make([]string, 1),
	}
	fmt.Println(me.person.age)
}

func (p *Person) ingrementAge() (int) {
	p.age++
	return p.age;
}
