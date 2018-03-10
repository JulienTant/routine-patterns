package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomStringsGenerator(nbString, nbChar int) <-chan string {
	var src = rand.NewSource(time.Now().UnixNano())

	cs := make(chan string)

	go func(nbString, nbChar int) {
		for i := 0; i < nbString; i++ {
			b := make([]byte, nbChar)

			for i, cache, remain := nbChar-1, src.Int63(), letterIdxMax; i >= 0; {
				if remain == 0 {
					cache, remain = src.Int63(), letterIdxMax
				}
				if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
					b[i] = letterBytes[idx]
					i--
				}
				cache >>= letterIdxBits
				remain--
			}

			cs <- string(b)
		}
		close(cs)

	}(nbString, nbChar)

	return cs
}

func main() {
	for i := range randomStringsGenerator(100, 16) {
		fmt.Println(i)
	}
}
