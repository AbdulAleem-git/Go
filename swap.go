package main

import "fmt"

func main(){
	x := 20 
	y := 30 
	fmt.Printf("before swap x = %d y = %d\n", x, y)

	x = x + y 
	y = x -y 
	x = x-y 
	fmt.Printf("after swap x = %d y = %d\n", x, y )


	// The other way of Swapping
	swap( &x  , &y)
	fmt.Printf("after swap x = %d y = %d\n", x, y )

}
func swap(x *int , y *int){
	*x , *y = *y , *x
}