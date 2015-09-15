package main

import "fmt"
import "syscall"
import "os"
import "math/rand"

func main() {
	food := RandomFood()
	env := os.Environ()

	fmt.Printf("Starting tmux session with name: %v\n", food)

	err := syscall.Exec("/usr/local/bin/tmux", []string{"tmux", "new-session", "-s", food}, env)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func RandomFood() string {
	foods := []string{
		"lasagna",
		"pizza",
		"empanadas",
		"tacos",
		"enchiladas",
		"burritos",
		"sushi",
		"fried chicken",
		"hot wings",
		"donuts",
		"cake",
		"muffins",
		"pie",
		"cinnamon rolls",
		"fries",
		"fried catfish",
		"ham",
		"turkey",
		"macaroni and cheese",
		"cheesy potatoes",
		"cheeseburgers",
		"cottage cheese",
		"sausage",
		"bratwurst",
		"mozzarella sticks",
		"cheese bread",
		"cheese curds",
		"pasta salad",
		"beef stew",
	}
	return foods[rand.Intn(len(foods))]
}
