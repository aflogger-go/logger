// Code generated by "stringer -type=Level"; DO NOT EDIT.

package logger

import "strconv"

const _Level_name = "DebugInfoWarnErrCrit"

var _Level_index = [...]uint8{0, 5, 9, 13, 16, 20}

func (i Level) String() string {
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}