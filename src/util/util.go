package util

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(length int) string {
	output := make([]byte, length)

	for i := range output {
		output[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(output)
}

func GetTimestamp() int64 {
	return time.Now().Local().UnixNano() / int64(time.Millisecond)
}
