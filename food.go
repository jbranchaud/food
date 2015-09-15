package main

import "fmt"
import "syscall"
import "os"

func main() {
	food := "lasagna"
	env := os.Environ()

	fmt.Printf("Starting tmux session with name: %v\n", food)

	err := syscall.Exec("/usr/local/bin/tmux", []string{"tmux", "new-session", "-s", food}, env)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
