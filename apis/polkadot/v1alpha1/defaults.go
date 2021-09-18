package v1alpha1

const (
	// DefaultSyncMode is the default blockchain sync mode
	DefaultSyncMode = FullSynchronization
)

// Resources
const (
	// DefaultNodeCPURequest is the cpu requested by polkadot node
	DefaultNodeCPURequest = "4"
	// DefaultNodeCPULimit is the cpu limit for polkadot node
	DefaultNodeCPULimit = "8"

	// DefaultNodeMemoryRequest is the memory requested by polkadot node
	DefaultNodeMemoryRequest = "4Gi"
	// DefaultNodeMemoryLimit is the memory limit for polkadot node
	DefaultNodeMemoryLimit = "8Gi"

	// DefaultNodeStorageRequest is the Storage requested by polkadot node
	DefaultNodeStorageRequest = "80Gi"
)
