package system

import "github.com/opencontainers/runc/libcontainer/userns"

var (
	// RunningInUserNS detects whether we are currently running in a user namespace.
	// Originally copied from github.com/lxc/lxd/shared/util.go
	// Deprecated: use github.com/opencontainers/runc/libcontainer/userns.RunningInUserNS
	RunningInUserNS = userns.RunningInUserNS

	// Deprecated: use github.com/opencontainers/runc/libcontainer/userns.UIDMapInUserNS
	UIDMapInUserNS = userns.RunningInUserNS

	// GetParentNSeuid returns the euid within the parent user namespace
	// Deprecated: use github.com/opencontainers/runc/libcontainer/userns.RunningInUserNS
	GetParentNSeuid = userns.RunningInUserNS
)
