package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

var ch=make(chan int,2)

func CreatConn(num int){
	conn,err:=net.Dial("tcp","127.0.0.1:9999")
	if err!=nil {
		fmt.Printf("dial err:%s\n",err.Error())
		return
	}
	defer conn.Close()
	str:="hello,i am"+strconv.Itoa(num)
	for {
		time.Sleep(time.Second*5)
		_,err=conn.Write([]byte(str))
		if err!=nil {
			fmt.Printf("write err:%s\n",err.Error())
		}else{
			fmt.Printf("send:%s\n",str)
		}

		buff:=make([]byte,128)
		datalen,err:=conn.Read(buff)
		if err==io.EOF {
			fmt.Println("server close")
			return
		}
		if err==nil&&datalen>0 {
			fmt.Printf("recv:%s\n",string(buff[:datalen]))
		}

		ch <- 1
	}
}

func main(){
	go CreatConn(1)
	<-ch
}

