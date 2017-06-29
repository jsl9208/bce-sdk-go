package blb

import (
	"fmt"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestCreateTCPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
		BackendPort:    8080,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateTCPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUDPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateUDPListenerArgs{
		LoadBalancerId:    "lb-f5d263e5",
		ListenerPort:      8888,
		BackendPort:       8888,
		Scheduler:         "LeastConnection",
		HealthCheckString: "hello",
	}
	err := blbClient.CreateUDPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateHTTPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateHTTPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		ListenerPort:   8899,
		BackendPort:    8899,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateHTTPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeTCPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
	}
	list, err := blbClient.DescribeTCPListener(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}

func TestDescribeUDPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeUDPListenerArgs{
		LoadBalancerId: "lb-07ab7a1d",
		// ListenerPort:   80,
	}
	list, err := blbClient.DescribeUDPListener(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("DescribeUDPListener", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}
func TestUpdateTCPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &UpdateTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
		BackendPort:    999,
	}
	err := blbClient.UpdateTCPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUDPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &UpdateUDPListenerArgs{
		LoadBalancerId:    "lb-f5d263e5",
		ListenerPort:      8888,
		BackendPort:       8019,
		Scheduler:         "RoundRobin",
		HealthCheckString: "A",
	}
	err := blbClient.UpdateUDPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteListeners(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DeleteListenersArgs{
		LoadBalancerId: "lb-e5b33752",
		PortList:       []int{8088},
	}
	err := blbClient.DeleteListeners(args)
	if err != nil {
		t.Error(err)
	}
}
