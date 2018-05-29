package convert

type stats struct {
	extensions map[string]int
	available  int
	processed  int
	skipped    int
	errors     int
	size       int
	count      int
}
