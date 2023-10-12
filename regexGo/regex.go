package regexGo

import (
	"fmt"
	"regexp"
)

func MatchString(expression, input string) bool {
	reg, err := regexp.Compile(expression)
	if err != nil {
		return false
	}

	return reg.MatchString(input)
}

func TestifyRegexExpression() {
	regexExp := "[a-zA-Z0-9_-]{5}"
	fmt.Println(MatchString(regexExp, "a_-1A"))
	fmt.Println(MatchString(regexExp, "a_-1#"))
}
