package main

import "fmt"

type Person struct {
	Name      string
	NickNames []string
}

func main() {
	greeting := [5]string{"Hi", "there.", "How", "are", "you?"}
	// fmt.Printf("%p\n", &greeting)
	update(greeting)
	fmt.Println(greeting)

	person := Person{Name: "Jhonatan", NickNames: []string{"Banana", "Potatoes", "Chips", "Pizza", "LOL"}}
	// fmt.Printf("%p\n", &person)
	changeName(person)
	fmt.Println(person)

	fmt.Println(len(person.NickNames))
	fmt.Println(cap(person.NickNames))
}

func update(greeting [5]string) {
	// fmt.Printf("%p\n", &greeting)
	greeting[0] = "Hello"
}

func changeName(person Person) {
	// fmt.Printf("%p\n", &person)
	person.Name = "Rodrigues"
}
