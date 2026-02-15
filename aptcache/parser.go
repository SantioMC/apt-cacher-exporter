package aptcache

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (LineData, error) {
	data := strings.Split(line, "|")

	if len(data) != 5 {
		return LineData{}, fmt.Errorf("invalid line format: %s", line)
	}

	timestamp, err := strconv.Atoi(data[0])
	if err != nil {
		return LineData{}, err
	}

	size, err := strconv.Atoi(data[2])
	if err != nil {
		return LineData{}, err
	}

	return LineData{
		Timestamp: timestamp,
		Kind:      data[1],
		Size:      size,
		Ip:        data[3],
		File:      data[4],
	}, nil
}
