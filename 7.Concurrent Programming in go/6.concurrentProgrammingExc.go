package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		fmt.Printf("%T\n", line)
		if err == io.EOF {
			fmt.Println(err)
			return sum
		}

		if err != nil {
			fmt.Println("error", err)
		}
		num, err := strconv.Atoi(line[:len(line)-1])
		fmt.Println("after conversion", num)
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt"}
	sum := 0

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		rd := bufio.NewReader(file)
		calculate := func() {
			fileSum := sumFile(*rd)
			sum += fileSum
		}
		go calculate()
	}
	time.Sleep(300 * time.Millisecond)
	fmt.Println(sum)
}
