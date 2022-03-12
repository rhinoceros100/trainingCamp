package biz

import "time"

type RunJob struct{}

func (rj *RunJob) Run() {
	go func() {
		//从kafka读数据
		rj.readFromKafka()

		//处理
		rj.dealData()

		//写入redis
		rj.writeRedis()
	}()
}

func (rj *RunJob) readFromKafka() {
	//TODO
	time.Sleep(time.Millisecond)
}

func (rj *RunJob) dealData() {
	//TODO
	time.Sleep(time.Millisecond)
}

func (rj *RunJob) writeRedis() {
	//TODO
	time.Sleep(time.Millisecond)
}
