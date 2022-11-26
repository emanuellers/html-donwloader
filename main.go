package main

import "log"

func main() {
	macthing := Html{}
	buffer, _ := macthing.GetBufferResponse("https://google.com")
	matches := macthing.GetHiperlinks(buffer, 0)

	test, duration := Timer(macthing.CheckFuntionalUrls, matches[0])
	log.Println(test, duration)

}
