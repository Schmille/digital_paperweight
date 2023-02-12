package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	limit := getLength()
	var index int64
	buff := make([]byte, 1)

	file, err := os.Create("paperweight")
	if err != nil {
		exitErr(err)
	}

	for index = 0; index < limit; index++ {
		rand.Read(buff)
		_, err := file.Write(buff)
		if err != nil {
			exitErr(err)
		}
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
