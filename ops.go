package main
import("fmt")

type Persons struct{
	Name string
	Age int
}

func(p *Persons) SayHello(){
	// p.Age=22;
	// p.Name="Degraft Frimpong";

	fmt.Printf("Hello ,my name is %s and I'am %d years old.\n",p.Name,p.Age);
}

