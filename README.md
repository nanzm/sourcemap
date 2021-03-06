# Source maps consumer for Golang

> fork form https://github.com/go-sourcemap/sourcemap

## Installation

Install:

```shell
go get -u github.com/nanzm/sourcemap
```

## Quickstart

```go
func ExampleParse() {
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
```
