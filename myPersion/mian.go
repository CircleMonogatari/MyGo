package main

import "MyGo/myPersion/model"

import "fmt"

func main()  {
	var p = model.NewPersion("pp", 30)
	p.Show()

	s := model.NewSon("ss", "pp", 12)

	s.Show()
	fmt.Println(s)
}