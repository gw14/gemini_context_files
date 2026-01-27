package main

import "fmt"

//position
type Position struct{
	x float64
	y float64
}

func (p *Position) moveX(step float64){
	p.x += step
}

func (p *Position) moveY(step float64){
	p.y += step
}

func (p Position) showPosition() string{
	return ("x: " + (string)p.x + ",y: " + (string)p.y)
}

type Player struct{
	name string
	Position
}

func (p Player) showPlayer() string{
	return ("name: " + p.name + ",Position: " + p.showPosition)
}

type Enemy struct{
	name string
	Position
}

func (e Enemy) showEnemy() string{
	return ("name: " + p.name + ", Position: " + p.showPosition)
}

type StaticObstacle struct{
	material string
	Position
}

func main(){
	hero := Player{
		name: "Guy",
		Position: Position{x: 80.0,y: 59.5},
	}
	boss := Enemy{
		name: "Avi",
		Position: Position{x: 50.0,y: 63.0},
	}
//	rock := StaticObstacle{
//		material: "stone",
//		Position: Position{x: 1.0,y: 5.0},
//	}

	fmt.Printf("%s\n",hero.showPlayer)
	fmt.Printf("%s\n",boss.showEnemy)
	hero.moveX(2)
	hero.moveY(3)
	boss.moveX(5)
	boss.moveY(-3)

}
