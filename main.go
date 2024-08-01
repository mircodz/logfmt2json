package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"

	"github.com/go-logfmt/logfmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		s := scanner.Bytes()
		decoder := logfmt.NewDecoder(bytes.NewReader(s))

		o := make(map[string]string)
		for decoder.ScanRecord() {
			for decoder.ScanKeyval() {
				o[string(decoder.Key())] = string(decoder.Value())
			}
		}

		bs, err := json.Marshal(o)
		if err != nil {
			continue
		}

		if _, err := os.Stdout.Write(bs); err != nil {
			continue
		}

		if _, err := os.Stdout.Write([]byte("\n")); err != nil {
			continue
		}

	}
}
