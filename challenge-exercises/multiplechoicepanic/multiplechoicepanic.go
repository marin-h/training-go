package main

import "fmt"

type Question struct {
	Options     int
	PointsWrong float32
	PointsRight float32
}

type Exam struct {
	PointsPass float32
	Questions  []Question
}

func (exam Exam) isExpectedToPassExam() bool {
	var oddsPass int
	var oddsFail int
	var examOdds [][]float32
	// put together the odds for each question
	for _, question := range exam.Questions {
		var questionOdds []float32
		questionOdds = append(questionOdds, question.PointsRight)
		for i := 0; i < question.Options-1; i++ {
			questionOdds = append(questionOdds, question.PointsWrong)
		}
		questionOdds = append(questionOdds, 0)
		examOdds = append(examOdds, questionOdds)
	}
	// iterate over the odds now, and decide if they pass or fail
	for i, questionOdds := range examOdds {
		for _, questionOdd := range questionOdds {
			odd := questionOdds[i] + questionOdd
			if odd >= exam.PointsPass {
				oddsPass = oddsPass + 1
			} else {
				oddsFail = oddsFail + 1
			}
		}
	}
	// check if odds are more passing than failing
	return oddsPass > oddsFail
}

func main() {

	exam1 := Exam{1, []Question{{2, -1, 2}, {2, -1, 2}}}
	exam2 := Exam{2, []Question{{2, -1, 0.5}, {4, -4, 1}, {2, -1, 1}}}
	exams := []Exam{exam1, exam2}
	for _, exam := range exams {
		if exam.isExpectedToPassExam() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
