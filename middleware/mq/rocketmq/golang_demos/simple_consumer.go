package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/liuliqiang/log4go"
)

func RunConsumer() {
	nameSrv := os.Getenv("NAMESRV_ADDR")
	if nameSrv == "" {
		nameSrv = "127.0.0.1:9876"
	}
	// Must resolve to IP because rocketmq-client-go's primitive.NewPassthroughResolver requires IP
	nameSrv = ResolveNameSrv(nameSrv)
	log4go.Info("nameSrv: %s", nameSrv)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test_group"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{nameSrv})),
	)
	if err != nil {
		fmt.Printf("create consumer error: %s\n", err)
		os.Exit(1)
	}
	err = c.Subscribe("test_topic", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
