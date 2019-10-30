package main

import (
	"fmt"
	"github.com/ArcturusZhang/module-demo-test/services/aznet/mgmt/2019-10-01/aznet"
	"github.com/ArcturusZhang/module-demo-test/services/bznet/mgmt/2019-11-01/bznet"
)

func main() {
	// Use Aznet
	azClient := aznet.AzNetsClient{}
	fmt.Println(azClient.Get())
	// Use Bznet
	bzClient := bznet.BzNetClient{}
	fmt.Println(bzClient.Get())
}
