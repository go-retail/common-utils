package rabbit

import (
	"fmt"

	utils "github.com/go-retail/common-utils/pkg/log"
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
	utils.FailOnError(err, "Failed to Connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to Open The Channel")

	//TODO Deliver to Exchange not a Queue
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	utils.FailOnError(err, "Unable to Declare a Queue")
	Rmq = RMQ{conn, ch, &q}
}

// Find a place to put this code ...

// defer rmq.Connection.Close()
// defer rmq.Channel.Close()
