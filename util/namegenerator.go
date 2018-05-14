package util

import (
	"time"
	"fmt"
	"math/rand"
)

func NameGenerator() string {
	y, m, d := time.Now().Date()
	seed := 1000 + rand.Intn(8999)
	name := fmt.Sprintf("%v%.2v%.2v%v", y, int(m), d, seed)
	return name
}
