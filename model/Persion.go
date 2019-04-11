package model

type persion struct{
	Name string
	age int 
	sal int
}


func Newpersion(name string) *persion{
	
	return &persion{
		Name:name,
	}
}

func (p *persion)Set_data(age int, sal int){
	p.age = age
	p.sal = sal
}


func (p *persion)Get_data()(int, int){
	return p.age ,p.sal
}
