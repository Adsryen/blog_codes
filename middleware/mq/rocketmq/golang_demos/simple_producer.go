package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/liuliqiang/log4go"
)

func RunProducer() {
	nameSrv := os.Getenv("NAMESRV_ADDR")
	if nameSrv == "" {
		nameSrv = "127.0.0.1:9876"
	}
	// Must resolve to IP because rocketmq-client-go's primitive.NewPassthroughResolver requires IP
	nameSrv = ResolveNameSrv(nameSrv)
	log4go.Info("nameSrv: %s", nameSrv)
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{nameSrv})),
		producer.WithRetry(2),
	)
	if err != nil {
		fmt.Printf("create producer error: %s\n", err)
		os.Exit(1)
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	topic := "test_topic"

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte(fmt.Sprintf("Hello RocketMQ %d", i)),
		}
		res, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
