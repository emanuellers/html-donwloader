package main

func main() {
	macthing := Request{}

	macthing.GetPageHtmlToFile("https://pkg.go.dev/regexp#Regexp", "teste/manu.html")
	buffer, _ := macthing.GetBufferResponse("https://pkg.go.dev/regexp#Regexp")
	macthing.GetHiperlinks(buffer, 0)
}
