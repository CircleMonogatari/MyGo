package snake

import (
	"fmt"
)

type Scene struct{
	scene [][]string
	X, Y int
}

func NewScene(MaxX, MaxY int) *Scene  {
	return &Scene{X:MaxX, Y:MaxY,}
}

func (s *Scene) Init()  {
	wall := "墙"
	s.scene = make([][]string, s.X)
	for i := 0; i < s.X; i++ {
		s.scene[i] = make([]string, s.Y)
	}



	for i := 0; i < s.Y; i++ {
		s.scene[0][i] = wall
		s.scene[s.X - 1][i] = wall
	}	
	for i := 0; i < s.Y; i++ {
		s.scene[i][0] = wall
		s.scene[i][s.Y-1] = wall
		}	
	
		for i := 1; i < s.X-1; i++ {
			for j := 1; j < s.Y-1; j++ {
				s.scene[i][j] = "  "
			}}

}

func (s *Scene) Draw()  {
		for i := 0; i < s.X; i++ {
		for j := 0; j < s.Y; j++ {
			fmt.Printf(s.scene[i][j])
		}
		fmt.Println()
	}
}