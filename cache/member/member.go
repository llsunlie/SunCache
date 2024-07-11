package member

import (
	"SunCache/cache/core"
	"SunCache/cache/core/lru"
	iu "SunCache/cache/interface"
	"SunCache/cache/log"
	"fmt"
)

type Member struct {
	name         string
	localCache   core.Cache
	gatherer     *Gatherer
	team         iu.TeamIU
	sourceGetter SourceGetter
}

func NewMember(name string, maxBytes int64, sourceGetter SourceGetter, cache core.Cache) (member *Member) {
	if sourceGetter == nil {
		panic("nil getter")
	}
	member = &Member{
		name:         name,
		localCache:   cache,
		gatherer:     NewGatherer(),
		sourceGetter: sourceGetter,
	}
	if member.localCache == nil {
		member.localCache = lru.NewSafeCache(maxBytes)
	}
	return
}

func (m *Member) RegisterTeam(team iu.TeamIU) {
	if m.team != nil {
		panic("Member RegisterPeer() called more than once")
	}
	m.team = team
}

func (m *Member) GetName() (name string) {
	return m.name
}

func (m *Member) Get(key string) (value iu.ByteViewIU, err error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("empty key")
	}
	if v, ok := m.getFromLocal(key); ok {
		log.Info("[%v] Cache hit", m.team.GetChatServerSocket())
		log.Info("[%v] Cache useBytes/maxBytes: %v/%v", m.team.GetChatServerSocket(), m.localCache.UseBytes(), m.localCache.MaxBytes())
		return v, nil
	} else {
		var v interface{}
		v, err = m.gatherer.Do(key, func() (interface{}, error) {
			if value, err = m.getFromPeers(key); err == nil {
				return value, nil
			}
			return m.getFromSource(key)
		})
		if err == nil {
			return v.(ByteView), nil
		}
		return
	}
}

func (m *Member) getFromLocal(key string) (value ByteView, ok bool) {
	if v, ok := m.localCache.Get(key); ok {
		return v.(ByteView), true
	}
	return ByteView{}, false
}

func (m *Member) getFromPeers(key string) (value ByteView, err error) {
	if m.team != nil {
		var v []byte
		v, err = m.team.GetValueFromRemote(m.name, key)
		if err != nil {
			return
		}
		return ByteView{value: v}, nil
	}
	return ByteView{}, fmt.Errorf("peers is nil")
}

func (m *Member) getFromSource(key string) (value ByteView, err error) {
	bytes, err := m.sourceGetter.Get(key)
	if err != nil {
		return
	}
	value = ByteView{value: cloneBytes(bytes)}
	m.localCache.Add(key, value)
	return
}
