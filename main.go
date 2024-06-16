package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var arr []string
	f, _ := os.OpenFile("./read.txt", os.O_RDONLY, 0777)
	defer f.Close()

	fileReader := bufio.NewReader(f)

	file, _ := os.OpenFile("./write.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		arr = append(arr, string(line))
	}

	for i := 0; i < len(arr); i++ {
		primer := regexp.MustCompile(`^[0-9]+[+\-*/][0-9]+=\?$`)
		isMatch := primer.MatchString(arr[i])
		if isMatch {
			primer = regexp.MustCompile(`^([0-9]+)([+\-*/])([0-9]+)=\?$`)
			matches := primer.FindStringSubmatch(arr[i])

			num1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			num2, _ := strconv.Atoi(matches[3])

			var result int
			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}
			output := fmt.Sprintf("%s%d\n", arr[i][:len(arr[i])-1], result)
			writer.WriteString(output)
		}
	}
}
