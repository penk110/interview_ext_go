package raft_node

import (
	"net"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/config"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
)

func NewRaftNode(serviceId string, serviceName string, raftConfig *config.RaftConfig, cacheClient cache.Cache, nodes []*config.Node) (*raft.Raft, error) {
	defaultConfig := raft.DefaultConfig()
	defaultConfig.LocalID = raft.ServerID(serviceId)

	hcLogOpt := &hclog.LoggerOptions{
		Name:            serviceName,
		Level:           hclog.LevelFromString("DEBUG"),
		Output:          os.Stdout,
		Mutex:           &sync.Mutex{},
		JSONFormat:      true,
		IncludeLocation: true,
		TimeFormat:      "2006-01-02 15:04:05",
	}
	defaultConfig.Logger = hclog.New(hcLogOpt)

	// logStore
	logStore, err := raftboltdb.NewBoltStore(raftConfig.LogStore)
	if err != nil {
		return nil, err
	}

	// snapshotStore
	snapshotStore := raft.NewDiscardSnapshotStore()

	// 节点信息
	stableStore, err := raftboltdb.NewBoltStore(raftConfig.StableStore)
	if err != nil {
		return nil, err
	}

	// 节点之间的通信
	addr, err := net.ResolveTCPAddr("tcp", raftConfig.Transport)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second*10, os.Stdout)
	if err != nil {
		return nil, err
	}

	// 自定义FSM
	fsm := NewFSM(cacheClient)

	raftNode, err := raft.NewRaft(defaultConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}
	var servers []raft.Server
	for _, node := range nodes {
		servers = append(servers, raft.Server{ID: node.ID, Address: node.Address})
	}

	configuration := raft.Configuration{
		Servers: servers,
	}

	raftNode.BootstrapCluster(configuration)

	return raftNode, nil
}
