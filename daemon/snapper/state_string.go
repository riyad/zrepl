// Code generated by "stringer -type=State"; DO NOT EDIT.

package snapper

import "strconv"

const (
	_State_name_0 = "SyncUpSyncUpErrWait"
	_State_name_1 = "Planning"
	_State_name_2 = "Snapshotting"
	_State_name_3 = "Waiting"
	_State_name_4 = "ErrorWait"
	_State_name_5 = "Stopped"
)

var (
	_State_index_0 = [...]uint8{0, 6, 19}
)

func (i State) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _State_name_0[_State_index_0[i]:_State_index_0[i+1]]
	case i == 4:
		return _State_name_1
	case i == 8:
		return _State_name_2
	case i == 16:
		return _State_name_3
	case i == 32:
		return _State_name_4
	case i == 64:
		return _State_name_5
	default:
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
