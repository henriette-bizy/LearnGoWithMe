package main 

import "fmt"

//order in which functions are written doesn't change

func main(){
	totalScore := ask(0,"fish");
ask(1,"cat");
ask(2,"dog");

}

func ask(number int , question string) int{
	var input string;
	fmt.Printf("{Question %d} please input the following word : %s \n", number , question);
	fmt.Scan(&input);
	fmt.Printf("%s was the input \n", input);

	if question == input {
		fmt.Println("correct")
		return 10
	
	}else{
		fmt.Println("incorrect")
		return 0
	}
}