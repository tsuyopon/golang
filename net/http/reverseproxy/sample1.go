//  https://go.dev/play/p/YhR1A08G0n
// リバースプロキシのサンプル
// 80番ポートで立ち上がり指定したサイトにリバースプロキシするだけのプログラム
package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

type customTransport struct {
	http.Transport
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println("custom RoundTrip")
	res, err := http.DefaultTransport.RoundTrip(req)
	return res, err
}

func ReverseProxy(target string) *httputil.ReverseProxy {
	slice := strings.Split(target, "://")
	if len(slice) != 2 {
		log.Fatalf("Invalid target URI")
	}
	scheme := slice[0]
	host := slice[1]
	director := func(req *http.Request) {
		log.Printf("custom Director")
		req.URL.Scheme = scheme
		req.URL.Host = host
	}
	dial := func(network, addr string) (net.Conn, error) {
		log.Println("custom Dial")
		return net.Dial(network, addr)
	}

	//transport := &http.Transport{Dial: dial,}
	/* prints:
	   2016/02/18 13:35:34 (func(string, string) (net.Conn, error))(0x401810)
	   2016/02/18 13:35:34 &http.Transport{idleMu:sync.Mutex{state:0, sema:0x0}, wantIdle:false, idleConn:map[http.connectMethodKey][]*http.persistConn(nil), idleConnCh:map[http.connectMethodKey]chan *http.persistConn(nil), reqMu:sync.Mutex{state:0, sema:0x0}, reqCanceler:map[*http.Request]func()(nil), altMu:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:0, readerWait:0}, altProto:map[string]http.RoundTripper(nil), Proxy:(func(*http.Request) (*url.URL, error))(nil), Dial:(func(string, string) (net.Conn, error))(0x401810), DialTLS:(func(string, string) (net.Conn, error))(nil), TLSClientConfig:(*tls.Config)(nil), TLSHandshakeTimeout:0, DisableKeepAlives:false, DisableCompression:false, MaxIdleConnsPerHost:0, ResponseHeaderTimeout:0, ExpectContinueTimeout:0, TLSNextProto:map[string]func(string, *tls.Conn) http.RoundTripper(nil), nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0}, done:0x0}, h2transport:(*http.http2Transport)(nil)}
	   2016/02/18 13:35:36 custom Director
	   2016/02/18 13:35:36 custom Dial
	*/

	transport := &customTransport{http.Transport{Dial: dial}}
	/* prints:
	   2016/02/18 13:36:45 (func(string, string) (net.Conn, error))(0x401950)
	   2016/02/18 13:36:45 &main.customTransport{Transport:http.Transport{idleMu:sync.Mutex{state:0, sema:0x0}, wantIdle:false, idleConn:map[http.connectMethodKey][]*http.persistConn(nil), idleConnCh:map[http.connectMethodKey]chan *http.persistConn(nil), reqMu:sync.Mutex{state:0, sema:0x0}, reqCanceler:map[*http.Request]func()(nil), altMu:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:0, readerWait:0}, altProto:map[string]http.RoundTripper(nil), Proxy:(func(*http.Request) (*url.URL, error))(nil), Dial:(func(string, string) (net.Conn, error))(0x401950), DialTLS:(func(string, string) (net.Conn, error))(nil), TLSClientConfig:(*tls.Config)(nil), TLSHandshakeTimeout:0, DisableKeepAlives:false, DisableCompression:false, MaxIdleConnsPerHost:0, ResponseHeaderTimeout:0, ExpectContinueTimeout:0, TLSNextProto:map[string]func(string, *tls.Conn) http.RoundTripper(nil), nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0}, done:0x0}, h2transport:(*http.http2Transport)(nil)}}
	   2016/02/18 13:36:46 custom Director
	   2016/02/18 13:36:46 custom RoundTrip
	*/
	log.Printf("%#v", transport.Dial)
	res := &httputil.ReverseProxy{Director: director, Transport: transport}
	return res
}

func main() {
	proxy := ReverseProxy("http://yahoo.co.jp")
	log.Printf("%#v", proxy.Transport)
	mux := http.NewServeMux()
	mux.Handle("/", proxy)
	http.ListenAndServe(":80", mux)
}
