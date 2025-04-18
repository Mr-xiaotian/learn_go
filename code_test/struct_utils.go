package main

import (
	"fmt"
	"math"
)

type Address struct {
	City string
}

type User struct {
	Name string
	Addr *Address
}

func cloneUser(u *User, newCity string) *User {
	return &User{
		Name: u.Name,
		Addr: &Address{City: newCity},
	}
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) HaveBirthday() {
	p.Age += 1
}

type Student Person

func AverageAge(students []Student) float64 {
	var totalAge int
	for _, student := range students {
		totalAge += student.Age
	}
	return float64(totalAge) / float64(len(students))
}

type Author struct {
	Name string
}

type Book struct {
	Title      string
	BookAuthor Author
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("《%s》 by %s", b.Title, b.BookAuthor.Name)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		r.Width * factor,
		r.Height * factor,
	}
}
