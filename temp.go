package main

import (
    "fmt"
	"github.com/yryz/ds18b20"
	
)

func main() {
    sensors, err := ds18b20.Sensors()
    if err != nil {
        panic(err)
    }

    for _,sensor:= range sensors {
        t, err := ds18b20.Temperature(sensor)
        if err == nil {
			fmt.Printf("Actual temp is: %.2f Celsius \n", t)
		fmt.Printf("%.2f \n",t)
		
    }
}}
