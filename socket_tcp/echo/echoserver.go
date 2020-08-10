/*
 * echoサーバのサンプル
 * 
 * 引用元参考: https://qiita.com/kawasin73/items/3371d35166af733c2ce4
 */
package main

import (
    "log"
    "net"
)

func handleConnection(conn *net.TCPConn) {
    defer conn.Close()

    // readのbuffer sizeは4K
    buf := make([]byte, 4*1024)

    for {
        n, err := conn.Read(buf)
        if err != nil {
            if ne, ok := err.(net.Error); ok {
                switch {
                case ne.Temporary():
                    continue
                }
            }
            log.Println("Read", err)
            return
        }

        n, err = conn.Write(buf[:n])
        if err != nil {
            log.Println("Write", err)
            return
        }
    }
}

func handleListener(l *net.TCPListener) error {
    defer l.Close()
    for {
        log.Println("AcceptTCP Start")
        conn, err := l.AcceptTCP()
        if err != nil {
            if ne, ok := err.(net.Error); ok {
                if ne.Temporary() {
                    log.Println("AcceptTCP", err)
                    continue
                }
            }
            return err
        }

        // goroutineを利用
        go handleConnection(conn)
    }
}

func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:12345")
    if err != nil {
        log.Println("ResolveTCPAddr", err)
        return
    }

    l, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        log.Println("ListenTCP", err)
        return
    }

    err = handleListener(l)
    if err != nil {
        log.Println("handleListener", err)
    }
}
