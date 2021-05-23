package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

//kafka写日志的模块

var (
	client sarama.SyncProducer //连接kafka的Producer
)

//初始化client
func Init(addrs []string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return err
	}
	//defer client.Close()
	return nil
}

func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("send message succuse,pid:%v,offset:%v\n", pid, offset)
}
