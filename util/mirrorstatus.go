package util

import "fmt"

type MirrorStatusRow struct {
	topic             string
	action            action
	configChange      string
	oldPartitionCount int32
	newPartitionCount int32
	status            status
	reason            string
}

func MirrorStatus(topic, config string, oldPartitionCount, newPartitionCount int32, isCreate, isDryRun bool, err error) MirrorStatusRow {
	var action action
	if isCreate {
		action = create
	} else {
		action = update
	}
	var status status
	var reason string
	if isDryRun {
		status = dryRun
	} else if err == nil {
		status = success
	} else {
		status = failure
		reason = err.Error()
	}

	return MirrorStatusRow{topic: topic, action: action, configChange: config, oldPartitionCount: oldPartitionCount, newPartitionCount: newPartitionCount, status: status, reason: reason}
}

func (m MirrorStatusRow) FieldValues() []string {
	return []string{m.topic, m.action.String(), m.configChange, fmt.Sprint(m.oldPartitionCount), fmt.Sprint(m.newPartitionCount), m.status.String(), m.reason}
}

func (m MirrorStatusRow) Headers() []string {
	return []string{"Topic", "Action", "Configs", "OldPartitionCount", "NewPartitionCount", "Status", "Reason"}
}

type action int

const (
	create action = iota
	update
)

func (s action) String() string {
	return [...]string{"Create", "Update"}[s]
}

type status int

const (
	dryRun status = iota
	success
	failure
)

func (s status) String() string {
	return [...]string{"DryRun", "Success", "Failure"}[s]
}
