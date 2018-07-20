package nsq

import (
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/fengyfei/gu/libs/nsq/nsq"
)

type Conf struct{ Config *nsq.Config }

type Producer struct{ *nsq.Producer }

type Consumer struct{ *nsq.Consumer }

func (c *Conf) NewConfig(setConf ...func(*nsq.Config)) *nsq.Config {
	conf := nsq.NewConfig()
	for _, set := range setConf {
		set(conf)
	}
	c.Config = conf
	return conf
}

func (c *Conf) NewProducer(addr string) (*Producer, error) {
	p, err := nsq.NewProducer(addr, c.Config)
	if err != nil {
		return nil, err
	}
	err = p.Ping()
	if err != nil {
		return nil, err
	}
	return &Producer{p}, nil
}

func main() {
	conf := Conf{nsq.NewConfig()}
	p, err := conf.NewProducer("127.0.0.1:4150")
	if err != nil {
		log.Println(err)
	}
	e := p.NewTopic("Topic-01", []byte("ssdfjnaowihgajndsk"))
	if e != nil {
		log.Println(err)
	}

}

func (p *Producer) NewTopic(topic string, massage []byte) interface{} {
	return p.Publish(topic, massage)
}

func (c *Conf) NewConsumer() {

}
