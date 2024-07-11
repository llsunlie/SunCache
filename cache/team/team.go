package team

import (
	"SunCache/cache/chatTool"
	"SunCache/cache/consistentHash"
	iu "SunCache/cache/interface"
	"SunCache/cache/log"
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Team struct {
	memLocker sync.RWMutex
	members   map[string]iu.MemberIU

	chatServer   *chatTool.Server
	addrLocker   sync.RWMutex
	addressTable *consistentHash.ConsistentHash
	chatClients  map[string]*chatTool.Client
}

func NewTeam(socket string, addresses ...string) (team *Team) {
	team = &Team{
		members: make(map[string]iu.MemberIU),

		chatServer: chatTool.NewServer(socket),
		addressTable: consistentHash.NewConsistentHash(
			consistentHash.DefaultMultiple,
			nil,
		),
		chatClients: make(map[string]*chatTool.Client),
	}
	team.AddAddress(socket)
	if addresses != nil {
		team.AddAddress(addresses...)
	}
	team.chatServer.RegisterTeam(team)
	return
}

func (t *Team) GetChatServerSocket() (address string) {
	return t.chatServer.GetSocket()
}

func (t *Team) AddMember(members ...iu.MemberIU) {
	t.memLocker.Lock()
	defer t.memLocker.Unlock()
	for _, m := range members {
		if _, ok := t.members[m.GetName()]; !ok {
			t.members[m.GetName()] = m
			m.RegisterTeam(t)
		}
	}
}

func (t *Team) GetMember(name string) (member iu.MemberIU) {
	t.memLocker.RLock()
	defer t.memLocker.RUnlock()
	return t.members[name]
}

func (t *Team) AddAddress(addresses ...string) {
	t.addrLocker.Lock()
	defer t.addrLocker.Unlock()
	for _, a := range addresses {
		if _, ok := t.chatClients[a]; !ok {
			t.addressTable.Add(a)
			t.chatClients[a] = chatTool.NewClient(a)
		}
	}
}

func (t *Team) GetAddress(key string) (address string) {
	t.addrLocker.RLock()
	defer t.addrLocker.RUnlock()
	return t.addressTable.Get(key)
}

func (t *Team) GetAddressAndClient(key string) (address string, client *chatTool.Client) {
	t.addrLocker.RLock()
	defer t.addrLocker.RUnlock()
	address = t.addressTable.Get(key)
	client = t.chatClients[address]
	return
}

func (t *Team) GetValueFromRemote(member, key string) (value []byte, err error) {
	address, client := t.GetAddressAndClient(key)
	if address == t.chatServer.GetSocket() {
		return nil, fmt.Errorf("[%v] The key (%v) need load from source", t.chatServer.GetSocket(), key)
	}
	log.Info("[%v] GetValueFromRemote [%v]", t.chatServer.GetSocket(), client.GetServerSocket())
	req := &chatTool.Request{
		Member: member,
		Key:    key,
	}
	var res *chatTool.Response
	res, err = client.Get(context.Background(), req)
	if err != nil {
		return
	}
	value = res.Value
	return
}

func (t *Team) RunChatServer() {
	log.Info("[%v] ChatServer is listening\n", t.chatServer.GetSocket())
	t.chatServer.Run()
}

func (t *Team) StartHttpServer(socket string) {
	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		member := r.URL.Query().Get("member")
		key := r.URL.Query().Get("key")
		log.Info("[%v] %v", socket, r.RequestURI)
		m := t.GetMember(member)
		value, err := m.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(value.Copy())
	}))
	log.Info("[%v] HttpServer is listening", socket)
	log.Errorln(http.ListenAndServe(socket, nil))
}
