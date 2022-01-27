package main

import (
	"fmt"
	"strconv"
)

func main() {
	file, err := NewFile(getFilePath())
	if err != nil {
		panic(err)
	}

	println("Parsing file...\n")

	var (
		fileSize = 0
		grades   = []float64{}
	)

	fileSize, err = strconv.Atoi(file[0])
	if err != nil {
		panic(err)
	}

	for _, strGrade := range file[1:] {
		grade, err := strconv.ParseFloat(strGrade, 64)
		if err != nil {
			panic(err)
		}

		grades = append(grades, grade)
	}

	println(fmt.Sprintf("File size: %d\nGrades:%+v", fileSize, grades))
}

func getFilePath() (str string) {
	println("***Welcome to the Exam Statistics Program!!***")
	println("Please enter a file name: (default path: './grades/zz_grades.txt')")

	fmt.Scanln(&str)

	if str == "" {
		return "./grades/zz_grades.txt"
	}

	return str
}
