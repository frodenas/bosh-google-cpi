package instance

import (
	"google.golang.org/api/compute/v1"
)

type Service interface {
	AddAccessConfig(id string, zone string, networkInterface string, accessConfig *compute.AccessConfig) error
	AddNetworkConfiguration(id string, networks Networks) error
	AttachDisk(id string, diskLink string) (string, string, error)
	AttachedDisks(id string) (AttachedDisks, error)
	CleanUp(id string)
	Create(vmConfig *VMConfig, networks Networks, registryEndpoint string) (string, error)
	Delete(id string) error
	DeleteAccessConfig(id string, zone string, networkInterface string, accessConfig string) error
	DeleteNetworkConfiguration(id string) error
	DetachDisk(id string, diskID string) error
	Find(id string, zone string) (*compute.Instance, bool, error)
	Reboot(id string) error
	SetMetadata(id string, vmMetadata Metadata) error
	SetTags(id string, zone string, instanceTags *compute.Tags) error
	UpdateNetworkConfiguration(id string, networks Networks) error
}

type AttachedDisks []string

type Metadata map[string]interface{}

type VMConfig struct {
	// Name is the optional name that will be assigned to the GCE VM on creation.
	// Names must be unique within a GCP project.
	// The nil value will means a UUID will be used for the VM name.
	Name              string
	Zone              string
	Stemcell          string
	Metadata          Metadata
	MachineType       string
	RootDiskSizeGb    int
	RootDiskType      string
	AutomaticRestart  bool
	OnHostMaintenance string
	Preemptible       bool
	ServiceScopes     ServiceScopes
}

type ServiceScopes []string
