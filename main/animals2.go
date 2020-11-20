package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func DoSomething(a Animal, behavior string) {
	switch behavior {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	cow := Cow{name: "cow", food: "grass", locomotion: "walk", sound: "moo"}
	bird := Bird{name: "bird", food: "worms", locomotion: "fly", sound: "peep"}
	snake := Snake{name: "snake", food: "mice", locomotion: "slither", sound: "hsss"}
	for {
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()
		if strings.Compare(request, "") == 0 {
			break
		}
		args := strings.Split(request, " ")
		if len(args) < 2 {
			break
		}
		switch name := args[0]; name {
		case "cow":
			DoSomething(cow, args[1])
		case "bird":
			DoSomething(bird, args[1])
		case "snake":
			DoSomething(snake, args[1])
		}
	}
}
