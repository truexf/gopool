package main

import (
	"fmt"
	"github.com/truexf/gopool"
)

type Person struct {
	Name string
	Age int
	Country string
}

func CreateObj() interface{} {
		fmt.Println("new obj") 
		return &Person{Name:"unknown", Country: "unknown"}
}
func main() {
	goPool := gopool.NewGoPool(100, CreateObj)
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 50; i++ {
			v := goPool.Get()
			if v == nil {
				fmt.Println("Pool get nil")
			} else {
				fmt.Printf("%v\n", v)
			}
		}
		ch <- 1
	}()

	go func() {
		for i := 0; i< 100; i++ {
			v := &Person{Name: "fangyousong", Age: 36, Country: "chinese"}
			goPool.Put(v)
		}
		ch <- 1
	}()

	<- ch
	<- ch
	fmt.Printf("poolsize: %d\n", goPool.Size())
	return
}