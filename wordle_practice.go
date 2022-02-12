package main

import (
	"bufio"
	"fmt"
	"os"
)

const WORDS_URL = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"

func main() {

	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')

	fmt.Println(s1, len(s1))
}
