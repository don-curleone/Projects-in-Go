package main

import "fmt"

func add(num1 int32, num2 int32) int32{
	return num1 + num2
}

func sub(num1 int32, num2 int32) int32{
	return num1 - num2
}

func mul(num1 int32, num2 int32) int32{
	return num1 * num2
}

func div(num1 int32, num2 int32) int32{
	return num1 / num2
}




func main(){

	var op string
	var num1 int32
	var num2 int32

	for true{
		 
	fmt.Print("Enter <operator> <arg1> <arg2> : ")
	fmt.Scanln(&op, &num1, &num2)

	if (op == "add"){
		fmt.Println(add(num1, num2))
	} else if (op == "sub"){
		fmt.Println(sub(num1, num2))
	} else if (op == "mul"){
		fmt.Println(mul(num1, num2))
	} else if (op == "div"){
		fmt.Println(div(num1, num2))
	} else {
		fmt.Println("Invalid operator")
	}
}
 



}