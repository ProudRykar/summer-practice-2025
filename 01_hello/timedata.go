package main

import (
        "fmt"
        "time"
)

func main() {
        name := "Nikita Zhukov"
        date := time.Now()

        msg := fmt.Sprintf(
                "Hello, my name is %s. Current date is %s",
                name,
                date.Format("January 2, 2006"),
        )

        fmt.Println(msg)
}
