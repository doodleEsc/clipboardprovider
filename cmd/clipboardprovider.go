package main

import (
	"clipboardprovider"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	method := strings.ToLower(os.Args[1])
	if method != "copy" && method != "paste" {
		return
	}

	switch method {
	case "copy":
		data, err := clipboardprovider.PipeInput()
		if err != nil {
			return
		}
		ps := &clipboardprovider.ProviderSet{
			&clipboardprovider.PB{},
		}
		ps.Copy(data)
	case "paste":
		ps := &clipboardprovider.ProviderSet{
			&clipboardprovider.PB{},
		}
		ps.Paste()
	}
}
