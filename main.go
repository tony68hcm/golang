package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	fmt.Print("You live in " + city)
	time.Sleep(5)
	fmt.Println("test test")
}
