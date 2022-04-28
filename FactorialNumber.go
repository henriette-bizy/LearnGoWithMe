package main

import "fmt"


var SumOfFactors uint64 = 1;

func factorize(n int) uint64{

for i:=1; i<=n ; i++{
	 SumOfFactors *= uint64(i);
}

return SumOfFactors;

}

func main(){

var number int
fmt.Println("Welcome to factorial season, Comment down the number you want to factorise")
fmt.Scan(&number)
fmt.Printf("the factorial number of %d is : %d",number , factorize(number));


}