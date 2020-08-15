package main

import (
	_ "net/http/pprof"

	myip "github.com/claytonblythe/myip/myip"
)

func main() {
	myip.My_ip()
}
