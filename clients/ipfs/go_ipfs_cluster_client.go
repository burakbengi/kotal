package ipfs

import (
	"os"
	"strings"

	ipfsv1alpha1 "github.com/kotalco/kotal/apis/ipfs/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GoIPFSClusterClient is ipfs cluster service client
// https://github.com/ipfs/ipfs-cluster
type GoIPFSClusterClient struct {
	peer *ipfsv1alpha1.ClusterPeer
}

const (
	// EnvGoIPFSClusterImage is the environment variable used for go ipfs cluster client image
	EnvGoIPFSClusterImage = "GO_IPFS_CLUSTER_IMAGE"
	// DefaultGoIPFSClusterImage is the default go ipfs cluster client image
	DefaultGoIPFSClusterImage = "kotalco/ipfs-cluster:v0.14.2"
	//  GoIPFSClusterHomeDir is go ipfs cluster image home dir
	GoIPFSClusterHomeDir = "/home/ipfs-cluster"
)

// Image returns go ipfs cluster image
func (c *GoIPFSClusterClient) Image() string {
	if os.Getenv(EnvGoIPFSClusterImage) == "" {
		return DefaultGoIPFSClusterImage
	}
	return os.Getenv(EnvGoIPFSClusterImage)
}

// Command returns go ipfs cluster entrypoint
func (c *GoIPFSClusterClient) Command() []string {
	return []string{"ipfs-cluster-service"}
}

// Command returns environment variables for the client
func (c *GoIPFSClusterClient) Env() []corev1.EnvVar {
	return nil
}

// Arg returns go ipfs cluster arguments
func (c *GoIPFSClusterClient) Args() (args []string) {
	args = append(args, GoIPFSClusterDaemonArg)

	if len(c.peer.Spec.BootstrapPeers) != 0 {
		args = append(args, GoIPFSClusterBootstrapArg, strings.Join(c.peer.Spec.BootstrapPeers, ","))
	}

	return
}

// HomeDir returns go ipfs cluster image home directory
func (c *GoIPFSClusterClient) HomeDir() string {
	return GoIPFSClusterHomeDir
}
