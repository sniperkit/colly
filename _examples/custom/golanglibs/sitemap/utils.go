package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ensurePathExists(path string) error {
	i := strings.LastIndexByte(path, '/')
	if i >= 0 {
		dir := path[:i]
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isRemoteURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func resolvePath(basePath string, relativePath string) string {
	if isRemoteURL(basePath) || isRemoteURL(relativePath) || filepath.IsAbs(relativePath) {
		return relativePath
	}
	return filepath.Join(filepath.Dir(basePath), relativePath)
}

func info(v ...interface{}) {
	if !isVerbose {
		return
	}
	log.Println(v...)
}

func roundU(val float64) int {
	if val > 0 {
		return int(val + 1.0)
	}
	return int(val)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

func parseTimeStamp(utime string) (*time.Time, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return nil, err
	}
	t := time.Unix(i, 0)
	return &t, nil
}

func toBytes(input string) []byte {
	return []byte(input)
}

func mapToString(input map[string]interface{}) string {
	return toString(input)
}

func toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

func convertInterfaceArray(input []map[string]interface{}) []interface{} {
	results := make([]interface{}, len(input))
	for _, result := range input {
		// resultSlice := result.(map[string]interface{})
		// pp.Println("resultSlice=", resultSlice)
		results = append(results, result)
	}
	return results
}

func convertInterface(input map[string]interface{}) []interface{} {
	results := make([]interface{}, len(input))
	for _, result := range input {
		resultSlice := result.(interface{})
		results = append(results, resultSlice)
	}
	return results
}

func shuffle(slc []interface{}) []interface{} {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}

func shuffleInts(slc []int) []int {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}

func shuffleStrings(slc []string) []string {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func updateRequestDelay(beat int, unit string) (delay time.Duration) {
	input := time.Duration(beat)
	switch strings.ToLower(unit) {
	case "microsecond":
		delay = input * time.Microsecond
	case "millisecond":
		delay = input * time.Millisecond
	case "minute":
		delay = input * time.Minute
	case "hour":
		delay = input * time.Hour
	case "second":
		fallthrough
	default:
		delay = input * time.Second
	}
	return
}
