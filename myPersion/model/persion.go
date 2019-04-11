package model

import "fmt"

type persion struct{
	Name string
	age int
}

func NewPersion(name string, age int) *persion{
	return &persion{
		Name:name,
		age :age,
	}
}

func (p*persion)Show(){
	fmt.Println(p)
}


//继承
type son struct{
	persion
	mom string
}

func NewSon(name string, mom string, age int) *son{

	s :=son{
		mom:mom,
		persion: persion{Name:name, age:age,},
	}

	return &s
}