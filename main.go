package main

import (
	. "exercism/easy"
	. "exercism/hard"
	. "exercism/medium"
	"fmt"
)

func main() {
	// easy

	// 1- https://exercism.org/tracks/go/exercises/largest-series-product
	var series string

	fmt.Printf("Enter series \n")
	fmt.Scan(&series)

	fmt.Println(LargestSeriesProduct(series, 3))

	// 2- https://exercism.org/tracks/go/exercises/raindrops
	var num int
	var res string

	fmt.Printf("Enter any number for factor check of 3, 5, 7 \n")
	fmt.Scan(&num)

	Raindrop(&num, &res)

	fmt.Printf(res)
	fmt.Printf("\n")

	// 3- https://exercism.org/tracks/go/exercises/isogram
	var word string
	var isogramRes bool

	fmt.Printf("Enter any word to check for isogram \n")
	fmt.Scan(&word)

	isogramRes = IsIsogram(word)

	if isogramRes {
		fmt.Printf("word is an isogram \n")
	} else {
		fmt.Printf("word is not an isogram \n ")
	}

	// 4- https://exercism.org/tracks/go/exercises/two-fer
	var doEnterFriendName bool
	var name string

	fmt.Printf("Do you want to enter your friends name true or false \n")
	fmt.Scan(&doEnterFriendName)

	if doEnterFriendName {
		fmt.Printf("Enter your friends name \n")
		fmt.Scan(&name)
		fmt.Println(TwoFer(name))
	} else {
		name = ""
		fmt.Println(TwoFer(name))
	}

	// 5- https://exercism.org/tracks/go/exercises/protein-translation
	var RNA string

	fmt.Printf("Enter RNA \n")
	fmt.Scan(&RNA)

	fmt.Println(ProteinTranslation(RNA))

	// medium

	// 1- https://exercism.org/tracks/go/exercises/simple-linked-list
	LinkedlistReverse()

	// 2- https://exercism.org/tracks/go/exercises/transpose
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	Transpose(input)

	// 3- https://exercism.org/tracks/go/exercises/change
	coins := []uint16{
		1, 5, 10, 25, 100,
	}
	fmt.Print(Change(40, coins))

	// 4- https://exercism.org/tracks/go/exercises/matching-brackets
	MatchingBrackets("((){}[()])")

	// 5- https://exercism.org/tracks/go/exercises/flatten-array
	texts := []string{
		"Ibrahim",
		"Jamil",
		"Hamza",
	}
	fmt.Println(ParallelFrequencyCount(texts))

	// hard
	// https://exercism.org/tracks/go/exercises/forth
	var code string = "3 5 + 2 * 10 /"

	result := Forth(code)
	fmt.Print(result)
}
