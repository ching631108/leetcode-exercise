package main

import "fmt"

func main() {
	result := ladderLength("hit", "cog", []string{
		"hot", "dot", "dog", "lot", "log", "cog"})
	// result := ladderLength("hit", "cog", []string{
	// 	"hot", "dot", "dog", "lot", "log"})
	fmt.Println(result)
}
