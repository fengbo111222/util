package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
  	producer_test()
}

func producer_test() {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{"192.168.66.34:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "topic002",
		//Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message by fb111"

		fmt.Print(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

	for{
		select {
		case  suc:=<-producer.Successes():
			fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
		case fail := <-producer.Errors():
			fmt.Println("err: ", fail.Err)
		}
	}

}
