// continuation of parserUsingFunc.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result // total visits per domain
	domains []string          // slice to keep domain names in a sequence
	total   int               // total visits to all domain
	lines   int               // number of parsed line
}

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		parsed, err := parse(&p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		update(&p, parsed)
	}

	// after "Domain", add spaces
	// before "Visits", add spaces
	fmt.Printf("%-20s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 35))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-20s %10d\n", domain, parsed.visits)
	}

	fmt.Println(strings.Repeat("_", 35))
	fmt.Printf("%-20s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err: ", err)
		return
	}
}

// constructor function
// constructs the initial value and returns it
func newParser() parser {
	return parser{
		sum: make(map[string]result),
	}
}

func parse(p *parser, line string) (result, error) {
	var (
		parsed result
		err    error
	)

	// split the string s around each instance of one or more consecutive white space
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v at line #%d", fields, p.lines)
		return parsed, err
	}

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %v at line #%d", fields[1], p.lines)
		return parsed, err
	}

	return parsed, err
}

func update(p *parser, parsed result) {
	domain, visits := parsed.domain, parsed.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// a value can be assigned to addressable elements
	// map elements are not addressable
	// hence, this won't work
	// p.sum[domain].visits += visits
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits, // current visits + old visits
	}

	p.total += visits
}

/*
./parserUsingPointer < parserUsingStruct.log1.txt
./parserUsingPointer < parserUsingStruct.log2.txt
./parserUsingPointer < parserUsingStruct.log3.txt
./parserUsingPointer < parserUsingStruct.log4.txt
*/
