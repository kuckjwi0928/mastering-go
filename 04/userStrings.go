package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Println

func main() {
	f(s.ToUpper("Hello there!"))
	f(s.ToLower("Hello THERE"))
	f(s.Title("tHis wiLL be A title"))
	f("EqualFold:", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold:", s.EqualFold("Mihalis", "MIHAli"))
	f("Prefix:", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix:", s.HasSuffix("Mihalis", "lis"))
	f("Index:", s.Index("Mihalis", "ha"))
	f("Index:", s.Index("Mihalis", "Ha"))
	f("Index:", s.Index("Mihalis", "i"))
	f("Index:", s.Index("Mihalis", "I"))
	f("Repeat:", s.Repeat("ab", 5))
	f("TrimSpace", s.TrimSpace(" \tThis is a line."))
	f("TrimLeft", s.TrimLeft(" \tThis is a\t line.", "\n\t "))
	f("TrimLeft", s.TrimRight(" \tThis is a\t line.", "\n\t "))
	f("Compare:", s.Compare("Mihalis", "MIHALIS"))
	f("Compare:", s.Compare("Mihalis", "Mihalis"))
	f("Compare:", s.Compare("mihalis", "MIHalis"))
	f("Fields:", s.Fields("This is a string!"))
	f("Fields:", s.Fields("Thisis\na\tstring!"))
	f(s.Split("abcd efg", ""))
	f(s.Replace("abcd efg", "", "_", -1))
	f(s.Replace("abcd efg", "", "_", 4))
	f(s.Replace("abcd efg", "", "_", 2))
	f(s.Join([]string{"Line 1", "Line2", "Line3"}, ","))
	f(s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f(s.TrimFunc("123 abc ABC \t", trimFunction))
}
