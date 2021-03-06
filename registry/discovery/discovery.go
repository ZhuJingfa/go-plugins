// Package discovery is an interface for scalable service discovery.
package discovery

import (
	"micro/go-micro/cmd"
	"micro/go-micro/registry"
	"micro/go-os/discovery"
)

func init() {
	cmd.DefaultRegistries["os"] = NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return discovery.NewRegistry(opts...)
}
