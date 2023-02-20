/*
 * 1つのgoroutineのみが変数にアクセスできるように排他制御を行う
 * https://go-tour-jp.appspot.com/concurrency/9
 *
 * GO標準ライブラリでは排他制御をsync.MutexとLock(), Unlock()の2つのメソッドにて
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	// sync.Mutexの宣言を行う
	mu sync.Mutex
	v  map[string]int
}

// mainからgoroutineで実行される関数
func (c *SafeCounter) Inc(key string) {
	// ロックする
	c.mu.Lock()
	c.v[key]++
	// ロックを解除する
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	// ロックする
	c.mu.Lock()
	// mutexがunlockすることを保証させるためにdeferとして記述することもあります。
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
