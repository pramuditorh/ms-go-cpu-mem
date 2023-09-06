package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func simulateCPUUsage(duration time.Duration) {
	if duration == 0 {
		for {
			math.Sqrt(999999999)
		}
	} else {
		endTime := time.Now().Add(duration)
		for time.Now().Before(endTime) {
			math.Sqrt(999999999)
		}
	}
}

func simulateMemoryUsage(duration time.Duration) {
	if duration == 0 {
		var memoryHog []byte
		for {
			memoryHog = append(memoryHog, make([]byte, 1024*1024)...) // Allocate 1 MB of memory
			time.Sleep(time.Millisecond)
		}
	} else {
		endTime := time.Now().Add(duration)
		var memoryHog []byte
		for time.Now().Before(endTime) {
			memoryHog = append(memoryHog, make([]byte, 1024*1024)...) // Allocate 1 MB of memory
			time.Sleep(time.Millisecond)
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello")
	})

	http.HandleFunc("/api/v1/simulate/cpu/", func(w http.ResponseWriter, r *http.Request) {
		durationStr := r.URL.Path[len("/api/v1/simulate/cpu/"):]
		var duration time.Duration
		if durationStr != "" {
			durationInt, err := strconv.Atoi(durationStr)
			if err != nil {
				http.Error(w, "Invalid duration", http.StatusBadRequest)
				return
			}
			duration = time.Duration(durationInt) * time.Second
		}
		go simulateCPUUsage(duration)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "CPU simulation started\n")
	})

	http.HandleFunc("/api/v1/simulate/memory/", func(w http.ResponseWriter, r *http.Request) {
		durationStr := r.URL.Path[len("/api/v1/simulate/memory/"):]
		var duration time.Duration
		if durationStr != "" {
			durationInt, err := strconv.Atoi(durationStr)
			if err != nil {
				http.Error(w, "Invalid duration", http.StatusBadRequest)
				return
			}
			duration = time.Duration(durationInt) * time.Second
		}
		go simulateMemoryUsage(duration)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Memory simulation started\n")
	})

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("Starting API server...")
	http.ListenAndServe(":8080", nil)
}
