// Regular Expression
// Example:
// Valid characters:
// A-Z, a-z, 0-9, -, _
// Starting character must be A-Z or a-z

package main

import (
	"fmt"
	"regexp"
)

func main() {
	isAlphanumericUnderScoreHyphen := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*$`).MatchString

	fmt.Println(isAlphanumericUnderScoreHyphen("AleX1_6-_"))
	fmt.Println(isAlphanumericUnderScoreHyphen("*leX1_6-_"))
	fmt.Println(isAlphanumericUnderScoreHyphen("_leX1_6-6"))
	fmt.Println(isAlphanumericUnderScoreHyphen("-leX1_6-6"))
	fmt.Println(isAlphanumericUnderScoreHyphen("0-eX1_6-6"))
	fmt.Println(isAlphanumericUnderScoreHyphen("ü_eX1_6-6"))
}

/*
Output:

true
false
false
false
false
false

*/
