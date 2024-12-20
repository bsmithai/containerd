package libcontainer

import (
	"github.com/cedana/runc/libcontainer/cgroups"
	"github.com/cedana/runc/libcontainer/intelrdt"
	"github.com/cedana/runc/types"
)

type Stats struct {
	Interfaces    []*types.NetworkInterface
	CgroupStats   *cgroups.Stats
	IntelRdtStats *intelrdt.Stats
}
