package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func server(wg *sync.WaitGroup) {
	server, err := net.Listen("tcp", ":8080")
	defer wg.Done()
	if err != nil {
		fmt.Println("error listen")
		return
	} else {
		for i := 0; i < 10; i++ {
			_, Aerr := server.Accept()
			fmt.Printf("client connecting %d\n", i)
			if Aerr == nil {
				msg := string("hello client ")
				msg = msg + strconv.Itoa(i)
				// _, err = connection.Write([]byte(msg + strconv.Itoa(i)))
			}
		}
	}
}

func client(wg *sync.WaitGroup) {
	defer wg.Done()
	var wg1 sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			defer wg1.Done()
			conn, err := net.Dial("tcp", "127.0.0.1:8080")
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			if err != nil {
				fmt.Println(err)
				return
			} else {
				buf := make([]byte, 1024)
				if length, err := conn.Read(buf); err == nil {
					if length > 0 {
						fmt.Printf("recv:%s\n", buf[0:length])
					} else {
						fmt.Println("nothing")
					}
				} else {
					fmt.Println(err)
				}
				conn.Close()
				fmt.Println("disconnect")
			}
		}()
		wg1.Add(1)
	}
	wg1.Wait()
}

func main() {
	var wg sync.WaitGroup
	go server(&wg)
	wg.Add(1)
	go client(&wg)
	wg.Add(1)
	wg.Wait()
	fmt.Println("all go routines finished executing")
}
