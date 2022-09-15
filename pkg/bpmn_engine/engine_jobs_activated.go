package bpmn_engine

import (
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine/variable_scope"
	"time"

	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20/process_instance"
)

// ActivatedJob is a struct to provide information for registered task handler
type activatedJob struct {
	processInstanceInfo      *ProcessInstanceInfo
	completeHandler          func()
	failHandler              func(reason string)
	key                      int64
	processInstanceKey       int64
	bpmnProcessId            string
	processDefinitionVersion int32
	processDefinitionKey     int64
	elementId                string
	createdAt                time.Time
	scope                    variable_scope.VarScope
	localScope               variable_scope.VarScope
}

// ActivatedJob represents an abstraction for the activated job
// don't forget to call Fail or Complete when your task worker job is complete or not.
type ActivatedJob interface {
	// GetInstanceKey get instance key from processInfo
	GetInstanceKey() int64
	// GetCreatedAt  get job create time
	GetCreatedAt() time.Time
	// GetState get instance state
	GetState() process_instance.State

	// GetKey the key, a unique identifier for the job
	GetKey() int64

	// GetVariable the varaible from variable scope  include local scope
	GetVariable(key string) interface{}

	// SetVariable set variable into variable scope
	SetVariable(key string, value interface{})

	// SetVariableLocal set variable into local variable scope
	SetVariableLocal(key string, value interface{})

	// GetProcessInstanceKey the job's process instance key
	GetProcessInstanceKey() int64

	// GetBpmnProcessId Retrieve id of the job process definition
	GetBpmnProcessId() string

	// GetProcessDefinitionVersion Retrieve version of the job process definition
	GetProcessDefinitionVersion() int32

	// GetProcessDefinitionKey Retrieve key of the job process definition
	GetProcessDefinitionKey() int64

	// GetElementId Get element id of the job
	GetElementId() string

	// Fail does set the state the worker missed completing the job
	// Fail and Complete mutual exclude each other
	Fail(reason string)

	// Complete does set the state the worker successfully completing the job
	// Fail and Complete mutual exclude each other
	Complete()
}

// GetCreatedAt implements ActivatedJob
func (aj *activatedJob) GetCreatedAt() time.Time {
	return aj.createdAt
}

// GetInstanceKey implements ActivatedJob
func (aj *activatedJob) GetInstanceKey() int64 {
	return aj.processInstanceInfo.GetInstanceKey()
}

// GetProcessInfo implements ActivatedJob
func (aj *activatedJob) GetProcessInfo() *ProcessInfo {
	return aj.processInstanceInfo.GetProcessInfo()
}

// GetState implements ActivatedJob
func (aj *activatedJob) GetState() process_instance.State {
	return aj.processInstanceInfo.GetState()
}

// GetElementId implements ActivatedJob
func (aj *activatedJob) GetElementId() string {
	return aj.elementId
}

// GetKey implements ActivatedJob
func (aj *activatedJob) GetKey() int64 {
	return aj.key
}

// GetBpmnProcessId implements ActivatedJob
func (aj *activatedJob) GetBpmnProcessId() string {
	return aj.bpmnProcessId
}

// GetProcessDefinitionKey implements ActivatedJob
func (aj *activatedJob) GetProcessDefinitionKey() int64 {
	return aj.processDefinitionKey
}

// GetProcessDefinitionVersion implements ActivatedJob
func (aj *activatedJob) GetProcessDefinitionVersion() int32 {
	return aj.processDefinitionVersion
}

// GetProcessInstanceKey implements ActivatedJob
func (aj *activatedJob) GetProcessInstanceKey() int64 {
	return aj.processInstanceKey
}

// GetVariable implements ActivatedJob
func (aj *activatedJob) GetVariable(key string) interface{} {
	if aj.localScope.GetVariable(key) != nil {
		return aj.localScope.GetVariable(key)
	}
	return aj.scope.GetVariable(key)
}

// SetVariable implements ActivatedJob
func (aj *activatedJob) SetVariable(key string, value interface{}) {
	aj.scope.SetVariable(key, value)
}

// SetVariableLocal implements ActivatedJob
func (aj *activatedJob) SetVariableLocal(key string, value interface{}) {
	aj.localScope.SetVariable(key, value)
}

// Fail implements ActivatedJob
func (aj *activatedJob) Fail(reason string) {
	aj.failHandler(reason)
}

// Complete implements ActivatedJob
func (aj *activatedJob) Complete() {
	aj.completeHandler()
	aj.scope.Propagation()
}
