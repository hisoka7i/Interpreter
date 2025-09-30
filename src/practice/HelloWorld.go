package main

import "fmt"

func main(){
	fmt.Println("Hello, world")
	var inter interface{} = 23  //this is called empty interface, and it can save any value it is like a object.
	fmt.Printf("%#v\n", inter)
}