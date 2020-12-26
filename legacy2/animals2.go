package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, sound string
}

type Bird struct {
	name, food, locomotion, sound string
}
type Snake struct {
	name, food, locomotion, sound string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}
func (a Cow) Move() {
	fmt.Println(a.locomotion)
}
func (a Cow) Speak() {
	fmt.Println(a.sound)
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}
func (a Bird) Move() {
	fmt.Println(a.locomotion)
}
func (a Bird) Speak() {
	fmt.Println(a.sound)
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}
func (a Snake) Move() {
	fmt.Println(a.locomotion)
}
func (a Snake) Speak() {
	fmt.Println(a.sound)
}

func main() {
	animals := make(map[string]Animal)
	for {
		var first, second, third string
		fmt.Printf("\n>")
		_, _ = fmt.Scan(&first)
		if first == "newanimal" {
			_, _ = fmt.Scan(&second)
			_, _ = fmt.Scan(&third)
			switch third {
			case "cow":
				animals[second] = Cow{name: "cow", food: "grass", locomotion: "walk", sound: "moo"}
			case "bird":
				animals[second] = Bird{name: "bird", food: "worms", locomotion: "fly", sound: "peep"}
			case "snake":
				animals[second] = Snake{name: "snake", food: "mice", locomotion: "slither", sound: "hsss"}
			}
		} else if first == "query" {
			_, _ = fmt.Scan(&second)
			_, _ = fmt.Scan(&third)
			switch third {
			case "eat":
				animals[second].Eat()
			case "move":
				animals[second].Move()
			case "speak":
				animals[second].Speak()
			}
		} else {
			break
		}
	}
}
