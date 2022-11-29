package main

import (
	"flag"
	"fmt"
	"github.com/liwei1dao/lego/sys/proto"
	"net"
	"os"
	"time"
	"war-game/comm"
	"war-game/pb"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "3563", "port")

type DefMessage struct {
	ComId  uint16 //主Id
	MsgId  uint16 //次Id
	MsgLen uint32 //消息体长度
	Buffer []byte //消息体
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	//go handleWrite(conn, done)
	go handleRead(conn, done)
	go handleWriteHello(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
func handleWriteHello(conn net.Conn, done chan string) {
	if err := proto.OnInit(map[string]interface{}{"MsgProtoType": 1}); err != nil {
		panic(fmt.Sprintf("初始化proto失败 %v", err))
	}
	msg := proto.EncodeToMesage(comm.FighterComId, comm.FighterReq, &pb.DemoTestReq{Name: "czy"}) //&pb.DemoTestResp{}
	bytes := proto.EncodeToByte(msg)
	_, e := conn.Write([]byte(bytes))
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
	}

}
func handleWrite(conn net.Conn, done chan string) {
	if err := proto.OnInit(map[string]interface{}{"MsgProtoType": 1}); err != nil {
		panic(fmt.Sprintf("初始化proto失败 %v", err))
	}
	msg := proto.EncodeToMesage(0, 1, nil) //&pb.DemoTestResp{}
	bytes := proto.EncodeToByte(msg)

	for i := 100; i > 0; i-- {
		time.Sleep(time.Second * 1)
		_, e := conn.Write([]byte(bytes))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
		}
	}

	//fmt.Println(string(bytes))
	/*	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}*/
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"

}
