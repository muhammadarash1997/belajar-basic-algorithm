package main

import (
	"time"
	"fmt"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Seconds Nanos
func main() {
	// Auto
	fmt.Println(timestamppb.Now())

	// Semi
	n := time.Now()
	fmt.Printf("seconds:%v nanos:%v\n", n.Unix(), n.Nanosecond())

	// Manual
	t, _ := time.Parse(time.RFC3339, "2022-09-08T06:10:44.762192792Z")
	fmt.Printf("seconds:%v nanos:%v\n", t.Unix(), t.Nanosecond())

	// Unix to RFC3339
	fmt.Println(time.Unix(t.Unix(), int64(t.Nanosecond())).UTC().Format(time.RFC3339))
}