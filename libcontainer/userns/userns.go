package userns

var (
	// RunningInUserNS detects whether we are currently running in a user namespace.
	// Originally copied from github.com/lxc/lxd/shared/util.go
	RunningInUserNS = runningInUserNS

	UIDMapInUserNS = uidMapInUserNS

	// GetParentNSeuid returns the euid within the parent user namespace
	GetParentNSeuid = getParentNSeuid
)
