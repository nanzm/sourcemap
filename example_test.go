package sourcemap_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/nanzm/sourcemap"
)

func TestExampleParse(t *testing.T) {
	mapURL := "http://code.jquery.com/jquery-2.0.3.min.map"
	resp, err := http.Get(mapURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	smap, err := sourcemap.Parse(mapURL, b)
	if err != nil {
		panic(err)
	}

	line, column := 5, 6789
	file, fn, line, col, ok := smap.Source(line, column)
	fmt.Println(file, fn, line, col, ok)
	// Output: http://code.jquery.com/jquery-2.0.3.js apply 4360 27 true
}
