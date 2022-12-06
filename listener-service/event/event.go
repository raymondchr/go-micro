package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the exchange
		"topic",      //type
		true,         //durable?
		false,        //auto delete?
		false,        //internally?
		false,        //noWait?
		nil,          //arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"apa aja", //name
		false,
		false,
		true,
		false,
		nil,
	)
}
