// Code generated by "stringer -type=ConnectionStatus ./v2"; DO NOT EDIT.

package yarpc

import "strconv"

const _ConnectionStatus_name = "UnavailableConnectingAvailable"

var _ConnectionStatus_index = [...]uint8{0, 11, 21, 30}

func (i ConnectionStatus) String() string {
	if i < 0 || i >= ConnectionStatus(len(_ConnectionStatus_index)-1) {
		return "ConnectionStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ConnectionStatus_name[_ConnectionStatus_index[i]:_ConnectionStatus_index[i+1]]
}
