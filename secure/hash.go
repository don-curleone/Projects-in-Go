package main

import (
	"fmt";
	"crypto/sha256";
	"os"
)

func main(){

	args := os.Args[1:]

	if(len(args) == 1){
		word := os.Args[1]
		hash := sha256.Sum256([]byte(word))
		fmt.Printf("%x\n", hash)

	} else if(os.Args[1] == "-f"){

		filepath := os.Args[2]
		data, err := os.ReadFile(filepath)

		if(err != nil){
			fmt.Println("File reading error", err)
        	return
		}
		hash := sha256.Sum256([]byte(data))
		fmt.Printf("%x\n", hash)

	}

}