package main

import(
	"fmt"
	"github.com/jreisinger/gokatas/counter"
)


func main(){
	c:=new(counter.Counter)
	c.Increment()
	fmt.Println(c.Get())
	c.Reset()
}