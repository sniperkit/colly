package cmdb

import (
	"net"
	"errors"
)

type Node interface {
	Address() net.IPAddr // unique id
	Tags() map[string]string
	GetTag(tagKey string) string
}

// NodeIterator return a batch of nodes on every call, return nil when no more data to read
type NodeIterator func() []Node

var NilNodeIterator NodeIterator = func() []Node {
	return nil
}

var GetNode = func(addr net.IPAddr) Node {
	return nil
}

var ListAllNodes = func() NodeIterator {
	return NilNodeIterator
}

var SearchNodes = func(tagKV ...string) NodeIterator {
	return NilNodeIterator
}

type DistributionPlan []DistributionStep

type DistributionStep struct {
	ToNode    net.IP
	NextSteps []DistributionStep
}

var GenerateDistributionPlan = func(srcNode net.IP, replayNodes []net.IP, destNodes net.IP) (DistributionPlan, error) {
	return nil, errors.New("not implemented")
}
