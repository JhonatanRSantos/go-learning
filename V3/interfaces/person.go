package interfaces

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Contact
	Values []string
}

type Contact struct {
	Email   string
	ZipCode string
}

func NewPerson() Person {
	return Person{}
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) UpdateFirstName(firstName string, values []string) {
	p.FirstName = firstName
	p.Values = append(p.Values, values...)
}
