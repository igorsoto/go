package main

import "testing"

func TestPerson(t *testing.T) {
	p := person{
		name:"Person1",
	}

	p.setNameByValue("Person2")
	if p.name != "Person1" {
		t.Error("When passing by value expect original reference to not change")
	}

	p.setNameByReference("Person3")
	if p.name != "Person3" {
		t.Error("When passing by reference expect original reference to change")
	}
}