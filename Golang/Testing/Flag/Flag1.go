package main

import (
	"flag"
	"fmt"
	"strings"
)

type sliceValue []string

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's not my business"
}

func main() {
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I love Golang")
	fmt.Println(languages)
	flag.Parse()
	fmt.Println(languages)
}
