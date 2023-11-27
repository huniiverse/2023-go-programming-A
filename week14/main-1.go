package main

import "fmt"

func main() {
	// 두개의 차이점 알아두기
	// var balls map[string]int
	balls := make(map[string]int)
	fmt.Printf("%v\n", balls)
	balls["성기훈"] = 20
	fmt.Println(balls["성기훈"])
	fmt.Println(balls["오일남"])
}
