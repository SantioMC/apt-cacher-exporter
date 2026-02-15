package aptcache

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getHits(data []LineData) int {
	served := count(data, func(line LineData) bool {
		return line.Kind == "O"
	})

	fetched := count(data, func(line LineData) bool {
		return line.Kind == "I"
	})

	return served - fetched
}

func getMisses(data []LineData) int {
	return count(data, func(line LineData) bool {
		return line.Kind == "I"
	})
}

func getMissBandwidth(data []LineData) int {
	return reduce(0, data, func(line LineData) int {
		if line.Kind == "I" {
			return line.Size
		}

		return 0
	})
}

func getCacheBandwidth(data []LineData) int {
	downloadBandwidth := reduce(0, data, func(line LineData) int {
		if line.Kind == "I" {
			return line.Size
		}

		return 0
	})

	serveBandwidth := reduce(0, data, func(line LineData) int {
		if line.Kind == "O" {
			return line.Size
		}

		return 0
	})

	maxData := max(downloadBandwidth, serveBandwidth)
	return maxData - downloadBandwidth
}

func GetData() Export {
	lines := []LineData{}

	files, err := filepath.Glob("/var/log/apt-cacher-ng/apt-cacher*.log")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		reader, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		content, err := io.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}

		if len(content) == 0 {
			continue
		}

		for _, line := range strings.Split(string(content), "\n") {
			if strings.TrimSpace(line) == "" {
				continue
			}

			data, err := parseLine(line)
			if err != nil {
				log.Println(err)
				continue
			}

			lines = append(lines, data)
		}
	}

	return Export{
		Hits:   getHits(lines),
		Misses: getMisses(lines),

		CacheBandwidth: getCacheBandwidth(lines),
		MissBandwidth:  getMissBandwidth(lines),

		Total: count(lines, func(line LineData) bool {
			return line.Kind != "E"
		}),
	}
}
