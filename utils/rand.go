package utils

import (
	"math/rand"
	"time"
)

func RandInt(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(n)
}