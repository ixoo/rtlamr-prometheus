package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var	homeMeteredConsumption = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_metered_consumption",
			Help: "Home consumption measurement (water, gas, electricity).",
		},
		[]string{"id","type"},
		)

func init() {
	prometheus.MustRegister(homeMeteredConsumption)
}

type Message struct {
	Time time.Time `json:"Time"`
	SCM  SCM       `json:"Message"`
}

type SCM struct {
	ID		uint32	`json:"ID"`
	Type		uint8	`json:"Type"`
	Consumption	float64	`json:"Consumption"`
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go func(){
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	stdinBuf := bufio.NewScanner(os.Stdin)

	for stdinBuf.Scan() {
		var msg Message
		err := json.Unmarshal(stdinBuf.Bytes(), &msg)
		if err != nil {
			log.Println(err)
			continue
		}

		homeMeteredConsumption.With(
			prometheus.Labels{
				"id":	fmt.Sprint(msg.SCM.ID),
				"type":	fmt.Sprint(msg.SCM.Type)}).Set(msg.SCM.Consumption)

	}
}
