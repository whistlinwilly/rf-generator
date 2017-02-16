package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("./dic/f.dic")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r, err := os.Open("./dic/r.dic")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	rCount := getLineCount(r)
	fCount := getLineCount(f)

	rand.Seed(time.Now().Unix())
	best := 0
	rf := ""
	for i := 0; i < 15; i++ {
		randR := rand.Intn(rCount)
		randF := rand.Intn(fCount)
		rWord := getLine(r, randR)
		fWord := getLine(f, randF)
		cur := getSearchResults(rWord, fWord)
		fmt.Printf("%v %v (%v)\n", rWord, fWord, cur)
		if cur > best {
			best = cur
			rf = rWord + " " + fWord
		}
	}
	fmt.Println("CHOOSE", rf)
}

func getLineCount(f *os.File) int {
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i++
	}
	return i
}

func getLine(f *os.File, n int) string {
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		if i == n {
			return scanner.Text()
		}
		i++
	}
	return ""
}

func getSearchResults(a, b string) int {
	re, _ := regexp.Compile("About .* results")
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.google.com/search?q=%22"+a+"%20"+b+"%22", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	results := re.FindString(string(body))
	if results == "" {
		return 0
	}
	s, err := strconv.Atoi(strings.Replace(strings.Split(results, " ")[1], ",", "", -1))
	if err != nil {
		return 0
	}
	return s
}
