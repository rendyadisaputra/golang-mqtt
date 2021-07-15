package main
 
import (
	"fmt"
	"log"
	"net/url"
	// "os"
	"time"
 
	mqtt "github.com/eclipse/paho.mqtt.golang"
)
 
func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
 
func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}
 
func listen(uri *url.URL, topic string) {
	client := connect("CID:sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		
		fmt.Printf("* Tictoc Topic [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
 
func main() {
	uri, err := url.Parse("mqtt://testbroh:pass@localhost:1883/testtopic/#")
	// uri = "tcp://localhost:1883"
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	log.Printf(" [x] %s", topic)
	if topic == "" {
		topic = "test"
	}
 
	go listen(uri, topic)
 
	// client := connect("pub", uri)
	// timer := time.NewTicker(3 * time.Second)
	wait := make(chan bool)
	go func() {
		log.Printf("waiting for messages")
		
	}()
    <-wait
}
