package tests

import (
	"errors"
	"fmt"
	"testing"
	"testing/synctest"
	"time"
)

// simple read function, timeouts after 1m
type result struct{}

func Read(ch <-chan result) (result, error) {
	select {
	case r := <-ch:
		return r, nil
	case <-time.After(60 * time.Second):
		return result{}, errors.New("timeout after 60s")
	}
}

// usage
func main() {
	ch := make(chan result)
	go func() {
		ch <- result{}
	}()
	val, err := Read(ch)
	fmt.Printf("val=%v, err=%v\n", val, err)
}

func TestReadTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan result)
		_, err := Read(ch)
		if err == nil {
			t.Fatal("expected timeout")
		}
	})
}

func TestSleepFast(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		time.Sleep(1 * time.Hour) // no sleeping, wake up immediately!!! :D
	})
}

// from documentation, runs immediately, no need to wait for 2s
func TestTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now() // always midnight UTC 2000-01-01
		go func() {
			time.Sleep(1 * time.Second)
			t.Log(time.Since(start)) // always logs "1s"
		}()
		time.Sleep(2 * time.Second) // the goroutine above will run before this Sleep returns
		t.Log(time.Since(start))    // always logs "2s"
	})
}

func TestWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var innerStarted bool
		done := make(chan struct{})

		go func() {
			innerStarted = true
			time.Sleep(time.Second)
			close(done)
		}()

		// Wait for the inner goroutine to block on time.Sleep.
		synctest.Wait() // <-- try commenting this out
		// innerStarted is guaranteed to be true here.
		fmt.Printf("inner started: %v\n", innerStarted)

		<-done
	})
}
