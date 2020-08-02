package main

import "fmt"

type Question struct {
	Options     int
	PointsWrong float32
	PointsRight float32
}

type Exam struct {
	PointsPass int
	Questions  []Question
}

func (exam Exam) isExpectedToPassExam() bool {
	var oddsPass int
	var oddsFail int
	var examOdds [][]float32

	// TODO: put this in another function
	for _, question := range exam.Questions {
		var questionOdds []float32
		questionOdds = append(questionOdds, question.PointsRight)
		for i := 0; i < question.Options-1; i++ {
			questionOdds = append(questionOdds, question.PointsWrong)
		}
		questionOdds = append(questionOdds, 0)
		examOdds = append(examOdds, questionOdds)
	}

	// for i, questionOdds := range examOdds {
	// 	for j, questionOdd := range questionOdds {
	// 	}
	// }

	fmt.Println(examOdds)
	return oddsPass > oddsFail
}

func main() {

	exam1 := Exam{1, []Question{{2, -1, 2}, {2, -1, 2}}}
	exam2 := Exam{2, []Question{{2, -1, 0.5}, {4, -4, 1}, {2, -1, 1}}}
	exams := []Exam{exam1, exam2}
	for _, exam := range exams {
		fmt.Println(exam.isExpectedToPassExam())
	}
}
