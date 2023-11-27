package main

import "fmt"

func main() {
	// var games map[int]string
	// games = make(map[int]string)

	// map literal
	games := map[int]string{
		456: "이훈진",
		218: "송윤빈",
		067: "최은솔",
		001: "어재선",
		199: "감승엽",
		101: "고은지",
	}
	// fmt.Println(games[100])
	name, ok := games[101]
	fmt.Println(name, ok)
	// map은 순서 개념이 없다, 하지만 짝은 항상 유지가 된다.

	// append
	// games[456] = "이훈진"
	// games[218] = "송윤빈"
	// games[067] = "최은솔"
	// games[001] = "어재선"
	// games[199] = "감승엽"
	// games[101] = "고은지"

	for _, v := range games {
		fmt.Println(v)
	}

	games[101] = "장덕수" // update
	delete(games, 199) // delete 삭제

	for k, v := range games {
		fmt.Println(k, v)
	}
}
