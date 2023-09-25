package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// "fmt"
// "time"

// func main() {
// 	var now time.Time = time.Now()
// 	//var year int = now.Year()
// 	year := now.Year()
// 	var month string = now.Month().String()
// 	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
// }

//	func main() {
//		HotSpurs := "hm ? j madi?"
//		replacePlayer := strings.NewReplacer("?", "son")
//		player := replacePlayer.Replace(HotSpurs)
//		fmt.Println(player)
//	}
// func main() {
// 	fmt.Print("Input score : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	inputScore, err := reader.ReadString('\n') // Opt 2
// 	log.Fatal(err)
// 	fmt.Println(inputScore)

// }
func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n') // Opt 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      // remove space bar
	inputScore, err := strconv.ParseFloat(inputScoreString, 64) // casting
	if err != nil {
		log.Fatal(err)
	}
	var grade string
	if inputScore >= 90 { // mismatched types string and untyped int
		grade = "A grade!"
	} else {
		grade = "BCDF grade~"
	}
	fmt.Println("you got " + grade)

}
