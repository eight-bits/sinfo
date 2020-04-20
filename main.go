package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

var key string

// Function get Memory
func getInfoMemory() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Error syscall %s", err)
		os.Exit(1)
	}
	syscall.Sysinfo(&info)
	ramTotal := (info.Totalram / 1024) / 1024
	ramFree := (info.Freeram / 1024) / 1024
	ramUsed := ramTotal - ramFree
	ramBuffer := (info.Bufferram / 1024) / 1024
	fmt.Printf("[  RAM   ] - total %d M, used/cache %d M, free %d M, buffer %d M\n", ramTotal, ramUsed, ramFree, ramBuffer)
	swapTotal := (info.Totalswap / 1024) / 1024
	swapFree := (info.Freeswap / 1024) / 1024
	swapUsed := swapTotal - swapFree
	fmt.Printf("[  SWAP  ] - total %d M, used %d M, free %d M\n", swapTotal, swapUsed, swapFree)
}

// Function get local time
func getInfoSystemTime() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Error syscall %s", err)
		os.Exit(1)
	}
	t := time.Now().Format("15:04:05")
	fmt.Printf("[  TIME  ] - %s\n", t)
}

// Function get uptime
func getInfoUptime() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Error syscall %s", err)
		os.Exit(1)
	}
	info.Uptime /= 60
	min := info.Uptime % 60
	info.Uptime /= 60
	hours := info.Uptime % 24
	days := info.Uptime / 24
	fmt.Printf("[ UPTIME ] - %d days, %d hours, %d min\n", days, hours, min)
}

// Function get loading
func getInfoSystemLoad() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Error syscall %s", err)
		os.Exit(1)
	}
	fmt.Printf("[  LOAD  ] - 1m %.02f, 5m %.02f, 15m %.02f\n", float64(info.Loads[0])/65536, float64(info.Loads[1])/65536, float64(info.Loads[2])/65536)
}

// Function get information All
func getInfoAll() {
	// get time
	getInfoSystemTime()
	// get uptime
	getInfoUptime()
	// get memory
	getInfoMemory()
	// get load
	getInfoSystemLoad()
}

func init() {
	flag.StringVar(&key, "s", "a", "Key [a - all, t - time, u - uptime, m - memory, l - loading]")
}

// Function main, point enter
func main() {
	flag.Parse()
	if key == "a" {
		getInfoAll()
		os.Exit(0)
	} else if key == "t" {
		getInfoSystemTime()
		os.Exit(0)
	} else if key == "u" {
		getInfoUptime()
		os.Exit(0)
	} else if key == "m" {
		getInfoMemory()
		os.Exit(0)
	} else if key == "l" {
		getInfoSystemLoad()
		os.Exit(1)
	} else {
		log.Fatalln("No undeclaredet parametr")
		os.Exit(1)
	}
}
