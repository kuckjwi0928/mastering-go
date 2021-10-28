package main

import (
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	tFile := "/Volumes/kuckjwi/source/mastering-go/06/tFile"
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []Entry

	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Entries = append(Entries, temp)
		}
	}

	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
