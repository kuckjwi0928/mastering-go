package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	myUrl string
	delay = 1
	w     sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer w.Done()
	data := make(chan myData, 1)
	httpClient := &http.Client{}

	go func() {
		response, err := httpClient.Get(myUrl)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done():
		fmt.Println(c.Err())
	case ok := <-data:
		err := ok.err
		if err != nil {
			fmt.Println("error select:", err)
			return err
		}
		resp := ok.r
		defer resp.Body.Close()
		realHttpData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", realHttpData)
	}

	return nil
}

func main() {
	myUrl = "https://www.naver.com"

	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s \n", myUrl)
	w.Add(1)
	go connect(c)
	w.Wait()
	fmt.Println("Exiting...")
}
