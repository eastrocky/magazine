// +build magazine

package main

import (
	"testing"

	"github.com/eastrocky/magazine"
)

// go test -tags magazine ./...
func TestEjectConfig(t *testing.T) {
	magazine.Eject("config.yml", Config{
		AWS: AWS{
			Region: "us-west-2", // default region
		},
	})
}
