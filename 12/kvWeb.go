package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type myElement struct {
	Name    string
	SurName string
	Id      string
}

var DATA = make(map[string]myElement)
var DATAFILE = "./12/dataFile.gob"

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}

	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	defer loadFrom.Close()
	if err != nil {
		fmt.Println("Empty key/value stroe!")
		return err
	}
	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}

func ADD(k string, v myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = v
		return true
	}

	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) == nil {
		return false
	}
	delete(DATA, k)
	return true
}

func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if !ok {
		return nil
	}
	n := DATA[k]
	return &n
}

func CHANGE(k string, v myElement) bool {
	DATA[k] = v
	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("Key: %s, value: %v\n", k, v)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("./12/home.gohtml"))
	myT.ExecuteTemplate(w, "home.gohtml", nil)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, "<h1>Home sweet Home!</h1>")
	fmt.Fprintf(w, `<a href="/">Home sweet home!</a>`)
	fmt.Fprintf(w, `<a href="/list">List all elements!</a>`)
	fmt.Fprintf(w, `<a href="/change">Change an element!</a>`)
	fmt.Fprintf(w, `<a href="/insert">Insert an element!</a>`)
	fmt.Fprintf(w, "<ul>")
	for k, v := range DATA {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV store!")
	tmpl := template.Must(template.ParseFiles("./12/update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		SurName: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	CHANGE(key, n)
	err := save()
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element of the KV store!")
	tmpl := template.Must(template.ParseFiles("./12/insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		SurName: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	ADD(key, n)
	err := save()
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {
	_ = load()

	PORT := ":8001"

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)

	_ = http.ListenAndServe(PORT, nil)
}
