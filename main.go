package main

import (
	"fmt"
	"github.com/pbnjay/memory"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	limit := getLength()
	safety := multiplyToBytes(2, "GB")
	freeMem := memory.FreeMemory()

	if freeMem < uint64(limit+safety) {
		fmt.Println("Not enough memory. Starting in stream writing mode. This may take a long time...")
		streamWrite(limit)
	} else {
		blockWrite(limit)
	}
}

func blockWrite(limit int64) {
	buff := make([]byte, limit)
	rand.Read(buff)
	err := ioutil.WriteFile("paperweight", buff, 777)
	if err != nil {
		exitErr(err)
	}
}

func streamWrite(limit int64) {
	const chunkSize = 16

	remaining := limit
	buff := make([]byte, chunkSize)

	file, err := os.Create("paperweight")
	if err != nil {
		exitErr(err)
	}

	for {
		if remaining <= 0 {
			break
		}

		currentChunkSize := min(remaining, int64(len(buff)))
		if currentChunkSize != int64(len(buff)) {
			buff = make([]byte, currentChunkSize)
		}

		rand.Read(buff)
		_, err := file.Write(buff)
		if err != nil {
			exitErr(err)
		}

		remaining -= currentChunkSize
	}
}

func getLength() int64 {
	arg := os.Args[1]
	if arg == "" {
		exitLength()
	}

	var unitPart string
	var sizePart string
	if isByteEnding(arg) {
		unitPart = arg[len(arg)-1:]
		sizePart = arg[:len(arg)-1]
	} else {
		sizePart = arg[:len(arg)-2]
		unitPart = arg[len(arg)-2:]
	}

	length, err := strconv.ParseInt(sizePart, 10, 64)
	if err != nil {
		exitLength()
	}

	if length < 0 {
		exitMsg("input must be positive")
	}

	length = multiplyToBytes(length, unitPart)

	return length
}

func multiplyToBytes(length int64, unit string) int64 {
	const (
		Multiplier = 1024
		Byte       = 1
		Kilobyte   = Byte * Multiplier
		Megabyte   = Kilobyte * Multiplier
		Gigabyte   = Megabyte * Multiplier
	)

	switch unit {
	default:
		fallthrough
	case "B":
		return length * Byte
	case "KB":
		return length * Kilobyte
	case "MB":
		return length * Megabyte
	case "GB":
		return length * Gigabyte
	}
}

func isByteEnding(input string) bool {
	return strings.HasSuffix(input, "B") && !strings.HasSuffix(input, "KB") && !strings.HasSuffix(input, "MB") && !strings.HasSuffix(input, "GB")
}

func exitLength() {
	exitMsg("argument must be a string representing a size (e.g. \"2GB\")")
}

func exitErr(err error) {
	exitMsg(err.Error())
}

func exitMsg(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func min(x int64, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
