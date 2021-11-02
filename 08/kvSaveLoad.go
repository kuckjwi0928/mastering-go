package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

const DATAFILE = "./08/dataFile.gob"

var DATA = make(map[string]string)

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("cannot create", DATAFILE)
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
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}
	decoder := gob.NewDecoder(loadFrom)
	_ = decoder.Decode(&DATA)
	return nil
}

func main() {
	DATA["name"] = "kuckjwi"
	_ = save()
	_ = load()
	fmt.Println(DATA)
}
