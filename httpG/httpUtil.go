package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	flagDebug = flag.Bool("debug", true, "Print message type")
)

func httpGet(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get error:", err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll error:", err)
		return err
	}

	fmt.Println(string(body))
	return nil
}

func main() {
	log.Println("in")
	flag.Parse()
	httpGet("http://pd1.pd.megvii-inc.com/videos/")

	log.Println("end")
}
