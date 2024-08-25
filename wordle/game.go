package main

import (
	"fmt";
	"os";
	"strings"
)

func check(og_word string, your_word string) ([5]int, [5]int){

	var full_indices = [5]int{-1, -1, -1, -1, -1}
	var partial_indices = [5]int{-1, -1, -1, -1, -1}

	if(og_word == your_word){
		fmt.Println("You win!")
		os.Exit(0)
	}

	for index := range og_word{
		if(og_word[index] == your_word[index]){

			full_indices[index] = index
		} else if(strings.ContainsAny(og_word, string(your_word[index]))){
			partial_indices[index] = index
		}
	}

	return full_indices, partial_indices

}

func main(){

	answer := "fangs"

	for i := range 6 {
		
		var guess string

		fmt.Printf("round %d: ", i+1)
		fmt.Scan(&guess)

		// validate input
		if len(guess) != 5 {
			fmt.Println("enter a 5 lettered word sucka!!!")
			continue
		}

		full, partial := check(answer, guess)

		// display your word
		for i := range guess{
			fmt.Printf("%s ", string(guess[i]))
		}

		fmt.Println("")

		// display grid
		for i := 0; i < 5; i++ {
			if(full[i] != -1){
				fmt.Printf("%s ", string(answer[i]))
			} else if(partial[i] != -1){
				fmt.Print("* ")
			} else{
				fmt.Print("_ ")
			}
		}

		fmt.Printf("\n\n")


	}

}