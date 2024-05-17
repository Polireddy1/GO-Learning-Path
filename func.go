package main
import "fmt"

func plus(a int,b int) int{
	return a+b
}
func plusplus(a int,b int,c int) int{
	return a+b+c
}

func sub(a int,b int) int{
	return a-b
}
func mod(a int,b int) int{
	return a%b
}
func mul(a int, b int) int{
	return a*b
}
func main(){ 
	res :=plus(1,2)
	fmt.Println("1+2=",res)

	res=plusplus(1,2,3)
	fmt.Println("1+2+3=",res)

	res=sub(9,5)
	fmt.Println("9-5 =",res)

	res=mod(7,3)
	fmt.Println("7%3 =",res)

	res=mul(9,8)
	fmt.Println("9*8 =",res)
}