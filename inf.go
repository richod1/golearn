package main

import "fmt"


type Human struct{
	name string
	age int
	phone string
}


type Student struct{
	Human
	school string
	loan float32
}

type Employee struct{
	Human
	company string
	money float32
}

type Men interface{
	SayHi()
	Sing(lyrics string)
}

func(h Human)SayHi(){
	fmt.Printf("Hi i am %s you can call me on %s\n",h.name,h.phone)

}

func (e Employee) SayHi(){
	fmt.Printf("Hi i am %s, I work at %s. Call me on %s\n",e.name,e.company,e.phone)
}