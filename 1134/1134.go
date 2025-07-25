package main

import (
	"bufio"
	"fmt"
	"os"
)

const msg = "MUITO OBRIGADO"

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var x int

	contadores := make([]int, 4)

	for {
		_, err := fmt.Fscanln(reader, &x)
		if err != nil {
			break
		}

		if x >= 1 && x <= 3 {
			contadores[x]++
		} else if x == 4 {
			break
		}
	}

	produtos := []string{"", "Alcool", "Gasolina", "Diesel"}

	fmt.Fprintln(writer, msg)

	for i := 1; i < len(produtos); i++ {
		fmt.Fprintf(writer, "%s: %d\n", produtos[i], contadores[i])
	}
}
