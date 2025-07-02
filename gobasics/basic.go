package main

import "fmt"                 //import fmt -> #include<iostream>  in cpp


type SellableProduct interface{
	buy() 
	getDiscount() int

}

func check_discountAndBuy(p SellableProduct){
	discount := p.getDiscount()
	if discount>30 {
		fmt.Println(("Discount is good,buying the product"))
		p.buy()
		return
	}  else{
		fmt.Println("Discount is not good,not buying the product")
		return
	}
}
type Product struct{
	name string
	price int              // ->struct is objects 
	company string
}

func newProduct(name string,price int,company string) *Product {  //taking product instance
	p:=Product{
		name:name,
		price: price,       // -> constructor of the product
		company: company,
	}
	//  return p        returning a copy of p -> creating double object          
	return &p   // this returning address of p
}

func fun(copyOfP Product){
	// this p is a copy of the product passed to this function
	copyOfP.name="macbook"

}

func pass_ref(p *Product){
	p.name = "macbook pro"            // -> this is changing the object
} 
 
// member function of a struct
func (p *Product) display(short bool){
	if(short){
		fmt.Println("Name     ",p.name)
		return
	}

	fmt.Println("Name",p.name)
	fmt.Println("Price",p.price)
	fmt.Println("Name",p.company)

}

func (i *Product) buy(){
	fmt.Println("name is ",i.name,"price is", i.price)
}

func (p *Product) getDiscount() int {
	discount := p.price*20/100
	fmt.Print("name of the discounted product",p.name,"discount is",discount)
	return discount
}


func main(){

	new_p:=newProduct("iphone 15",1000,"apple inc");   // new_p is the pointer to the product
	fmt.Println(new_p.name)
	fmt.Println(new_p.price)
	fmt.Println(new_p.company)
	
	pass_ref(new_p)
	
	fmt.Println(new_p.name)
	new_p.display(true)

	fmt.Println("........................")
	
	
	check_discountAndBuy(new_p)
	
	fmt.Println("..................")





	p:= Product{
		name:"iphone",
		price : 1000,
		company: "Apple",
	}
	fmt.Println(p.name)

	fun(p)        // passing a copy of the object and getting the same object not updated object

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
	//loops()
	// arrays()
		//maps()

	// x,y :=check_oddEven(10)
	// fmt.Println(x,y)
pointers()
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

func arrays(){
	var arr2 [5]int   // array
	arr2[0]=1
	arr2[1]=15
	arr2[2]=1321
	arr2[3]=132
	arr2[4]=13
	//fmt.Print(arr1)
	
	var arr1 []int;   // this is a slice -> dynamic size array
	arr1=append(arr1, 34,76,86,98)  // -> append only works in slice
	 arr := []int{1,2,3,4}
	combined := append(arr1,arr...)       // -> ellipsis operator
	fmt.Print(combined) 


}

func maps(){

	products := map[string]int{
		"Product 1":123,
		"Product 2":324,
		"Product 3":14324,
	}

	fmt.Print(products)

	customMap := map[string]int{}
	fmt.Print(customMap)


	customMap1 :=make( map[string]int)
	fmt.Print(customMap1)

	customMap["string"]=322

}

func check_oddEven(num int) (string,int){
	if(num %2==0){
		return "Even", 1
	} else{
		return "Odd",0
	}
	
}

func pointers(){
	i:=1120
	val:=&i
	fmt.Println("value of i",i)
	fmt.Println("pointer to i",&val)
	fmt.Println("value of val",*val)
}