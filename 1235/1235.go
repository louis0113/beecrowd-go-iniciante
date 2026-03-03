package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		n, ini1, ini2, fi1         int
		msg, m1, m2, mt, arrin, nm []string
		chars                      [][]rune
		tm1, tm2, in, ns           string
		err                        error
	)

	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)

	for x := range n {
		in, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		in = strings.TrimSpace(in)
		arrin = append(arrin, in)
		msg = append(msg, arrin[x])
		chars = append(chars, []rune(msg[x]))
	}

	for y := range n {
		ini1, ini2, fi1 = len(chars[y])/2-1, len(chars[y])-1, len(chars[y])/2

		for i := ini1; i >= 0; i-- {
			tm1 += string(chars[y][i])
		}
		m1 = append(m1, tm1)

		for j := ini2; j >= fi1; j-- {
			tm2 += string(chars[y][j])
		}

		m2 = append(m2, tm2)
		ns = m1[y] + m2[y]
		nm = append(nm, ns)
		mt = append(mt, nm[y])
		tm1, tm2 = "", ""
	}

	for a := range n {
		fmt.Println(mt[a])
	}

}
