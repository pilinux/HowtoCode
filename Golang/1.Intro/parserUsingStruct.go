// continuation of parserUsingMap.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type parser struct {
	sum     map[string]int // total visits per domain
	domains []string       // slice to keep domain names in a sequence
	total   int            // total visits to all domain
	lines   int            // number of parsed line
}

func main() {
	p := parser{
		sum: make(map[string]int),
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		// split the string s around each instance of one or more consecutive white space
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v at line #%d\n", fields, p.lines)
			fmt.Println("existing...")
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %v at line #%d\n", fields[1], p.lines)
			fmt.Println("existing...")
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		p.sum[domain] += visits
		p.total += visits
	}

	// after "Domain", add spaces
	// before "Visits", add spaces
	fmt.Printf("%-20s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 35))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		visits := p.sum[domain]
		fmt.Printf("%-20s %10d\n", domain, visits)
	}

	fmt.Println(strings.Repeat("_", 35))
	fmt.Printf("%-20s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err: ", err)
		return
	}
}

/*
./parserUsingStruct < parserUsingStruct.log1.txt
./parserUsingStruct < parserUsingStruct.log2.txt
./parserUsingStruct < parserUsingStruct.log3.txt
./parserUsingStruct < parserUsingStruct.log4.txt
*/
