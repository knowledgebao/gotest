package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	flagDebug = flag.Bool("debug", true, "Print message type")
)

func main() {
	flag.Parse()

	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	h := http.Header{"Origin": {"http://" + u.Host}}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), h)

	if err != nil {
		panic(err)
	}

	timeStart := time.Now()
	var numSum int64
	go func() {
		for {
			t, buf, err := conn.ReadMessage()
			if *flagDebug {
				log.Println("type:", t, len(buf), err)
				if len(buf) > 500 {
					numSum++
					if time.Since(timeStart)/time.Second > 0 {
						log.Println(numSum / int64(time.Since(timeStart)/time.Second))
					}
				}
				continue
			}
			os.Stdout.Write(buf)
		}
	}()

	b := bufio.NewReader(os.Stdin)
	for {
		line, err := b.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			break
		}
	}

	conn.Close()
}
