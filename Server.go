package main

import (
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn){
	defer conn.Close()
	ip:=conn.RemoteAddr().String()
	fmt.Printf("new client:%s\n",ip)
	for {
        buff:=make([]byte,128)
        bufflen,err:=conn.Read(buff)
		if err==io.EOF {
			fmt.Printf("%s close\n",ip)
			return
		}
		if err!=nil {
			fmt.Printf("%s read err:%s\n",ip,err.Error())
			return
		}
		if bufflen>0 {
			fmt.Printf("%s recv:%s\n",ip,string(buff[0:bufflen]))
			conn.Write(buff[0:bufflen])
		}
	}
}

func main()  {
	listen,err:=net.Listen("tcp","127.0.0.1:9999")
	if err!=nil {
		fmt.Printf("tcp listen err:%s\n",err.Error())
		return
	}
	defer  listen.Close()
	fmt.Println("tcp lisen 127.0.0.1:9999'")
	for {
		conn,err:=listen.Accept()
		if err!=nil {
			continue
		}
		go handle(conn)
	}
}