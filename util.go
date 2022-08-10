package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)[:len]
}
