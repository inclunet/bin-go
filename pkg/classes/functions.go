package classes

import (
	"math/rand"
	"time"
)

func GetRandomNumber(start int, end int) int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	return rand.Intn(end-start+1) + start
}
