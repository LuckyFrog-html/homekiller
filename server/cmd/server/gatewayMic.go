package main

import "server/gatewayMicroservice/cmd/gateway"

func main() {
	gateway.Start()
	gateway.TestConnections()
}
