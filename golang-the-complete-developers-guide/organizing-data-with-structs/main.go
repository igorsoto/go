package main

import "fmt"

// value type
type contactInfo struct {
	email string
}

// value type
type person struct {
	name string
	contactInfo
}

func (p person) setNameByValue(name string) person {
	p.name = name
	return p
}

func (p *person) setNameByReference(name string) {
	(*p).name = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	var p = person {
		name:"Igor",
		contactInfo: contactInfo{email: "user@test.com"},
	}
	p.print()

	copy:= p.setNameByValue("Marie")
	p.print()
	copy.print()

	p.setNameByReference("Marie")
	p.print()

}
