// Code generated by "stringer -type=NmCheckpointCreateFlags"; DO NOT EDIT.

package gonetworkmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NmCheckpointCreateFlagsNone-0]
	_ = x[NmCheckpointCreateFlagsDestroyAll-1]
	_ = x[NmCheckpointCreateFlagsDeleteNewConnections-2]
	_ = x[NmCheckpointCreateFlagsDisconnectNewDevices-4]
	_ = x[NmCheckpointCreateFlagsAllowOverlapping-8]
}

const (
	_NmCheckpointCreateFlags_name_0 = "NmCheckpointCreateFlagsNoneNmCheckpointCreateFlagsDestroyAllNmCheckpointCreateFlagsDeleteNewConnections"
	_NmCheckpointCreateFlags_name_1 = "NmCheckpointCreateFlagsDisconnectNewDevices"
	_NmCheckpointCreateFlags_name_2 = "NmCheckpointCreateFlagsAllowOverlapping"
)

var (
	_NmCheckpointCreateFlags_index_0 = [...]uint8{0, 27, 60, 103}
)

func (i NmCheckpointCreateFlags) String() string {
	switch {
	case i <= 2:
		return _NmCheckpointCreateFlags_name_0[_NmCheckpointCreateFlags_index_0[i]:_NmCheckpointCreateFlags_index_0[i+1]]
	case i == 4:
		return _NmCheckpointCreateFlags_name_1
	case i == 8:
		return _NmCheckpointCreateFlags_name_2
	default:
		return "NmCheckpointCreateFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
