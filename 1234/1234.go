package main

import(
	"bufio"
	"os"
	"fmt"
	"unicode"
)

func main () {
	
	var (
		in, out string
		chars []rune
	)

	reader := bufio.NewReader(os.Stdin)

	in, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	chars = []rune(in)

	for y := range len(chars){
		if y%2 == 0 {
			out += string(unicode.ToUpper(chars[y]))
		} else {
			out += string(chars[y])
		}
	}

	fmt.Print(out)
	}

