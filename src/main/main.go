package main

import (
	"encoding/json"
	"fmt"
	"github.com/goiiot/libmqtt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"main/mqttclient"
)

const Port  = "8000"

/**
 * Application Bootstrap
 */
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/devices", GetDevices).Methods("GET")


	client := mqttclient.MqttConnect()



	client.Handle(".*", func(topic string, qos libmqtt.QosLevel, msg []byte) {
		log.Printf("[%v] message: %v", topic, string(msg))
	})

	log.Println(fmt.Sprintf("Bound to port %s", Port))
	log.Fatal(http.ListenAndServe( fmt.Sprintf(":%s", Port), router))


}



func GetDevices(response http.ResponseWriter, r *http.Request) {

	fmt.Println("getting devices")

	var devices []Device

	devices = append(devices, Device{ID:("one"), Name:("devie one") })


	json.NewEncoder(response).Encode(devices)

}


type Device struct {
	ID   string   `json:"id,omitempty"`
	Name string   `json:"firstname,omitempty"`
}



