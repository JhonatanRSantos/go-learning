package main

import (
	"fmt"

	"github.com/JhonatanRSantos/go-learning/V3/interfaces"
)

type Person interfaces.Person

func main() {
	contact := interfaces.Contact{}

	personOne := Person{
		FirstName: " Nolan",
		LastName:  "Grayson",
		Contact:   contact,
	}

	personTwo := Person{
		"Deborah", "Grayson", 40, contact, []string{"Ones"},
	}

	personThree := Person{}
	personThree.FirstName = "Mark"
	personThree.LastName = "Grayson"

	fmt.Println(personOne)
	fmt.Println(personTwo)
	fmt.Printf("%+v\n", personThree)

	personFour := interfaces.NewPerson()
	personFour.UpdateFirstName("King", []string{"Ones"})
	personFour.Print()
}
