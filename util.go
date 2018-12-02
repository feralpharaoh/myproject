package main

import (
	"math/rand"
	"time"
)

var runesForString []rune = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)

func RandomString(length int) string {
	val := make([]rune, length)
	for i := range val {
		val[i] = runesForString[rand.Intn(len(runesForString)-1)]
	}
	return string(val)
}

func init() {
	rand.Seed(time.Now().Unix())
}
