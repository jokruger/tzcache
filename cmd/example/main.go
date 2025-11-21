package main

import (
	"github.com/jokruger/tzcache"
)

func main() {
	tzc := tzcache.NewSafe()
	l1, err := tzc.Location("UTC")
	if err != nil {
		panic(err)
	}
	l2, err := tzc.Location("Europe/Kyiv")
	if err != nil {
		panic(err)
	}
	println("UTC Location:", l1.String())
	println("Europe/Kyiv Location:", l2.String())
	println("Cache Size:", tzc.Size())
}
