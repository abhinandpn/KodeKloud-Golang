package datatype

// package for print somthing
import "fmt"

func Display() {

	name := "abhinandpn" // short variable declartion (datatype >> string)
	number := 1          // short variable declartion (datatype >> number)
	float := 2.4         // short variable declartion (datatype >> float)4

	fmt.Println(name) // printing command
	fmt.Println(number)
	fmt.Println(float)
}

func VarDisplay() {

	first := "abhinand"
	last := "pn"

	fmt.Println("name is ", first+last)
}

func InputStr() {
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scanf("%s", &name)
	fmt.Println("Your name is : ", name)
}

func MultIput() {
	var a string
	var b int

	fmt.Println("enter a string and initger")

	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count : ", count)
	fmt.Println("error : ", err)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}
