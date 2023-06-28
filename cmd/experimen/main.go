package main

import (
	"fmt"
	"io"

	"golang.org/x/net/proxy"
)

func main() {
	d, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, nil)
	if err != nil {
		panic(err)
	}

	conn, err := d.Dial("tcp", "ifconfig.me:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET http://ifconfig.me/ip HTTP/1.0\r\n\r\n")

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		//fmt.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)

	}
	fmt.Println("total size:", len(buf), string(buf))
}
