package main

import (
	"C"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thbkrkr/leds/ws2811"
	"github.com/thbkrkr/qli/client"
)

func main() {
	qli, err := client.NewClientFromEnv()
	handlErr(err, "Fail to create qli client", true)

	err = ws2811.Init(18, 1500, 255)
	handlErr(err, "Fail to init ws2811", true)

	for cmd := range qli.Sub() {
		go apply(cmd)
	}
}

func apply(cmd string) {
	leds := map[string]string{}

	err := json.Unmarshal([]byte(cmd), &leds)
	if err != nil {
		handlErr(err, "Fail to unmarshal json cmd", false)
		return
	}

	for k, rgb := range leds {
		numLed, err := strconv.Atoi(k)
		if err != nil {
			handlErr(err, "Fail to parse led index", false)
			return
		}

		color, err := colorToInt(rgb)
		if err != nil {
			handlErr(err, "Fail to parse color", false)
			return
		}

		ws2811.SetLed(numLed, color)
	}

	err = ws2811.Render()
	handlErr(err, "Fail to init ws2811", false)
}

func colorToInt(rgb string) (uint32, error) {
	rgbArr := strings.Split(rgb, ",")
	red, err := strconv.Atoi(rgbArr[0])
	if err != nil {
		return 0, err
	}
	green, err := strconv.Atoi(rgbArr[1])
	if err != nil {
		return 0, err
	}
	blue, err := strconv.Atoi(rgbArr[2])
	if err != nil {
		return 0, err
	}
	return uint32((red << 16) | (green << 8) | blue), nil
}

func handlErr(err error, context string, exit bool) {
	if err != nil {
		fmt.Printf("%s: %s\n", context, err)
		if exit {
			os.Exit(1)
		}
	}
}
