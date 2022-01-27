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

	println(fmt.Sprintf("File size: %d\nGrades:%+v\n", fileSize, grades))

	min, max := getMinMax(grades)
	println(fmt.Sprintf("Min: %f, Max: %f", min, max))

	letterGrades := getLetterGrades(grades)
	println(fmt.Sprintf("Letter grades: %+v", letterGrades))
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

func getMinMax(floats []float64) (float64, float64) {
	var (
		min, max float64
	)

	for i, f := range floats {
		if i == 0 {
			min, max = f, f
			continue
		}

		if f < min {
			min = f
		}

		if f > max {
			max = f
		}
	}

	return min, max
}

func getLetterGrades(floats []float64) map[string]int {
	m := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
		"F": 0,
	}

	for _, f := range floats {
		if f >= 90 {
			m["A"]++
		} else if f >= 80 {
			m["B"]++
		} else if f >= 70 {
			m["C"]++
		} else if f >= 60 {
			m["D"]++
		} else {
			m["F"]++
		}
	}

	return m
}
