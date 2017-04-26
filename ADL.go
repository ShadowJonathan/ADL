package main

import (
	"io/ioutil"
	"net/url"
	"strings"
)

func main() {
	go dl()
	b, err := ioutil.ReadFile("queue.txt")
	if err == nil {
		things := strings.Split(string(b), "\n")
		for _, s := range things {
			go func() { dlchan <- s }()
		}
	}
	<-finished
}

func dl() {
	for _, s := dlchan {
		down(s)
	}
}

func down(s string) error {

}

var dlchan = make(chan string, 100)
var finished = make(chan bool)

type download struct {
	args []dlarg
	url  *url.URL
}

type dlarg struct {
	method string
	value  string
}
