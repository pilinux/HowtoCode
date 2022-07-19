package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum map[string]int

		// slice to keep domain names in a sequence
		domains []string

		total int
		lines int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		// split the string s around each instance of one or more consecutive white space
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v at line #%d\n", fields, lines)
			fmt.Println("existing...")
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %v at line #%d\n", fields[1], lines)
			fmt.Println("existing...")
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		sum[domain] += visits
		total += visits
	}

	// after "Domain", add spaces
	// before "Visits", add spaces
	fmt.Printf("%-20s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 35))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-20s %10d\n", domain, visits)
	}

	fmt.Println(strings.Repeat("_", 35))
	fmt.Printf("%-20s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err: ", err)
		return
	}
}

/*
./parserUsingMap < parserUsingMap.log1.txt
./parserUsingMap < parserUsingMap.log2.txt
./parserUsingMap < parserUsingMap.log3.txt
./parserUsingMap < parserUsingMap.log4.txt
*/
