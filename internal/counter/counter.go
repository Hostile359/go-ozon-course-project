package counter

import (
	"expvar"
	"strconv"
	"sync"
)

var (
	successReqCounter *counter
	failReqCounter *counter
	validErrCounter *counter
	internalErrCounter *counter
	cacheHitCounter *counter
	cacheMissCounter *counter
)

type counter struct {
	cnt int
	m   *sync.RWMutex
}

func (c *counter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.cnt++
}

func (c *counter) String() string {
	c.m.RLock()
	defer c.m.RUnlock()
	return strconv.FormatInt(int64(c.cnt), 10)
}

func IncSuccessReq() {
	successReqCounter.Inc()
}

func IncFailReq() {
	failReqCounter.Inc()
}

func IncValidErr() {
	validErrCounter.Inc()
}

func IncInternalErr() {
	internalErrCounter.Inc()
}

func IncCacheHit() {
	cacheHitCounter.Inc()
}

func IncCacheMiss() {
	cacheMissCounter.Inc()
}

func init() {
	successReqCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("SuccessRequests", successReqCounter)

	failReqCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("FailRequests", failReqCounter)

	validErrCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("ValidErrors", validErrCounter)

	internalErrCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("InternalErrors", internalErrCounter)

	cacheHitCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("CacheHits", cacheHitCounter)

	cacheMissCounter = &counter{m: &sync.RWMutex{}}
	expvar.Publish("CacheMisses", cacheMissCounter)
}
