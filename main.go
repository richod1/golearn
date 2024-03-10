package main

import (
	"fmt"
)

// import ("rsc.io/quote/v4")

const A int = 10

// struct in go lang
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// interface in go
type Shape interface{
	Area() float32;
}
type Rectangle struct{
	Length float32;
	Width float32;

}
func(r Rectangle) Area() float32{
	return r.Length*r.Width;
}

type Circle struct{
	Radius float32;
}

func(c Circle) Area() float32{
	return 3.14*c.Radius*c.Radius;
}

func main() {
	var myName string = "Degraft Frimpong"
	fmt.Println(myName)
	myLastName := "frimpong"
	fmt.Println("my last name is :", myLastName)

	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println(a, b, c, d)

	// block format in variable decleraton
	var (
		s int
		k int    = 20
		e string = "hello"
	)
	fmt.Println(s, k, e)

	var _firstname = "Degraft"
	_firstname = "kweku"
	fmt.Println(_firstname)
	const PI = 3.14
	fmt.Println(PI)
	fmt.Println(A)

	// arrays
	var myArray = [...]int{2, 4, 6, 8, 10}
	myArray[0] = 40
	fmt.Println(myArray[0])

	var initArray = [4]string{"mazda", "offroad", "hintz", "cargo"}
	fmt.Println(len(initArray))

	// slice
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(mySlice))

	// making slice out of array
	var arr = [6]int{2, 4, 6, 8, 9, 0}
	myslice := arr[2:4]
	fmt.Println(myslice)

	myslice1 := []int{3, 6, 9, 14, 15, 17, 19}
	myslice2 := []int{2, 4, 6, 8, 10}
	myslice1 = append(myslice1, 20, 21, 22)

	fmt.Println(myslice1)

	myslice3 := append(myslice1, myslice2...)
	fmt.Println(myslice3)

	// conditions
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	y := 20
	x := 40
	if x > y {
		fmt.Println("x is greater")
	}

	time := 20
	if time < 20 {
		fmt.Println("Good Day")
	} else {
		fmt.Println("Good evening")
	}

	// switch statements
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("today is monday")
	case "Tuesday":
		fmt.Println("today is tuesday")
	case "Wednesday":
		fmt.Println("today is wednesday")
	default:
		fmt.Println("Not a week day")
	}

	// for loop in go
	for i := 0; i < 20; i++ {
		if i == 15 {
			break
		}
		fmt.Println(i)
	}

	// nested forloop in go
	adj := [2]string{"big", "teasty"}
	fruit := [3]string{"apple", "orange", "banana"}

	for index := 0; index < len(adj); index++ {
		for j := 0; j < len(fruit); j++ {
			fmt.Println(adj[index], fruit[j])
		}
	}

	// using range in iteration

	fruits := []string{"Apple", "mongo", "lemon", "pawpaw"}
	for idx, vlu := range fruits {
		fmt.Println(idx, vlu)
	}

	// to show only value omitting the index
	for _, vlu := range fruits {
		fmt.Println(vlu)
	}
	// function calling
	message()
	detailName("degraft", "frimpong", "22")
	// usage of struct
	var person1 Person

	person1.name = "Degraft frimpong"
	person1.age = 22
	fmt.Println("Name ", person1.name)
	fmt.Println("Age ", person1.age)


	// maps in go similar to objects in js
	var cars=map[string]string{"brand":"Ford","model":"Mustaung","year":"1964"};
	fmt.Println("Brand ",cars["brand"])


	// usage of interface and struct
	var shape1 Shape;
	var shape2 Shape;

	rect:=Rectangle{Length:5,Width: 3};
	circle:=Circle{Radius: 2};

	shape1=rect;
	shape2=circle;

	fmt.Println("Area of shape1 {Rectangle} :",shape1.Area());
	fmt.Println("Area of shape2 {Circle} :",shape2.Area())


}
func message() {
	fmt.Println("hello from msg function")
}

// function with parameters
func detailName(fname string, lname string, age string) {
	fmt.Println("my firstname is " + fname + " and my lastname is " + lname + " and i am " + age + " years old")
}
