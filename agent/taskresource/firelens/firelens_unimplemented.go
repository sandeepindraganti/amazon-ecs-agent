// +build !linux

// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package firelens

import (
	"errors"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/api/task/status"
	"github.com/aws/amazon-ecs-agent/agent/taskresource"
	resourcestatus "github.com/aws/amazon-ecs-agent/agent/taskresource/status"
)

const (
	// ResourceName is the name of the firelens resource.
	ResourceName = "firelens"
	// FirelensConfigTypeFluentd is the type of a fluentd firelens container.
	FirelensConfigTypeFluentd = "fluentd"
	// FirelensConfigTypeFluentbit is the type of a fluentbit firelens container.
	FirelensConfigTypeFluentbit = "fluentbit"
)

// FirelensResource represents the firelens resource.
type FirelensResource struct{}

// NewFirelensResource returns a new FirelensResource.
func NewFirelensResource(cluster, taskARN, taskDefinition, ec2InstanceID, dataDir, firelensConfigType string,
	ecsMetadataEnabled bool, containerToLogOptions map[string]map[string]string) *FirelensResource {
	return &FirelensResource{}
}

// SetDesiredStatus safely sets the desired status of the resource.
func (firelens *FirelensResource) SetDesiredStatus(status resourcestatus.ResourceStatus) {}

// GetDesiredStatus safely returns the desired status of the resource.
func (firelens *FirelensResource) GetDesiredStatus() resourcestatus.ResourceStatus {
	return resourcestatus.ResourceStatusNone
}

// GetName safely returns the name of the resource.
func (firelens *FirelensResource) GetName() string {
	return "undefined"
}

// DesiredTerminal returns true if the resource's desired status is REMOVED.
func (firelens *FirelensResource) DesiredTerminal() bool {
	return false
}

// KnownCreated returns true if the resource's known status is CREATED.
func (firelens *FirelensResource) KnownCreated() bool {
	return false
}

// TerminalStatus returns the last transition state of firelens resource.
func (firelens *FirelensResource) TerminalStatus() resourcestatus.ResourceStatus {
	return resourcestatus.ResourceStatusNone
}

// NextKnownState returns the state that the resource should progress to based
// on its `KnownState`.
func (firelens *FirelensResource) NextKnownState() resourcestatus.ResourceStatus {
	return resourcestatus.ResourceStatusNone
}

// ApplyTransition calls the function required to move to the specified status.
func (firelens *FirelensResource) ApplyTransition(nextState resourcestatus.ResourceStatus) error {
	return errors.New("not implemented")
}

// SteadyState returns the transition state of the resource defined as "ready".
func (firelens *FirelensResource) SteadyState() resourcestatus.ResourceStatus {
	return resourcestatus.ResourceStatusNone
}

// SetKnownStatus safely sets the currently known status of the resource.
func (firelens *FirelensResource) SetKnownStatus(status resourcestatus.ResourceStatus) {}

// SetAppliedStatus sets the applied status of resource and returns whether
// the resource is already in a transition.
func (firelens *FirelensResource) SetAppliedStatus(status resourcestatus.ResourceStatus) bool {
	return false
}

// GetKnownStatus safely returns the currently known status of the resource.
func (firelens *FirelensResource) GetKnownStatus() resourcestatus.ResourceStatus {
	return resourcestatus.ResourceStatusNone
}

// SetCreatedAt sets the timestamp for resource's creation time.
func (firelens *FirelensResource) SetCreatedAt(createdAt time.Time) {}

// GetCreatedAt sets the timestamp for resource's creation time.
func (firelens *FirelensResource) GetCreatedAt() time.Time {
	return time.Time{}
}

// Create creates the firelens resource.
func (firelens *FirelensResource) Create() error {
	return errors.New("not implemented")
}

// Cleanup cleans up the firelens resource.
func (firelens *FirelensResource) Cleanup() error {
	return errors.New("not implemented")
}

// StatusString returns the string of the firelens resource status.
func (firelens *FirelensResource) StatusString(status resourcestatus.ResourceStatus) string {
	return "undefined"
}

// GetTerminalReason returns the terminal reason for the resource.
func (firelens *FirelensResource) GetTerminalReason() string {
	return "undefined"
}

// MarshalJSON marshals FirelensResource object.
func (firelens *FirelensResource) MarshalJSON() ([]byte, error) {
	return nil, errors.New("not implemented")
}

// UnmarshalJSON unmarshals FirelensResource object.
func (firelens *FirelensResource) UnmarshalJSON(b []byte) error {
	return errors.New("not implemented")
}

// Initialize fills in the resource fields.
func (firelens *FirelensResource) Initialize(resourceFields *taskresource.ResourceFields,
	taskKnownStatus status.TaskStatus,
	taskDesiredStatus status.TaskStatus) {
}
