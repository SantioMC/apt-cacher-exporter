package aptcache

type Export struct {
	Hits   int
	Misses int

	CacheBandwidth int
	MissBandwidth  int

	Total int
}

// "I" indicates a download from upstream, and "O" indicates us serving the request
type LineData struct {
	Timestamp int
	Kind      string
	Size      int
	Ip        string
	File      string
}
