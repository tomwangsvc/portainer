package main

import "fmt"

func main() {
	container := NewContainer([][2]string{{"{", "}"}, {"[", "]"}, {"(", ")"}})
	fmt.Println(container.Valid("[]{}[[[{}]{}]]{}{[]{[]}}"))
}
