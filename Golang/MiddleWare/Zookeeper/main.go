package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"encoding/json"
)

type ServiceNode struct{ // zookeeper node object
	Name string `json: "name"`
	Host string `json: "host"`
	Port int `json: "port"`
}

type SdClient struct{ // Client object
	zkServers []string // Multiple nodes
	zkRoot string // Root node
	conn *zk.Conn
}

func (s *SdClient) ensureRoot() error{
	exists, _, err := s.conn.Exists(s.zkRoot) // check if root exists
	if err != nil{
		return err
	}
	if !exists { // if node does not exists, create one
		_, err := s.conn.Create(s.zkRoot, []byte(""), 0, zk.WorldACL(zk.PermAll)) // create zkRoot with no data 
		if err != nil{
			return err
		}
	}
	return nil
}

func (s *SdClient) ensurePath(name string) error { // if service node not exists, create it
	path := s.zkRoot + "/" + name
	exists, _ , err := s.conn.Exists(path)
	if err != nil {return err}
	if !exists {
		_, err := s.conn.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
		if err != nil{return err}
	}
	return nil
}

func NewClient(zkServer []string, zkRoot string, timeout int) (*SdClient, error){
	client := &SdClient{} 
	client.zkServers = zkServer
	client.zkRoot = zkRoot // simply fill in data
	conn, _ ,err := zk.Connect(zkServer, time.Duration(timeout) * time.Second)  // connect to zookeeper cluster
	if err != nil {return nil, err}
	client.conn = conn
	err = client.ensureRoot() // ensure root
	if err != nil {
		client.conn.Close()
		return nil, err
	}
	return client, nil
}

// Protected Node 临时节点
// 如果客户端创建了一个节点，但是刚刚的server突然断了，那么临时节点可以帮助我们找回节点
// 之后在另一个server上创建
func (s *SdClient) Register(node *ServiceNode) error{
	if err := s.ensurePath(node.Name); err != nil{ // 
		return err
	}
	path := s.zkRoot + "/" + node.Name + "/n"
	data, err := json.Marshal(node) // convert node struct into json
	if err != nil {return err}
	_, err = s.conn.CreateProtectedEphemeralSequential(path, data, zk.WorldACL(zk.PermAll)) // Create Protected Node
	if err != nil {return err}
	return nil
}

func (s *SdClient) GetNodes(name string) ([]*ServiceNode, error){
	path := s.zkRoot + "/" + name
	childs, _, err := s.conn.Children(path) 
	if err != nil{
		if err == zk.ErrNoNode {
			return []*ServiceNode{}, err
		}
		return nil, err
	}
	nodes := []*ServiceNode{}
	for _, child := range childs{
		childPath := path + "/" + child
		data, _, err := s.conn.Get(childPath)
		if err != nil{
			if err == zk.ErrNoNode {
				return []*ServiceNode{}, err
			}
			return nil, err
		}
		node := &ServiceNode{}
		err = json.Unmarshal(data, node)
		if err != nil{
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func zkOperationTest(servers []string) error{
	fmt.Println("ZooKeeper Operation Test")
	conn, _, err := zk.Connect(servers, time.Second * 5)
	if err != nil{return err}
	defer conn.Close()
	path := "/zk_test"
	data := []byte("Testing_Data")
	_, err = conn.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	if err != nil{return err}

	exists, _, err := conn.Exists(path)
	if err != nil{return err}
	if exists{
		fmt.Println("Path " + path + " created successfully")
	}

	data, str, err := conn.Get(path)
	if err != nil{return err}
	fmt.Printf("Path[%s]=[%s].\n", path, data)
	fmt.Printf("state:\n")

	err = conn.Delete(path, str.Version)
	if err != nil{return err}

	fmt.Println("Node has been deleted")
	return nil
}

func main(){
	servers := []string{"192.168.157.20:2181", "192.168.157.20:2182", "192.168.157.20:2183"}
	err := zkOperationTest(servers)
	if err != nil{return}

	client, err := NewClient(servers, "/api", 10) // for example, this is an API client
	if err != nil{
		client.conn.Close()
	}
	defer client.conn.Close()
	node1 := &ServiceNode{"user", "127.0.0.1", 2889}
	node2 := &ServiceNode{"user", "127.0.0.1", 2890}
	node3 := &ServiceNode{"user", "127.0.0.1", 2891}
	if err := client.Register(node1); err != nil{
		panic(err)
	}
	if err := client.Register(node2); err != nil{
		panic(err)
	}
	if err := client.Register(node3); err != nil{
		panic(err)
	}
	nodes, err := client.GetNodes("user")
	if err != nil{
		panic(err)
	}
	for _, node := range nodes{
		fmt.Println(node)
	}
}
