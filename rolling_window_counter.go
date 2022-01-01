package main

import (
	"container/ring"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//计数桶
type CounterBucket struct {
	success     int64
	total       int64
	windowStart int64 //本桶计数时间
}

func (c *CounterBucket) AddRequest(isSuccess bool) {
	atomic.AddInt64(&c.total, 1)
	if isSuccess {
		atomic.AddInt64(&c.success, 1)
	}
}

func (c *CounterBucket) GetBucketCounter() (total, success int64) {
	return c.total, c.success
}

func (c *CounterBucket) Reset(windowStart int64) {
	atomic.StoreInt64(&c.total, 0)
	atomic.StoreInt64(&c.success, 0)
	atomic.StoreInt64(&c.windowStart, windowStart)
}

//滑动计数器
type RollingWindowCounter struct {
	r    *ring.Ring //环，存储计数数据
	head *ring.Ring
	tail *ring.Ring
}

func (rw *RollingWindowCounter) Init(size int) {
	rw.r = ring.New(size + 1)
	rw.head = rw.r
	rw.tail = rw.r

	for i := 0; i < rw.r.Len(); i++ {
		rw.head.Value = &CounterBucket{
			success:     0,
			total:       0,
			windowStart: time.Now().Unix(),
		}
		rw.head = rw.head.Next()
	}
	rw.head = rw.tail.Next()
}

func (rw *RollingWindowCounter) AddCount(isSuccess bool) error {
	if rw.r == nil {
		return errors.New("RollingWindowCounter not initial err")
	}

	cb, ok := rw.head.Value.(*CounterBucket)
	if !ok {
		fmt.Println("not ok")
		return errors.New("RollingWindowCounter type err")
	}
	timeNow := time.Now().Unix()
	if timeNow == cb.windowStart {
		cb.AddRequest(isSuccess)
		return nil
	}

	//rw.head所在的桶计数时间与现在不一致，寻找与当前计数时间一致的桶
	oldWindowStart := cb.windowStart
	diff := timeNow - oldWindowStart
	for i := int64(1); i <= diff; i++ {
		rw.head = rw.head.Next()
		if rw.head == rw.tail {
			rw.tail = rw.tail.Next()
		}

		cb2, ok := rw.head.Value.(*CounterBucket)
		if !ok {
			return errors.New("RollingWindowCounter type err")
		}
		cb2.Reset(oldWindowStart + i)
		if cb2.windowStart == timeNow {
			cb2.AddRequest(isSuccess)
			return nil
		}
	}

	return nil
}

func (rw *RollingWindowCounter) Sum() (total, success int64) {
	total, success = 0, 0
	for r := rw.tail; r != rw.head; r = r.Next() {
		if cb, ok := r.Value.(*CounterBucket); ok {
			//fmt.Printf("total:%d, success:%d, windowStart:%d\n", cb.total, cb.success, cb.windowStart)
			total += cb.total
			success += cb.success
		}
	}
	return
}

func main() {
	rollingWindowCounter := &RollingWindowCounter{}
	rollingWindowCounter.Init(10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wg := sync.WaitGroup{}
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 4000; i++ {
				random := r.Intn(10)
				isSuccess := false
				if random < 8 {
					isSuccess = true
				}
				_ = rollingWindowCounter.AddCount(isSuccess)
				time.Sleep(time.Millisecond * time.Duration(random))
			}
		}()
	}

	wg.Wait()
	fmt.Println(time.Now())
	total, success := rollingWindowCounter.Sum()
	fmt.Printf("total:%d, success:%d\n", total, success)
}
