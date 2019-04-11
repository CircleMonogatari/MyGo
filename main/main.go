package main

import (
	"fmt"
	"MyGo/mydemo"
	"MyGo/model"
)
func main()  {
	mydemo.Say()

	p := model.Newpersion("saofuren")
	p.Set_data(18, 5000)
	fmt.Println(p)

	fmt.Println(p.Get_data())
}