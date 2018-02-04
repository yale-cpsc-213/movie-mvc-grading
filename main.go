package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yale-cpsc-213/movie-mvc-grading/grade"
	"github.com/yale-cpsc-213/movie-mvc-grading/questions"
)

const usage string = `
Creates and grades movie mvc homework for CPSC213. Usage:
cpsc213moviemvc test URL
`

func main() {
	log.SetFlags(log.Lshortfile)
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	url := os.Args[2]
	switch os.Args[1] {
	case "grade-all":
		err := grade.Grading(url)
		if err != nil {
			log.Fatal(err)
		}
	case "test":
		questions.TestAll(url, true)
	default:
		fmt.Println("ERROR! Bad input. See below for usage.")
		fmt.Println(usage)
		return
	}

}
