package main

import (
	"fmt"
	"net"
	"os"
)



func get_internal() string{
	var addr_s string
	addrs,err:=net.InterfaceAddrs()
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	for _,a:=range addrs{
		if ipnet,ok :=a.(*net.IPNet);ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() !=nil{
				//fmt.Println(ipnet.IP.String())
				addr_s = ipnet.IP.String()
				break
			}
		}
	}
	return addr_s
}

func main()  {
	fmt.Println(get_internal())
}
