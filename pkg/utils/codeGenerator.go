package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	CHARSET    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RANDLENGTH = 5
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateCode(prefix string) string {
	t := time.Now()
	var y, m, d string
	y = makeTwoDigitChunk(t.Year() % 100)
	m = makeTwoDigitChunk(int(t.Month()))
	d = makeTwoDigitChunk(t.Day())
	code := fmt.Sprintf("%s%s%s%s%s", prefix, y, m, d, generateRandStringWithLength(RANDLENGTH))
	return strings.ToUpper(code)
}

func makeTwoDigitChunk(expression int) string {
	if expression < 10 {
		return fmt.Sprintf("0%d", expression)
	} else {
		return fmt.Sprintf("%d", expression)
	}
}

func generateRandStringWithLength(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}
	return string(b)
}
