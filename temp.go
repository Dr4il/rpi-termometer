package main

import (
        "fmt"
        "github.com/yryz/ds18b20"
        "os"
)

func main() {

        sensors, err := ds18b20.Sensors()
        if err != nil {
                panic(err)
        }
        z, err := os.Create("/var/www/html/index.html")
        if err != nil {
                fmt.Println(err)
                return
}

        for _, sensor := range sensors {
                t, err := ds18b20.Temperature(sensor)
                if err == nil {
                        fmt.Printf("Actual temp is: %.2f Celsius \n", t)
                        fmt.Printf("%.2f \n", t)
                }
        }
        for _,sensor := range sensors {
                 p, err := ds18b20.Temperature(sensor)
                if err == nil {
			x := fmt.Sprintf("Actual temp in room is: %.2f Celsius \n", p)
                        z.WriteString(x)
                        if err  != nil {
                        fmt.Println(err)
                        z.Close()
                        return
}
}
}
}
