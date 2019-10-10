package main

import (
	"regexp"

	"github.com/hashicorp/packer/packer/plugin"
)

var (
	VirtualSystemTypeRegex = regexp.MustCompile("<vssd:VirtualSystemType>(.*?)</vssd:VirtualSystemType>")
	AddressRegex           = regexp.MustCompile("<rasd:Address>(.*?)</rasd:Address>")
	CaptionRegex           = regexp.MustCompile("<rasd:Caption>.*?</rasd:Caption>")
	DescriptionRegex       = regexp.MustCompile("<rasd:Description>.*?</rasd:Description>")
	ElementNameRegex       = regexp.MustCompile("<rasd:ElementName>.*?</rasd:ElementName>")
	ResourceSubTypeRegex   = regexp.MustCompile("<rasd:ResourceSubType>.*?</rasd:ResourceSubType>")
	ResourceTypeRegex      = regexp.MustCompile("<rasd:ResourceType>.*?</rasd:ResourceType>")
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterPostProcessor(&PostProcessor{})
	server.Serve()
}
