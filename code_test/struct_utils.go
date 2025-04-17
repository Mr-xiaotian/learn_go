package main

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
