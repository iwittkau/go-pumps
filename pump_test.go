package go_pumps

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPump_Test(t *testing.T) {

	stop := make(chan struct{}, 1)
	data := make(chan interface{}, 100)

	pp := NewPump(&stop, &data, PumpTestHandler)

	ff, err := pp.Run()

	if err != nil {
		t.Error("Pump should be running")
	}

	if pp.Stopped() {
		t.Error("Pump should be running")
	}

	for i := 0; i < 10; i++ {
		ff.Input(rand.Float64())
		time.Sleep(time.Millisecond * 100)
	}

	close(stop)

	if !pp.Stopped() {
		t.Error("Pump should not be running")
	}

}

func TestPump_TestFail(t *testing.T) {

	stop := make(chan struct{}, 1)
	data := make(chan interface{}, 1)

	pp := NewPump(&stop, &data, PumpTestHandler)

	ff1, err1 := pp.Run()

	if err1 != nil {
		t.Error("Should not be an error:", err1)
	}

	if !ff1.Input(rand.Float64()) {
		t.Error("Feed should not fail")
	}

	go func() {
		for {
			if !ff1.Input(rand.Float64()) {
				return
			}
		}
	}()

	go func() {
		for {
			if !ff1.Input(rand.Float64()) {
				return
			}
		}
	}()

	go func() {
		for {
			if pp.Stopped() {
				return
			}
		}
	}()

	time.Sleep(time.Second)

	close(stop)

	if ff1.Input(rand.Float64()) {
		t.Error("Feed should fail")
	}

	ff2, err2 := pp.Run()

	if err2 == nil {
		t.Error("Pump should not be running")
	} else {
		t.Log(err2)
	}

	if ff2 != nil {
		t.Error("Feed should be nil")
	}

}

func PumpTestHandler(data *interface{}) {
	fmt.Println(*data)
}
