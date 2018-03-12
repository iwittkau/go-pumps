package main

import (
	"fmt"

	"math/rand"
	"time"

	pumps "github.com/iwittkau/go-pumps"
)

func main() {

	stop := make(chan struct{}, 1)
	data := make(chan interface{}, 100)

	pp := pumps.NewPump(&stop, &data, PumpHandler)

	ff, err := pp.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println("Pump stopped =", pp.Stopped())

	ff.Input(rand.Float64())
	ff.Input(rand.Float64())
	ff.Input(rand.Float64())
	ff.Input(rand.Float64())

	time.Sleep(time.Second * 5)
	close(stop)
	fmt.Println("Pump stopped =", pp.Stopped())
	result := ff.Input(rand.Float64())
	fmt.Println("Feed accepted =", result)

}

func PumpHandler(data *interface{}) {
	var value float64
	value = (*data).(float64)
	fmt.Println(value)
	time.Sleep(time.Second)
}
