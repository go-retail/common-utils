package rabbit

import (
	"fmt"

	"github.com/go-retail/common-utils/pkg/logutils"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

//Rmq ..
var Rmq RMQ

//RMQ ..
type RMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

//RMQConfiguration ..
type RMQConfiguration struct {
	RabbitmqUsername string
	RabbitmqPassword string
	RabbitmqHost     string
}

var rmqConfigs RMQConfiguration

//InitRMQ ..
func InitRMQ() {
	rmqConfigs.RabbitmqUsername = viper.GetString("RabbitmqUsername")
	rmqConfigs.RabbitmqPassword = viper.GetString("RabbitmqPassword")
	rmqConfigs.RabbitmqHost = viper.GetString("RabbitmqHost")
	//TODO externalize RabbitMQ Port to config
	urlString := fmt.Sprintf(
		"amqp://%s:%s@%s:5672", rmqConfigs.RabbitmqUsername, rmqConfigs.RabbitmqPassword, rmqConfigs.RabbitmqHost)
	conn, err := amqp.Dial(urlString)
	logutils.FailOnError(err, "Failed to Connect to RabbitMQ")

	ch, err := conn.Channel()
	logutils.FailOnError(err, "Failed to Open The Channel")

	//TODO Deliver to Exchange not a Queue
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	logutils.FailOnError(err, "Unable to Declare a Queue")
	Rmq = RMQ{conn, ch, &q}
}

// Find a place to put this code ...

// defer rmq.Connection.Close()
// defer rmq.Channel.Close()

//PublishOnQueue ..
func PublishOnQueue(msg []byte, queueName string) error {
	if Rmq.Connection == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	// Declare a queue that will be created if not exists with some args
	queue, err := Rmq.Channel.QueueDeclare(
		queueName, // our queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	// Publishes a message onto the queue.
	err = Rmq.Channel.Publish(
		"",         // use the default exchange
		queue.Name, // routing key, e.g. our queue name
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg, // Our JSON body as []byte
		})
	fmt.Printf("A message was sent to queue %v: %v", queueName, msg)
	return err
}

//Publish ..
func Publish(msg []byte, exchangeName string, queueName string) error {
	if Rmq.Connection == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	// Declare a queue that will be created if not exists with some args
	queue, err := Rmq.Channel.QueueDeclare(
		queueName, // our queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	// Publishes a message onto the queue.
	err = Rmq.Channel.Publish(
		exchangeName, // use the exchange
		queue.Name,   // routing key, e.g. our queue name
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg, // Our JSON body as []byte
		})
	fmt.Printf("A message was sent to queue %v: %v", queueName, msg)
	return err
}
