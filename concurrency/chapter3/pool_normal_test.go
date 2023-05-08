package chapter3

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen:%v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			//connectToService()
			//fmt.Fprintln(conn, "")
			conn.Close()
		}

	}()

	return &wg
}

//	测试命令为:go test pool_normal_test.go -benchtime=10s -bench=.

//	测试结果为:
//	goos: darwin
//	goarch: amd64
//	BenchmarkNetworkRequest-12    	      10	1003252912 ns/op
//	PASS
//	ok  	command-line-arguments	11.048s

// 添加同步池之后
// 测试结果为:
// goos: darwin
// goarch: amd64
// BenchmarkNetworkRequest-12    	    5508	   2765755 ns/op
// PASS
// ok  	command-line-arguments	37.806s
func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host:%v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{New: connectToService}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}
