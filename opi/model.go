package opi

import (
	"fmt"

	"code.cloudfoundry.org/eirini/util"
)

const (
	RunningState = "RUNNING"
	PendingState = "CLAIMED"
	CrashedState = "CRASHED"
	UnknownState = "UNKNOWN"
)

type LRPIdentifier struct {
	GUID, Version string
	Hasher        util.Hasher
}

func (i *LRPIdentifier) Name() string {
	hashed, _ := i.Hasher.Hash(fmt.Sprintf("%s-%s", i.GUID, i.Version))
	//TODO: figure out what to do with this error
	return hashed
}

func (i *LRPIdentifier) NotSureWhatToCallThis() string {
	return fmt.Sprintf("%s-%s", i.GUID, i.Version)
}

// An LRP, or long-running-process, is a stateless process
// where the scheduler should attempt to keep N copies running,
// killing and recreating as needed to maintain that guarantee
type LRP struct {
	LRPIdentifier
	Name             string
	Image            string
	Command          []string
	Env              map[string]string
	Health           Healtcheck
	Ports            []int32
	TargetInstances  int
	RunningInstances int
	Metadata         map[string]string
}

type Instance struct {
	Index int
	Since int64
	State string
}

type Healtcheck struct {
	Type      string
	Port      int32
	Endpoint  string
	TimeoutMs uint
}

// A Task is a one-off process that is run exactly once and returns a
// result
type Task struct {
	Image   string
	Command []string
	Env     map[string]string
}

//go:generate counterfeiter . Desirer
type Desirer interface {
	Desire(lrp *LRP) error
	List() ([]*LRP, error)
	Get(identifier LRPIdentifier) (*LRP, error)
	GetInstances(identifier LRPIdentifier) ([]*Instance, error)
	Update(lrp *LRP) error
	Stop(identifier LRPIdentifier) error
}

//go:generate counterfeiter . TaskDesirer
type TaskDesirer interface {
	Desire(task *Task) error
	DesireStaging(task *Task) error
	Delete(name string) error
}
