package ecs

import (
	"sep-golang/sep"
)

type Id uint32

type Entity struct {
	id     Id
	tag    string
	active bool
	dcap   *sep.DeviceCapability
}
