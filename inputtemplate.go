package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	fmt.Printf("str : %s\n",str)

}
