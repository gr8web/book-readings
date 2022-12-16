package rgx

import (
	"fmt"
	"testing"
)

func TestMatchstar(t *testing.T) {
	fmt.Println(matchstar('a', "", "baaab"))
	fmt.Println(matchstar('a', "", "aaaa"))
}

func TestMatch(t *testing.T) {
	fmt.Println(Match("a*", "baaab"))
	fmt.Println(Match(".a*", "baaab"))
}
