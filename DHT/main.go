package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
)

// 第一个方法是注册方法
type BTServer struct {
	// 储存一部分用户CID
	AllNodeInfos []NodeInfo // 其实不需要保存所有的节点信息的
}

func NewBtServer() *BTServer {
	b := BTServer{}
	return &b
}

func (b BTServer)BootNeighbors(targetNode NodeID) []Bucket {
	// 距离指数上升 32, 64, 128, 256, 512, 1024
	// findIntervalDynamic 可以找出 a 在哪个区间
	findIntervalDynamic := func(a int)int{
		if a < 32 {
			return 0
		}
		return int(math.Log2(float64(a))) -5 + 1
	}

	result := make([]Bucket, 0, MaxBucketCnt)
	for _, node := range b.AllNodeInfos{
		targetDistance := CompareDistance(targetNode, node)
		if findIntervalDynamic(targetDistance) >= len(result)-1 {
			result = slices.Grow(result, targetLen - len(result))
		}
	}

	return 
}

type Bucket struct {
	FromDistance int
	ToDistance   int
	NodeInfos    []NodeInfo
}

// NodeInfo 用于在节点间传递轻量级信息
type NodeInfo struct {
	ID   NodeID
	IP   string
	Port string
}

const IDLength = 32
const MaxBucketCnt = 20
const MaxBucketPeerIdCnt = 20

type NodeID [IDLength]byte // [32]byte

func NewNodeID(data string) NodeID {
	// sha256.Sum256 固定返回长度为 32
	// 	// mock ipAddress + port ..

	sum := sha256.Sum256([]byte(data))
	return NodeID(sum)
}

// 计算两个节点之间的异或距离
func CompareDistance(Node1, Node2 NodeID) int {
	return bytes.Compare(Node1[:], Node2[:])
}

type Node struct {
	NodeId          NodeID
	ContentRouting  []Bucket          // router table
	ServerRef       *BTServer         // 持有 Server 的引用，用于模拟网络请求
	ProviderStore   map[string]string //  index, 存储某一个 cid文件哪个用户拥有, cid --> peerId
	LocalBlockStore map[string]string // local file, mock
}

func NewNode(ip, port string, server *BTServer) *Node {
	node := Node{}
	node.NodeId = NewNodeID(fmt.Sprintf("%s:%s", ip, port))
	node.ServerRef = server

	return node.BootStrap(ip, port)
}

func (node Node) BootStrap(ip, port string) *Node {
	// 都是网络操作, 内存快速mock实现
	newNodeInfo := NodeInfo{node.NodeId, ip, port}
	node.ServerRef.AllNodeInfos = append(node.ServerRef.AllNodeInfos, newNodeInfo)
	node.ContentRouting = 

	return &node
}
