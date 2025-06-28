package main

import "fmt"                 //import fmt -> #include<iostream>  in cpp

func main(){
	fmt.Println("hello");
	fmt.Println('h');     // ->single character
	var productName string= "hiiii"


	pName := "hiiii"      // -> shorthand
	
	fmt.Print(pName)     
	fmt.Println(productName)
	
	
	
	//var val1 bool      // ->false
	//var val2 string      // ->null
	//var newVal int      // ->0
	//fmt.Println((newVal))
	loops()
		
}


func loops(){
	for i:=0;i<5;i++{
		fmt.Print(i);
	}		
	
	for j := range 3{
		fmt.Print("Range",j)
	}

	for i,char:= range "Trishit"{
		fmt.Println(i,char);
	}

	j:=0
	for j<3{
		fmt.Print(j)
		j++
	}

	for{
		fmt.Print("infinite loop")
	}
}