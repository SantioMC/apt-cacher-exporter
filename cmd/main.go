package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/santiomc/apt-cacher-exporter/aptcache"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "apt-cacher-exporter",
	})

	app.Get("/metrics", getMetrics)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 3000")
}

func getMetrics(ctx fiber.Ctx) error {
	data := aptcache.GetData()

	return ctx.SendString(clean(`
	# HELP apt_cacher_requests_hits_total The total number of successful cache hits in the cache.
	# TYPE apt_cacher_requests_hits_total counter
	apt_cacher_requests_hits_total ` + strconv.Itoa(data.Hits) + `

	# HELP apt_cacher_requests_misses_total The total number of misses in the cache.
	# TYPE apt_cacher_requests_misses_total counter
	apt_cacher_requests_misses_total ` + strconv.Itoa(data.Misses) + `

	# HELP apt_cacher_requests_hit_bandwidth The total amount of network served from the cache.
	# TYPE apt_cacher_requests_hit_bandwidth gauge
	apt_cacher_requests_hit_bandwidth ` + strconv.Itoa(data.CacheBandwidth) + `

	# HELP apt_cacher_requests_miss_bandwidth The total amount of network served from upstream.
	# TYPE apt_cacher_requests_miss_bandwidth gauge
	apt_cacher_requests_miss_bandwidth ` + strconv.Itoa(data.MissBandwidth) + `
	`))
}

func clean(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\t", ""))
}
