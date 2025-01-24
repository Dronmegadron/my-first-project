package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go-first-floor-sprint-four/ftracker"
)

func main() {
	action := 10000
	trainingType := "Бег"
	duration := 1.0
	weight := 67.0
	height := 173.0
	lengthPool := 50
	countPool := 20

	result := ftracker.ShowTrainingInfo(action, trainingType, duration, weight, height, lengthPool, countPool)
	fmt.Println(result)
}
