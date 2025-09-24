package pcre

import "fmt"

type Regexp struct {
}

func Compile(pattern string, flags int) (Regexp, error) {
	return Regexp{}, fmt.Errorf("pcre.Compile not supported on windows")
}
