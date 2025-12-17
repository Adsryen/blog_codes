package main

import (
	"fmt"
	"net"
	"time"
)

func MainSimple() {
	fmt.Println("Starting Consumer...")
	go RunConsumer()

	fmt.Println("Waiting for consumer to start...")
	time.Sleep(5 * time.Second)

	fmt.Println("Starting Producer...")
	RunProducer()

	// Keep main alive to let consumer receive messages
	select {}
}

func ResolveNameSrv(addr string) string {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Printf("ResolveTCPAddr error: %s\n", err)
		return addr
	}
	// Verify if IP is valid, though ResolveTCPAddr usually returns valid IP in String()
	if tcpAddr.IP == nil {
		return addr
	}
	return tcpAddr.String()
}
