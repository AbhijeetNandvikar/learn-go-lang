package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in format of question,answer")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Error message :", err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error message :", err)
	}

	problems := parseLines(lines)
	score := 0

	for i, p := range problems {
		fmt.Printf("Q #%d: %s =  \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("correct!!")
			score++
		}
	}
	fmt.Printf("Total Score: %d out of %d. \n", score, len(problems))
	// Correct Answers count => score

	// fmt.Println("Questions:")
	// // var score int
	// for i, v := range lines {
	// 	var answer int
	// 	fmt.Println("Q-%v: %v \n", i, v[0])
	// 	// fmt.Println(v, i)
	// 	fmt.Scanf("Answer:", answer)

	// 	fmt.Println("Your input is:", answer)
	// }

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
