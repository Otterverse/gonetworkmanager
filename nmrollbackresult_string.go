// Code generated by "stringer -type=NmRollbackResult"; DO NOT EDIT.

package gonetworkmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NmRollbackResultOk-0]
	_ = x[NmRollbackResultErrNoDevice-1]
	_ = x[NmRollbackResultErrDeviceUnmanaged-2]
	_ = x[NmRollbackResultErrFailed-3]
}

const _NmRollbackResult_name = "NmRollbackResultOkNmRollbackResultErrNoDeviceNmRollbackResultErrDeviceUnmanagedNmRollbackResultErrFailed"

var _NmRollbackResult_index = [...]uint8{0, 18, 45, 79, 104}

func (i NmRollbackResult) String() string {
	if i >= NmRollbackResult(len(_NmRollbackResult_index)-1) {
		return "NmRollbackResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NmRollbackResult_name[_NmRollbackResult_index[i]:_NmRollbackResult_index[i+1]]
}
