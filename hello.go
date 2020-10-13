package hello

import (
	"github.com/zhouwentong1993/hello0"
	"rsc.io/quote/v3"
)

func Hello() string {
	return hello.Hello()
}

func Proverb() string {
	return quote.Concurrency()
}
