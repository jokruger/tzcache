package main

import (
	"github.com/jokruger/tzcache"
)

func main() {
	tzc := tzcache.NewUnsafe()
	l1, err := tzc.Get("UTC")
	if err != nil {
		panic(err)
	}
	l2, err := tzc.Get("Europe/Kyiv")
	if err != nil {
		panic(err)
	}
	println("UTC Location:", l1.String())
	println("Europe/Kyiv Location:", l2.String())
	println("Cache Size:", tzc.Size())
}
