package rbd

import (
	gofigCore "github.com/akutz/gofig"
	gofig "github.com/akutz/gofig/types"
)

const (
	// Name is the name of the storage driver
	Name = "rbd"

	// ConfigDefaultPool is the config key for default pool
	ConfigDefaultPool = Name + ".defaultPool"

	// ConfigTestModule is the config key for testing kernel module presence
	ConfigTestModule = Name + ".testModule"

	// ConfigCephArgs is the config key for the CEPH_ARGS env var
	ConfigCephArgs = Name + ".cephArgs"
)

func init() {
	registerConfig()
}

func registerConfig() {
	r := gofigCore.NewRegistration("RBD")
	r.Key(gofig.String, "", "rbd", "", ConfigDefaultPool)
	r.Key(gofig.Bool, "", true, "", ConfigTestModule)
	r.Key(gofig.String, "", "", "", ConfigCephArgs)
	gofigCore.Register(r)
}
