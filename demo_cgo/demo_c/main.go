package main

import "fmt"

func set_num(i *int){
	*i = 15
}

func main()  {
	var num int
	num = 10

	set_num(&num)

	fmt.Println(num)
}