package snake

import (
	"fmt"
)


func init()  {
	fmt.Println("snake start")
	GameStart()
	fmt.Println("snake End")
}

func GameStart()  {
	draw := NewScene(10, 10)
	draw.Init()
	draw.Draw()
}