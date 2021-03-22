package system

import "github.com/opencontainers/runc/libcontainer/userns"

// RunningInUserNS detects whether we are currently running in a user namespace.
// Originally copied from github.com/lxc/lxd/shared/util.go
// Deprecated: use github.com/opencontainers/runc/libcontainer/userns.RunningInUserNS
var RunningInUserNS = userns.RunningInUserNS
