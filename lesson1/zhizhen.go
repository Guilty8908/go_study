package main
import "fmt"

func main(){
	var x int
	x = 1
	var p *int
	p = &x
	fmt.Println("p","=", p)
	fmt.Println("*p=",*p)
	fmt.Println("x=",x)
	*p = 2
	fmt.Println("x=",x)
	fmt.Println("hello  golang")
}
