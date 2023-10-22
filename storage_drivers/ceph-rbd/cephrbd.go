package cephrbd

import (
	"context"

	"github.com/RoaringBitmap/roaring"
	tridentconfig "github.com/netapp/trident/config"
	"github.com/netapp/trident/storage"
	sa "github.com/netapp/trident/storage_attribute"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/utils"
)

type CephRBDStorageDriver struct {
	initialized bool
	Config      drivers.CephRBDStorageConfig
}

type Telemetry struct {
	tridentconfig.Telemetry
	Plugin string `json:"plugin"`
}

func (rbd CephRBDStorageDriver) Name() string {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) BackendName() string {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Initialize(_ context.Context, _ tridentconfig.DriverContext, _ string, _ *drivers.CommonStorageDriverConfig, _ map[string]string, _ string) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Initialized() bool {
	panic("not implemented") // TODO: Implement
}

// Terminate tells the driver to clean up, as it won't be called again.
func (rbd CephRBDStorageDriver) Terminate(ctx context.Context, backendUUID string) {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Create(ctx context.Context, volConfig *storage.VolumeConfig, storagePool storage.Pool, volAttributes map[string]sa.Request) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) CreatePrepare(ctx context.Context, volConfig *storage.VolumeConfig) {
	panic("not implemented") // TODO: Implement
}

// CreateFollowup adds necessary information for accessing the volume to VolumeConfig.
func (rbd CephRBDStorageDriver) CreateFollowup(ctx context.Context, volConfig *storage.VolumeConfig) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) CreateClone(ctx context.Context, sourceVolConfig *storage.VolumeConfig, cloneVolConfig *storage.VolumeConfig, storagePool storage.Pool) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Import(ctx context.Context, volConfig *storage.VolumeConfig, originalName string) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Destroy(ctx context.Context, volConfig *storage.VolumeConfig) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Rename(ctx context.Context, name string, newName string) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Resize(ctx context.Context, volConfig *storage.VolumeConfig, sizeBytes uint64) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Get(ctx context.Context, name string) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetInternalVolumeName(ctx context.Context, name string) string {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetStorageBackendSpecs(ctx context.Context, backend storage.Backend) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetStorageBackendPhysicalPoolNames(ctx context.Context) []string {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetProtocol(ctx context.Context) tridentconfig.Protocol {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) Publish(ctx context.Context, volConfig *storage.VolumeConfig, publishInfo *utils.VolumePublishInfo) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) CanSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) (*storage.Snapshot, error) {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetSnapshots(ctx context.Context, volConfig *storage.VolumeConfig) ([]*storage.Snapshot, error) {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) CreateSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) (*storage.Snapshot, error) {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) RestoreSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) DeleteSnapshot(ctx context.Context, snapConfig *storage.SnapshotConfig, volConfig *storage.VolumeConfig) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) StoreConfig(ctx context.Context, b *storage.PersistentStorageBackendConfig) {
	panic("not implemented") // TODO: Implement
}

// GetExternalConfig returns a version of the driver configuration that
// lacks confidential information, such as usernames and passwords.
func (rbd CephRBDStorageDriver) GetExternalConfig(ctx context.Context) interface{} {
	panic("not implemented") // TODO: Implement
}

// GetVolumeExternal accepts the internal name of a volume and returns a VolumeExternal
// object.  This method is only available if using the passthrough store (i.e. Docker).
func (rbd CephRBDStorageDriver) GetVolumeExternal(ctx context.Context, name string) (*storage.VolumeExternal, error) {
	panic("not implemented") // TODO: Implement
}

// GetVolumeExternalWrappers reads all volumes owned by this driver from the storage backend and
// writes them to the supplied channel as VolumeExternalWrapper objects.  This method is only
// available if using the passthrough store (i.e. Docker).
func (rbd CephRBDStorageDriver) GetVolumeExternalWrappers(_ context.Context, _ chan *storage.VolumeExternalWrapper) {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetUpdateType(ctx context.Context, driver storage.Driver) *roaring.Bitmap {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) ReconcileNodeAccess(ctx context.Context, nodes []*utils.Node, backendUUID string, tridentUUID string) error {
	panic("not implemented") // TODO: Implement
}

func (rbd CephRBDStorageDriver) GetCommonConfig(_ context.Context) *drivers.CommonStorageDriverConfig {
	panic("not implemented") // TODO: Implement
}
