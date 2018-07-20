package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

// Producer 生产者
func Producer() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewProducer", err)
		panic(err)
	}

	for {
		if err := producer.Publish("test", []byte(fmt.Sprintf("Hello World "))); err != nil {
			fmt.Println("Publish", err)
			panic(err)
		}

		time.Sleep(time.Second * 3)

	}
}

//消费者1
func Consumer1() {
	consumer, err := nsq.NewConsumer("test", "channel-a", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer fail:", err)
		panic(err)
	}

	// 添加消息处理的具体实现
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("Consumer1:", string(msg.Body))
		return nil
	}))

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd fail:", err)
		panic(err)
	}
}

//消费者2
func Consumer2() {
	consumer, err := nsq.NewConsumer("test", "channel-a", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer fail:", err)
		panic(err)
	}

	// 添加消息处理的具体实现
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("Consumer2:", string(msg.Body))
		return nil
	}))

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd fail:", err)
		panic(err)
	}
}

func main()  {
	Consumer1()
	Consumer2()
	Producer()
}