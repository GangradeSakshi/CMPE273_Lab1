package main

import (
  "fmt"
  "time"
)

func sleep(){
    <-time.After(5 * time.Second)
    fmt.Println("Lab Assignment")
}

func main() {
  fmt.Println("Sleep started")
  sleep()
  fmt.Println("Sleep finished")
}
