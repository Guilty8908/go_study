package main
import "fmt"

func main(){
	var x int
	x = 1
	var p *int
	p = &x
	fmt.Println(p)
	add(&x)
	fmt.Println(x)	
}

func add(p *int){
	*p = *p + 1
}
