package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"net/http/cookiejar"
	io "io/ioutil"
	"bufio"
	"strings"
	"strconv"
	"sort"
)
var client http.Client
func init() {
    jar, err := cookiejar.New(nil)
    if err != nil {
        log.Fatalf("Got error while creating cookie jar %s", err.Error())
    }
    client = http.Client{
        Jar: jar,
    }
}
func main() {
    cookie := &http.Cookie{
        Name:   "session",
        Value:  os.Getenv("API_TOKEN"),
        MaxAge: 300,
    }
    req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/1/input", nil)
    if err != nil {
        log.Fatalf("Got error %s", err.Error())
    }
    req.AddCookie(cookie)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error occured. Error is: %s", err.Error())
    }
    defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		findMaxCapacity(capacityPerElf(bodyString))
		findSumTopThree(listOfCapacity(bodyString))
	}
}
func listOfCapacity(data string) []int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	a := make([]int, 0)
	elf := 0
	capacity := 0
	for scanner.Scan() {
		line := scanner.Text()
    	if line == "" {
			a = append(a, capacity)
			capacity = 0
			elf += 1
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			capacity += value
			
		}
	}
	sort.Ints(a[:])
	return a
}
func capacityPerElf(data string) map[int]int {
	scanner := bufio.NewScanner(strings.NewReader(data))
	m := make(map[int]int)
	elf := 0
	capacity := 0
	for scanner.Scan() {
		line := scanner.Text()
    	if line == "" {
			m[elf] = capacity
			capacity = 0
			elf += 1
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			capacity += value
			
		}
	}
	return m
}

func findMaxCapacity(m map[int]int) int {
	max := 0
	for k := range m {
		if m[k] > max {
			max = m[k]
		}
	}
	return max
}

func findSumTopThree(a []int) int {
	sumTopThree := 0
	for i := 1; i <= 3; i++ {
		sumTopThree +=  a[len(a) - i]
	}
	fmt.Print(sumTopThree)
	return sumTopThree
}