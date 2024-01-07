package ecs

import (
	"github.com/Tylores/sep-models/sep"
)

type Id uint32

type Entity struct {
	id              Id
	tag             string
	active          bool
	DevicCapability *sep.DeviceCapability
}
