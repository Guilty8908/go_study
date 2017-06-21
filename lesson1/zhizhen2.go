package  main
import "fmt"

func main(){
	var x int
	var y int
	x = 1
	y = 2
	swap(&x,&y)
	fmt.Println(x,y)
}

func swap(p *int,q *int){
	var t = *p
	*p = *q
	*q = t
}
