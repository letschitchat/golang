package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) { /*传参conn*/
	fmt.Fprintf(conn, "%s", time.Now().String()) /*取现在时间变成string传入到conn*/
	conn.Close()                                 /*关闭连接*/
}
func main() { /*main函数*/
	l, err := net.Listen("tcp", ":8080") /*用net.listen监听8080端口*/
	if err != nil {                      /*查看有没有错误*/
		log.Fatal(err) /*有错误就logfatal*/
	}
	for { /*for循环*/
		conn, err := l.Accept() /*l是listen的socket，获取l，接受过来的连接*/
		if err != nil {         /*判断是否有错*/
			log.Fatal(err) /*有错打日志*/
		}
		go handle(conn)
	}
}

//进程，线程，协程（由重到轻）
//进程，协程，线程（历史出现次序）
