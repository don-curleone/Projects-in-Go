package main

import (
	"fmt";
	"math/rand";
)

func main(){
	answer := rand.Intn(100)

	for true{
		
	var guess int
	
	fmt.Print("Enter a number: ")
	fmt.Scan(&guess)

	
	if(guess == answer){
		fmt.Println("Congrats! You win")
		break
	} else if(guess > answer){
		fmt.Println("Lower")
	} else if(guess < answer){
		fmt.Println("Higher")
	}

	}
	

}