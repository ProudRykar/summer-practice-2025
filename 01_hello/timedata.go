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
                date.Format("November 11, 2000"),
        )

        fmt.Println(msg)
}
