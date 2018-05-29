package sitemap

import (
	"os"
)

func checkSitemap(loc string) bool {
	return true
}

func CreateDirs(dirs []string) (res map[string]bool) {
	res = make(map[string]bool, len(dirs))
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err == nil {
			res[dir] = true
		}
	}
	return
}
