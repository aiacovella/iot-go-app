package mqttclient


import (
	"fmt"
	"github.com/goiiot/libmqtt"
	"log"
	"time"
)

//type MqttClient struct {}

func mqttInit()(*libmqtt.AsyncClient) {

	fmt.Println("connecting to MQTT server")

	//fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))


	// Create a client and enable auto reconnect when connection lost
	// We primarily use `RegexRouter` for client
	client, err := libmqtt.NewClient(
		// server address(es)
		libmqtt.WithServer("localhost:1883"),
		// enable keepalive (10s interval) with 20% tolerance
		libmqtt.WithKeepalive(10, 1.2),
		// enable auto reconnect and set backoff strategy
		libmqtt.WithAutoReconnect(true),
		libmqtt.WithBackoffStrategy(time.Second, 5*time.Second, 1.2),
		// use RegexRouter for topic routing if not specified
		// will use TextRouter, which will match full text
		libmqtt.WithRouter(libmqtt.NewRegexRouter()),
	)

	if err != nil {
		// handle client creation error
		panic("create mqtt client failed")

	}else{

		log.Println("Mqtt client connected")
		return client

	}




}



func MqttConnect()(*libmqtt.AsyncClient) {

	client := mqttInit()

	client.Connect(func(server string, code byte, err error) {
		if err != nil {
			// failed
			panic(err)
		}

		if code != libmqtt.CodeSuccess {
			// server rejected or in error
			panic(code)
		}

		// success
		// you are now connected to the `server`
		// (the `server` is one of your provided `servers` when create the client)
		// start your business logic here or send a signal to your logic to start

		// subscribe some topic(s)
		client.Subscribe([]*libmqtt.Topic{
			{Name: "foo"},
			{Name: "bar", Qos: libmqtt.Qos1},
		}...)

		// publish some topic message(s)
		client.Publish([]*libmqtt.PublishPacket{
			{TopicName: "foo", Payload: []byte("bar"), Qos: libmqtt.Qos0},
			{TopicName: "bar", Payload: []byte("foo"), Qos: libmqtt.Qos1},
		}...)


	})

	return client

}