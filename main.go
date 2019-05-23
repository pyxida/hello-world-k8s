package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "my_custom_metric",
			Help: "custom metrics for application",
		},
		[]string{"mylabel"})
)

func incHandler(w http.ResponseWriter, r *http.Request) {
	httpReqs.WithLabelValues("Value").Add(1)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	httpReqs.WithLabelValues("Values").Set(0)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from : %s\n", os.Getenv("HOSTNAME"))
	fmt.Fprintf(w, "Value from config : MYVAR => %s\n", os.Getenv("MYVAR"))

	b, _ := ioutil.ReadFile(os.Getenv("MYFILE"))

	fmt.Fprintf(w, "Value from file : \n")
	fmt.Fprintf(w, "%s => \n%s", os.Getenv("MYFILE"), string(b))
}

func main() {
	prometheus.MustRegister(httpReqs)
	http.HandleFunc("/inc", incHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
