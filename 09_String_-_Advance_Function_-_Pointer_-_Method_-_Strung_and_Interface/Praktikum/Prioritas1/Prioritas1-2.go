package main

import "fmt"

type Student struct {
	name  []string
	score []float64
}

func (s Student) averageScore() float64 {
	arrScore := s.score
	var totalScore float64
	for _, v := range arrScore {
		totalScore += v
	}

	var average float64 = totalScore / float64(len(s.score))
	return average
}

func (s Student) minScore() (string, float64) {
	arrScore := s.score
	arrName := s.name
	var minScore float64
	var name string

	for i := 0; i < len(arrScore); i++ {
		if i == 0 {
			minScore = arrScore[i]
			name = arrName[i]
		} else if minScore < arrScore[i] {
			minScore = arrScore[i]
			name = arrName[i]
		}
	}
	return name, minScore
}

func (s Student) maxScore() (string, float64) {
	arrScore := s.score
	arrName := s.name
	var maxScore float64
	var name string

	for i := 0; i < len(arrScore); i++ {
		if i == 0 {
			maxScore = arrScore[i]
			name = arrName[i]
		} else if maxScore > arrScore[i] {
			maxScore = arrScore[i]
			name = arrName[i]
		}
	}
	return name, maxScore
}

func main() {
	myStudent := Student{
		[]string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		[]float64{80, 75, 70, 75, 60},
	}

	minName, minScore := myStudent.minScore()
	maxName, maxScore := myStudent.maxScore()

	fmt.Println("Average Score: ", myStudent.averageScore())
	fmt.Println("Min Score of Students: ", minName, "(", minScore, ")")
	fmt.Println("Max Score of Students: ", maxName, "(", maxScore, ")")

}
