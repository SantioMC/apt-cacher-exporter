# apt-cacher-exporter

A pretty simple prometheus exporter for apt-cacher-ng, providing basic metrics for the cache and bandwidth usage.

## Usage

```bash
docker run -d -v /var/log/apt-cacher-ng:/var/log/apt-cacher-ng ghcr.io/santiomc/apt-cacher-exporter:latest
```

It is important to mount the log directory to the container, as it will read the logs from there. The exporter will
read all files in the directory that match the pattern `apt-cacher*.log`. (Same behaviour as what apt-cacher-ng does
for its own statistics.) The directory to be mounted in the container is `/var/log/apt-cacher-ng` by default. This is
also where the exporter will write its logs to by default.

## Metrics

| Metric                             | Description                                                          |
| ---------------------------------- | -------------------------------------------------------------------- |
| apt_cacher_requests_hits_total     | The total number of successful cache hits.                           |
| apt_cacher_requests_misses_total   | The total number of misses in the cache.                             |
| apt_cacher_requests_hit_bandwidth  | The total amount of network served from the cache. (Bandwidth saved) |
| apt_cacher_requests_miss_bandwidth | The total amount of network downloaded from upstream.                |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
