package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Vimeo struct {
	Thumbnail_url string
}

func run() error {

	var url = "https://vimeo.com/api/oembed.json?url=https://vimeo.com/"

	if len(os.Args) < 2 {
		return fmt.Errorf("Parameter is invalid")
	}

	url += os.Args[1]

	if len(os.Args) == 3 {
		url += "&width=" + os.Args[2]
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to connect api")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response")
	}

	jsonBytes := ([]byte)(body)
	data := new(Vimeo)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return fmt.Errorf("Json Unmarshal error")
	}

	fmt.Printf("%s", data.Thumbnail_url)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
