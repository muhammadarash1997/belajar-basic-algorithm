package main

import (
	"time"
	"fmt"
	// timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Seconds Nanos
func main() {
	// // Auto
	// fmt.Println(timestamppb.Now())

	// // Semi
	// n := time.Now()
	// fmt.Printf("seconds:%v nanos:%v\n", n.Unix(), n.Nanosecond())

	// // Manual
	// t, _ := time.Parse(time.RFC3339, "2022-09-16T00:32:38.762192792Z")
	// fmt.Printf("seconds:%v nanos:%v\n", t.Unix(), t.Nanosecond())

	// Unix to RFC3339
	// fmt.Println(time.Unix(t.Unix(), int64(t.Nanosecond())).UTC().Format(time.RFC3339))

	// Unix to RFC3339
	// fmt.Println(time.Unix(1663288498, 0).UTC().Format(time.RFC3339))

	// Example Get ETA in Time
	// tm, _ := time.Parse(time.RFC3339, "2022-09-15T03:49:44.762192792Z")
	// fmt.Printf("ETA %v", (tm.Unix()-time.Now().UTC().Unix())/60)

	// Example Get ETA in Time
	now, _ := time.Parse(time.RFC3339, "2022-09-20T04:53:21.0Z")
	fmt.Printf("ETA %v", (1663649762-now.Unix())/60)
}