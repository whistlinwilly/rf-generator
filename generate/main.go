package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dic, err := os.Open("../dic/en_US.dic")
	if err != nil {
		log.Fatal(err)
	}
	defer dic.Close()

	r, err := os.Create("../dic/r.dic")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	f, err := os.Create("../dic/f.dic")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(dic)
	i := 0
	for scanner.Scan() {
		i++
		s := scanner.Text()
		if s[0] == 'R' || s[0] == 'r' {
			_, err := r.WriteString(strings.Split(s, "/")[0] + "\n")
			if err != nil {
				panic(err)
			}
		}
		if s[0] == 'F' || s[0] == 'f' {
			_, err := f.WriteString(strings.Split(s, "/")[0] + "\n")
			if err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
