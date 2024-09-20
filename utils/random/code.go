package random

import (
	"math/rand"
	"time"
)

var stringCode = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var bytes = make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = stringCode[rand.Intn(len(stringCode))]
	}
	return string(bytes)

}
