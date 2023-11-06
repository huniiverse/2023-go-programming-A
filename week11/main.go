package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])
	//var primes [3]int = [3]int{2, 3, 5} // primes 배열을 배열 리터럴로 초기화
	primes := [3]int{2, 3, 5} // 단축 연산자
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) // boolean 타입의 제로값, false

	//fmt.Println(primes[4]) // invalid argument: index 4 out of bounds [0:3]
	i := 0
	//for i < 4 { //panic: runtime error: index out of range [3] with length 3
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}
	// fmt.Println()
	// for j := 0; j < len(primes); j++ {
	// 	fmt.Println(primes[j])
	// }
	fmt.Println()
	for prime := range primes { // 값만 출력하려 했으나 인덱스가 출력됨
		fmt.Println(prime)
	}

}
