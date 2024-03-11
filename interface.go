package main

// interface
type Animal interface{
	Speak() string
}

type Dog struct{
	// do somthing here
	bark string
}

func(d Dog) Speak() string{
	return d.bark
}

type Cat struct{
	// do somthing here
	meow string
}

func(c Cat) Speak() string{
	return c.meow;
}