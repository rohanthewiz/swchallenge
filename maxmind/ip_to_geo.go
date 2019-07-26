package maxmind

import (
	"errors"
	"swchallenge/geo"
	"sync"
)

var MaxMindMap maxMindMapWrap

func init() {
	MaxMindMap.maxMindMap = make(map[string]geo.Geo, 16) // preallocate a few to start
}

type maxMindMapWrap struct {
	maxMindMap map[string]geo.Geo
	mu         sync.RWMutex
}

func (mm maxMindMapWrap) GetGeo(ip string) (geo.Geo, error) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()
	g, ok := mm.maxMindMap[ip]
	if !ok {
		return g, errors.New("not in map")
	}
	return g, nil
}

func (mm maxMindMapWrap) SetGeo(ip string, g geo.Geo) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.maxMindMap[ip] = g
}

// On the todo of course is a way to expire map entries
