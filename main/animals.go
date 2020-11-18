package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name, food, locomotion, sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func (a Animal) DoSomething(behavior string) {
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
		cow := Animal{name: "cow", food: "grass", locomotion: "walk", sound: "moo"}
		bird := Animal{name: "bird", food: "worms", locomotion: "fly", sound: "peep"}
		snake := Animal{name: "snake", food: "mice", locomotion: "slither", sound: "hsss"}
		switch name := args[0]; name {
		case "cow":
			cow.DoSomething(args[1])
		case "bird":
			bird.DoSomething(args[1])
		case "snake":
			snake.DoSomething(args[1])
		}
	}
}
