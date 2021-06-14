package common

import (
	"math/rand"
	"time"
)

// generate an EOS compliant account name
// length < 12
// charset: a..z12345
func ToolsAccountGenerateName(fiveCharsMaxPrefix string) string{

	if len(fiveCharsMaxPrefix) > 5{
		fiveCharsMaxPrefix = fiveCharsMaxPrefix[0:5]
	}

	return fiveCharsMaxPrefix+randStringBytesMaskImprSrc(7)
}

func ToolsWalletGenerateName(prefix string) string{

	return prefix+randStringBytesMaskImprSrc(7)
}

var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "12345"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
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

	return string(b)
}