// Code generated by "stringer -type=JobStatus"; DO NOT EDIT

package main

import "fmt"

const _JobStatus_name = "InitialisedDetachingSnapshottingCreatingImageDone"

var _JobStatus_index = [...]uint8{0, 11, 20, 32, 45, 49}

func (i JobStatus) String() string {
	if i < 0 || i >= JobStatus(len(_JobStatus_index)-1) {
		return fmt.Sprintf("JobStatus(%d)", i)
	}
	return _JobStatus_name[_JobStatus_index[i]:_JobStatus_index[i+1]]
}