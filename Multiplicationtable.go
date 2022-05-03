package main

import "fmt"


func multiplicationTable(number int){
	for i:=1; i<=10 ; i++{
   fmt.Printf(" %d * %d =  %d \n", i,number, (i*number))
	}
}


func main(){
    var n int
	println("Put the number to find for multiplication")
	fmt.Scan(&n)
    multiplicationTable(n)

}