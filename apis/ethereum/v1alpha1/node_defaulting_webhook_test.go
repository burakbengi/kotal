package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Ethereum defaulting", func() {
	It("Should default nodes joining mainnet", func() {
		availabilityConfig := AvailabilityConfig{
			HighlyAvailable: true,
		}

		node1 := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				AvailabilityConfig: availabilityConfig,
				Client:             BesuClient,
				Network:            MainNetwork,
			},
		}

		node2 := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-2",
			},
			Spec: NodeSpec{
				AvailabilityConfig: availabilityConfig,
				Client:             BesuClient,
				Network:            MainNetwork,
				SyncMode:           FullSynchronization,
			},
		}

		node1.Default()
		node2.Default()

		// node1 defaulting
		Expect(node1.Spec.TopologyKey).To(Equal(DefaultTopologyKey))
		Expect(node1.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node1.Spec.SyncMode).To(Equal(DefaultPublicNetworkSyncMode))
		Expect(node1.Spec.Resources.CPU).To(Equal(DefaultPublicNetworkNodeCPURequest))
		Expect(node1.Spec.Resources.CPULimit).To(Equal(DefaultPublicNetworkNodeCPULimit))
		Expect(node1.Spec.Resources.Memory).To(Equal(DefaultPublicNetworkNodeMemoryRequest))
		Expect(node1.Spec.Resources.MemoryLimit).To(Equal(DefaultPublicNetworkNodeMemoryLimit))
		Expect(node1.Spec.Resources.Storage).To(Equal(DefaultMainNetworkFastNodeStorageRequest))
		Expect(node1.Spec.Logging).To(Equal(DefaultLogging))
		// node2 defaulting
		Expect(node2.Spec.TopologyKey).To(Equal(DefaultTopologyKey))
		Expect(node2.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node2.Spec.SyncMode).To(Equal(FullSynchronization))
		Expect(node2.Spec.Resources.CPU).To(Equal(DefaultPublicNetworkNodeCPURequest))
		Expect(node2.Spec.Resources.CPULimit).To(Equal(DefaultPublicNetworkNodeCPULimit))
		Expect(node2.Spec.Resources.Memory).To(Equal(DefaultPublicNetworkNodeMemoryRequest))
		Expect(node2.Spec.Resources.MemoryLimit).To(Equal(DefaultPublicNetworkNodeMemoryLimit))
		Expect(node2.Spec.Resources.Storage).To(Equal(DefaultMainNetworkFullNodeStorageRequest))
		Expect(node2.Spec.Logging).To(Equal(DefaultLogging))

	})

	It("Should default nodes joining rinkeby", func() {
		availabilityConfig := AvailabilityConfig{
			HighlyAvailable: true,
		}

		node := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				AvailabilityConfig: availabilityConfig,
				Client:             BesuClient,
				Network:            RinkebyNetwork,
			},
		}

		node.Default()
		Expect(node.Spec.TopologyKey).To(Equal(DefaultTopologyKey))
		Expect(node.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node.Spec.SyncMode).To(Equal(DefaultPublicNetworkSyncMode))
		Expect(node.Spec.Resources.CPU).To(Equal(DefaultPublicNetworkNodeCPURequest))
		Expect(node.Spec.Resources.CPULimit).To(Equal(DefaultPublicNetworkNodeCPULimit))
		Expect(node.Spec.Resources.Memory).To(Equal(DefaultPublicNetworkNodeMemoryRequest))
		Expect(node.Spec.Resources.MemoryLimit).To(Equal(DefaultPublicNetworkNodeMemoryLimit))
		Expect(node.Spec.Resources.Storage).To(Equal(DefaultTestNetworkStorageRequest))
		Expect(node.Spec.Logging).To(Equal(DefaultLogging))
	})

	It("Should default geth node joining rinkeby", func() {
		availabilityConfig := AvailabilityConfig{
			HighlyAvailable: true,
		}

		node := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				AvailabilityConfig: availabilityConfig,
				Client:             GethClient,
				Network:            RinkebyNetwork,
			},
		}

		node.Default()
		Expect(node.Spec.TopologyKey).To(Equal(DefaultTopologyKey))
		Expect(node.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node.Spec.SyncMode).To(Equal(SnapSynchronization))
		Expect(node.Spec.Resources.CPU).To(Equal(DefaultPublicNetworkNodeCPURequest))
		Expect(node.Spec.Resources.CPULimit).To(Equal(DefaultPublicNetworkNodeCPULimit))
		Expect(node.Spec.Resources.Memory).To(Equal(DefaultPublicNetworkNodeMemoryRequest))
		Expect(node.Spec.Resources.MemoryLimit).To(Equal(DefaultPublicNetworkNodeMemoryLimit))
		Expect(node.Spec.Resources.Storage).To(Equal(DefaultTestNetworkStorageRequest))
		Expect(node.Spec.Logging).To(Equal(DefaultLogging))
	})

	It("Should default nodes joining network pow consensus", func() {
		node := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				Genesis: &Genesis{
					ChainID:   55555,
					NetworkID: 55555,
					Ethash:    &Ethash{},
				},
				Client: BesuClient,
			},
		}

		node.Default()
		var block0 uint = 0
		Expect(node.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node.Spec.SyncMode).To(Equal(DefaultPrivateNetworkSyncMode))
		Expect(node.Spec.Resources.CPU).To(Equal(DefaultPrivateNetworkNodeCPURequest))
		Expect(node.Spec.Resources.CPULimit).To(Equal(DefaultPrivateNetworkNodeCPULimit))
		Expect(node.Spec.Resources.Memory).To(Equal(DefaultPrivateNetworkNodeMemoryRequest))
		Expect(node.Spec.Resources.MemoryLimit).To(Equal(DefaultPrivateNetworkNodeMemoryLimit))
		Expect(node.Spec.Resources.Storage).To(Equal(DefaultPrivateNetworkNodeStorageRequest))
		Expect(node.Spec.Logging).To(Equal(DefaultLogging))
		// genesis defaulting
		Expect(node.Spec.Genesis.Coinbase).To(Equal(DefaultCoinbase))
		Expect(node.Spec.Genesis.MixHash).To(Equal(DefaultMixHash))
		Expect(node.Spec.Genesis.Difficulty).To(Equal(DefaultDifficulty))
		Expect(node.Spec.Genesis.GasLimit).To(Equal(DefaultGasLimit))
		Expect(node.Spec.Genesis.Nonce).To(Equal(DefaultNonce))
		Expect(node.Spec.Genesis.Timestamp).To(Equal(DefaultTimestamp))
		// forks defaulting
		Expect(node.Spec.Genesis.Forks.Homestead).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.DAO).To(BeNil())
		Expect(node.Spec.Genesis.Forks.EIP150).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP155).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP158).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Byzantium).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Constantinople).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Petersburg).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Istanbul).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.MuirGlacier).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Berlin).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.London).To(Equal(block0))
	})

	It("Should default nodes joining network with poa consensus", func() {
		node := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				Genesis: &Genesis{
					ChainID:   55555,
					NetworkID: 55555,
					Clique:    &Clique{},
				},
				Client: BesuClient,
				RPC:    true,
			},
		}

		node.Default()
		var block0 uint = 0
		Expect(node.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node.Spec.SyncMode).To(Equal(DefaultPrivateNetworkSyncMode))
		Expect(node.Spec.Hosts).To(Equal(DefaultOrigins))
		Expect(node.Spec.CORSDomains).To(Equal(DefaultOrigins))
		Expect(node.Spec.RPCPort).To(Equal(DefaultRPCPort))
		Expect(node.Spec.RPCAPI).To(Equal(DefaultAPIs))
		Expect(node.Spec.Resources.CPU).To(Equal(DefaultPrivateNetworkNodeCPURequest))
		Expect(node.Spec.Resources.CPULimit).To(Equal(DefaultPrivateNetworkNodeCPULimit))
		Expect(node.Spec.Resources.Memory).To(Equal(DefaultPrivateNetworkNodeMemoryRequest))
		Expect(node.Spec.Resources.MemoryLimit).To(Equal(DefaultPrivateNetworkNodeMemoryLimit))
		Expect(node.Spec.Resources.Storage).To(Equal(DefaultPrivateNetworkNodeStorageRequest))
		Expect(node.Spec.Logging).To(Equal(DefaultLogging))
		// genesis defaulting
		Expect(node.Spec.Genesis.Coinbase).To(Equal(DefaultCoinbase))
		Expect(node.Spec.Genesis.MixHash).To(Equal(DefaultMixHash))
		Expect(node.Spec.Genesis.Difficulty).To(Equal(DefaultDifficulty))
		Expect(node.Spec.Genesis.GasLimit).To(Equal(DefaultGasLimit))
		Expect(node.Spec.Genesis.Nonce).To(Equal(DefaultNonce))
		Expect(node.Spec.Genesis.Timestamp).To(Equal(DefaultTimestamp))
		// forks defaulting
		Expect(node.Spec.Genesis.Forks.Homestead).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.DAO).To(BeNil())
		Expect(node.Spec.Genesis.Forks.EIP150).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP155).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP158).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Byzantium).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Constantinople).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Petersburg).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Istanbul).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.MuirGlacier).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Berlin).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.London).To(Equal(block0))
		// clique defaulting
		Expect(node.Spec.Genesis.Clique.BlockPeriod).To(Equal(DefaultCliqueBlockPeriod))
		Expect(node.Spec.Genesis.Clique.EpochLength).To(Equal(DefaultCliqueEpochLength))
	})

	It("Should default nodes joining network with ibft2 consensus", func() {
		node := Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node-1",
			},
			Spec: NodeSpec{
				Genesis: &Genesis{
					ChainID:   55555,
					NetworkID: 55555,
					IBFT2:     &IBFT2{},
				},
				Client:  BesuClient,
				WS:      true,
				GraphQL: true,
			},
		}

		node.Default()
		var block0 uint = 0
		Expect(node.Spec.P2PPort).To(Equal(DefaultP2PPort))
		Expect(node.Spec.SyncMode).To(Equal(DefaultPrivateNetworkSyncMode))
		Expect(node.Spec.Hosts).To(Equal(DefaultOrigins))
		Expect(node.Spec.CORSDomains).To(Equal(DefaultOrigins))
		Expect(node.Spec.WSPort).To(Equal(DefaultWSPort))
		Expect(node.Spec.WSAPI).To(Equal(DefaultAPIs))
		Expect(node.Spec.GraphQLPort).To(Equal(DefaultGraphQLPort))
		Expect(node.Spec.Resources.CPU).To(Equal(DefaultPrivateNetworkNodeCPURequest))
		Expect(node.Spec.Resources.CPULimit).To(Equal(DefaultPrivateNetworkNodeCPULimit))
		Expect(node.Spec.Resources.Memory).To(Equal(DefaultPrivateNetworkNodeMemoryRequest))
		Expect(node.Spec.Resources.MemoryLimit).To(Equal(DefaultPrivateNetworkNodeMemoryLimit))
		Expect(node.Spec.Resources.Storage).To(Equal(DefaultPrivateNetworkNodeStorageRequest))
		Expect(node.Spec.Logging).To(Equal(DefaultLogging))
		// genesis defaulting
		Expect(node.Spec.Genesis.Coinbase).To(Equal(DefaultCoinbase))
		Expect(node.Spec.Genesis.MixHash).To(Equal(DefaultMixHash))
		Expect(node.Spec.Genesis.Difficulty).To(Equal(DefaultDifficulty))
		Expect(node.Spec.Genesis.GasLimit).To(Equal(DefaultGasLimit))
		Expect(node.Spec.Genesis.Nonce).To(Equal(DefaultNonce))
		Expect(node.Spec.Genesis.Timestamp).To(Equal(DefaultTimestamp))
		// forks defaulting
		Expect(node.Spec.Genesis.Forks.Homestead).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.DAO).To(BeNil())
		Expect(node.Spec.Genesis.Forks.EIP150).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP155).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.EIP158).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Byzantium).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Constantinople).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Petersburg).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Istanbul).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.MuirGlacier).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.Berlin).To(Equal(block0))
		Expect(node.Spec.Genesis.Forks.London).To(Equal(block0))
		// IBFT2 defaulting
		Expect(node.Spec.Genesis.IBFT2.BlockPeriod).To(Equal(DefaultIBFT2BlockPeriod))
		Expect(node.Spec.Genesis.IBFT2.EpochLength).To(Equal(DefaultIBFT2EpochLength))
		Expect(node.Spec.Genesis.IBFT2.RequestTimeout).To(Equal(DefaultIBFT2RequestTimeout))
		Expect(node.Spec.Genesis.IBFT2.MessageQueueLimit).To(Equal(DefaultIBFT2MessageQueueLimit))
		Expect(node.Spec.Genesis.IBFT2.DuplicateMessageLimit).To(Equal(DefaultIBFT2DuplicateMessageLimit))
		Expect(node.Spec.Genesis.IBFT2.FutureMessagesLimit).To(Equal(DefaultIBFT2FutureMessagesLimit))
		Expect(node.Spec.Genesis.IBFT2.FutureMessagesMaxDistance).To(Equal(DefaultIBFT2FutureMessagesMaxDistance))
	})
})
