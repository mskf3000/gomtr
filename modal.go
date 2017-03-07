package gomtr

import (
	"fmt"
	"github.com/gogather/safemap"
)

// parsed ttl item data
type TTLData struct {
	TTLID  int
	ipType string
	ip     string
	time   int64
	err    error
}

// task
type mtrTask struct {
	id       int64
	callback func(interface{})
	ttls     int
	ttlData  *safemap.SafeMap // item is ttlData, key is ttl
}

func (mt *mtrTask) save(ttl int, data *TTLData) {
	mt.ttlData.Put(fmt.Sprintf("%d", ttl), data)
}

func (mt *mtrTask) check() bool {
	for i := 1; i <= mt.ttls; i++ {
		_, ok := mt.ttlData.Get(fmt.Sprintf("%d", i))
		if !ok {
			return false
		}
	}
	return true
}

func (mt *mtrTask) clear() {
	for key, _ := range mt.ttlData.GetMap() {
		mt.ttlData.Remove(key)
	}
}