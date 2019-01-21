package main

import (
	"fmt" 
	"flag"
)

/**
	Find the sum of all the multiples of 3 or 5 below 1000.
*/

func main(){
	fmt.Println("Find sum of all the multiples of 3 and 5 below a maximum number (1000 is default).")
	
	maxNumber := flag.Int("maxNumber", 1000, "Maximum Number")
	flag.Parse() //Parse command line arguement for maximum number

	sum := 0 //Initialise sum

	for i:=0; i < *maxNumber; i++ {
		if ( i % 3 == 0 || i % 5 == 0){
			sum += i
		}
	}
	fmt.Println(sum) //print result
}