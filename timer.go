package main

import "time"

type StrIntFunction func(string) int

func Timer(function StrIntFunction, url string) (int, time.Duration) {
	initialTime := time.Now()
	result := function(url)
	return result, time.Duration(time.Duration(time.Now().Sub(initialTime)).Milliseconds())
}
