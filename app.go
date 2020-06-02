package main

import (
	"strconv"
	"strings"
)

func main() {
	//output := Slice("2 4 1 2 5 3-3 2 3 1")
	//output := Slice("1 1-1 1")
	//output := Slice("2 2 2-10 10 10")
	//output := Slice("1 1-100")
	//log.Println(output)
}

func Slice(data string) int {
	split := strings.Split(data, "-")
	splitX := strings.Split(split[0], " ")
	splitY := strings.Split(split[1], " ")

	var i int64
	var j int64
	var count int

	for i = 0; i <= int64(len(splitY)-1); i++ {
		limY, _ := strconv.ParseInt(splitY[i], 5, 0)
		for j = 0; j <= limY-1; j++ {
			limX, _ := strconv.ParseInt(splitX[j], 5, 0)
			if i < limX {
				count++
			}
		}
	}

	return count
}
