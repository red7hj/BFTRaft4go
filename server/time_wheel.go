package server

import (
	"context"
	"math/rand"
	"time"
	"log"
)

func RandomTimeout(mult float32) int {
	lowRange := 1000 * mult
	highRange := 5000 * mult
	return int(lowRange + highRange*rand.Float32())
}
