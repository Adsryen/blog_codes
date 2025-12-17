.PHONY: all

all:
	go run main.go

create_topic:
	docker exec rmqbroker sh \
		-c "mqadmin updateTopic -n rmqnamesrv:9876 -c DefaultCluster -t test_topic"