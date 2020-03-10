package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://ryft-public-sample-data.s3.amazonaws.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	texto := string(body)

	s := strings.Split(texto, "<Key>")

	x := strings.Split(s[0], "<Name>")
	y := strings.Split(x[1], "</Name>")

	bucketName := y[0]

	var keys []string
	var ext []string

	for _, value := range s {
		z := strings.Split(value, "</Key>")
		keys = append(keys, z[0])
	}

	var objects int = 0
	dir := make(map[string]bool)

	for _, value := range keys {

		if strings.Contains(value, "/") {
			temp2 := strings.Split(value, "/")
			dir[temp2[0]] = true
		}
		if strings.Contains(value, ".") {
			objects++
			temp := strings.Split(value, ".")
			ext = append(ext, temp[len(temp)-1])
		}

	}

	counts := make(map[string]int)
	for i, word := range ext {
		if i > 0 {
			_, ok := counts[word]
			if ok {
				counts[word]++
			} else {
				counts[word] = 1
			}
		}
	}

	fmt.Println("AWS S3 Explorer")
	fmt.Println("Bucket Name: ", bucketName)
	fmt.Println("Number of objects: ", objects)
	fmt.Println("Number of directories: ", len(dir))
	fmt.Print("Extensions: ")
	for index, element := range counts {
		fmt.Print(index, "(", element, ") ")
	}

}

