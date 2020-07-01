package controllers

import (
	"fmt"

	ethereumv1alpha1 "github.com/mfarghaly/kotal/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ethereum client arguments", func() {

	var bootnodes []string
	rinkeby := "rinkeby"
	bootnode := "enode://publickey@ip:port"
	coinbase := ethereumv1alpha1.EthereumAddress("0x5A0b54D5dc17e0AadC383d2db43B0a0D3E029c4c")
	nodekey := ethereumv1alpha1.PrivateKey("0x608e9b6f67c65e47531e08e8e501386dfae63a540fa3c48802c8aad854510b4e")

	cases := []struct {
		title     string
		bootnodes []string
		network   *ethereumv1alpha1.Network
		result    []string
	}{
		{
			"node joining rinkeby",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name: "node-1",
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuDataPath,
				PathBlockchainData,
			},
		},
		{
			"bootnode joining rinkeby",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Bootnode: true,
							Nodekey:  nodekey,
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuNodePrivateKey,
				BesuDataPath,
				PathBlockchainData,
			},
		},
		{
			"bootnode joining rinkeby with rpc",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Bootnode: true,
							Nodekey:  nodekey,
							RPC:      true,
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuNodePrivateKey,
				BesuDataPath,
				PathBlockchainData,
				BesuRPCHTTPEnabled,
			},
		},
		{
			"bootnode joining rinkeby with rpc settings",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Bootnode: true,
							Nodekey:  nodekey,
							RPC:      true,
							RPCPort:  8599,
							RPCAPI: []ethereumv1alpha1.API{
								ethereumv1alpha1.ETHAPI,
								ethereumv1alpha1.Web3API,
								ethereumv1alpha1.NetworkAPI,
							},
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuNodePrivateKey,
				BesuDataPath,
				PathBlockchainData,
				BesuRPCHTTPEnabled,
				BesuRPCHTTPPort,
				"8599",
				BesuRPCHTTPAPI,
				"eth,web3,net",
			},
		},
		{
			"bootnode joining rinkeby with rpc, ws settings",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Bootnode: true,
							Nodekey:  nodekey,
							RPC:      true,
							RPCHost:  "0.0.0.0",
							RPCPort:  8599,
							RPCAPI: []ethereumv1alpha1.API{
								ethereumv1alpha1.ETHAPI,
								ethereumv1alpha1.Web3API,
								ethereumv1alpha1.NetworkAPI,
							},
							WS:     true,
							WSHost: "127.0.0.1",
							WSPort: 8588,
							WSAPI: []ethereumv1alpha1.API{
								ethereumv1alpha1.Web3API,
								ethereumv1alpha1.ETHAPI,
							},
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuNodePrivateKey,
				BesuDataPath,
				PathBlockchainData,
				BesuRPCHTTPEnabled,
				BesuRPCHTTPPort,
				"8599",
				BesuRPCHTTPHost,
				"0.0.0.0",
				BesuRPCHTTPAPI,
				"eth,web3,net",
				BesuRPCWSEnabled,
				BesuRPCWSHost,
				"127.0.0.1",
				BesuRPCWSPort,
				"8588",
				BesuRPCWSAPI,
				"web3,eth",
			},
		},
		{
			"bootnode joining rinkeby with rpc, ws, graphql settings and cors domains",
			bootnodes,
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join: rinkeby,
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Bootnode: true,
							Nodekey:  nodekey,
							RPC:      true,
							RPCHost:  "0.0.0.0",
							RPCPort:  8599,
							RPCAPI: []ethereumv1alpha1.API{
								ethereumv1alpha1.ETHAPI,
								ethereumv1alpha1.Web3API,
								ethereumv1alpha1.NetworkAPI,
							},
							CORSDomains: []string{"cors.example.com"},
							WS:          true,
							WSHost:      "127.0.0.1",
							WSPort:      8588,
							WSAPI: []ethereumv1alpha1.API{
								ethereumv1alpha1.Web3API,
								ethereumv1alpha1.ETHAPI,
							},
							GraphQL:     true,
							GraphQLHost: "127.0.0.2",
							GraphQLPort: 8511,
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuNetwork,
				rinkeby,
				BesuNodePrivateKey,
				BesuDataPath,
				PathBlockchainData,
				BesuRPCHTTPCorsOrigins,
				BesuRPCHTTPEnabled,
				BesuRPCHTTPPort,
				"8599",
				BesuRPCHTTPHost,
				"0.0.0.0",
				BesuRPCHTTPAPI,
				"eth,web3,net",
				BesuRPCWSEnabled,
				BesuRPCWSHost,
				"127.0.0.1",
				BesuRPCWSPort,
				"8588",
				BesuRPCWSAPI,
				"web3,eth",
				BesuGraphQLHTTPEnabled,
				BesuGraphQLHTTPHost,
				"127.0.0.2",
				BesuGraphQLHTTPPort,
				"8511",
				BesuGraphQLHTTPCorsOrigins,
				"cors.example.com",
			},
		},
		{
			"miner node of private network that connects to bootnode",
			[]string{bootnode},
			&ethereumv1alpha1.Network{
				Spec: ethereumv1alpha1.NetworkSpec{
					Join:    rinkeby,
					Genesis: &ethereumv1alpha1.Genesis{},
					Nodes: []ethereumv1alpha1.Node{
						{
							Name:     "node-1",
							Miner:    true,
							Coinbase: coinbase,
						},
					},
				},
			},
			[]string{
				BesuNatMethod,
				BesuDataPath,
				PathBlockchainData,
				BesuBootnodes,
				bootnode,
				BesuMinerEnabled,
				BesuMinerCoinbase,
				string(coinbase),
			},
		},
	}

	for _, c := range cases {
		func() {
			cc := c
			It(fmt.Sprintf("Should create correct client arguments for %s", cc.title), func() {
				cc.network.Default()
				args := reconciler.createArgsForClient(&cc.network.Spec.Nodes[0], cc.network, cc.bootnodes)
				Expect(args).To(ContainElements(cc.result))
			})
		}()
	}

})
