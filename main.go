package main

import "github.com/anthdm/projectx/network"

// Server
// Transport => tcp,udp
// Block
// Tx
// Keypair

// 自定义的地址类型，实现了 net.Addr 接口
type CustomAddr struct {
	network string
	address string
}

// // 实现 Addr 接口的 Network 方法
// func (ca *CustomAddr) Network() string {
// 	return ca.network
// }

// // 实现 Addr 接口的 String 方法
// func (ca *CustomAddr) String() string {
// 	return ca.address
// }

func main() {
	// 创建一个本地地址
	localAddr := &CustomAddr{
		network: "LOCAL",
		address: "127.0.0.1:8080",
	}
	trLocal := network.NewLocalTransport(localAddr)

	// opts := network.ServerOpts{
	// 	Transports:x
	// }
}

func Network() string {
	return "LOCAL"
}
