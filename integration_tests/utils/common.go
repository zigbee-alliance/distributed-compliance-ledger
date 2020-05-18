package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandString() string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var b strings.Builder

	for i := 0; i < 5; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func ParseUint(num string) uint64 {
	val, _ := strconv.ParseUint(num, 10, 64)
	return val
}
