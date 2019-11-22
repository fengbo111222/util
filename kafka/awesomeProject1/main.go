package main

import ("fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
)

func main() {
 	consumer_test()
}

func consumer_test() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"192.168.66.34:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("topic002", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for  {
		select {

		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}

	}


}
func consumer_test1(){

	addrs := []string{"192.168.66.34:9092"}
	//addrs := []string{"192.168.152.48:9092"}
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("topic003", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("Consumed message offset", msg.Offset)
			fmt.Println(string(msg.Value))
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	fmt.Println("Consumed:", consumed)

}
