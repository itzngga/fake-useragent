package useragent

import (
	"github.com/puzpuzpuz/xsync/v3"
	"math/rand/v2"
)

type useragent struct {
	data     *xsync.MapOf[string, []string]
	dataLen  *xsync.MapOf[int64, string]
	totalLen *xsync.Counter
}

var (
	UA = useragent{data: xsync.NewMapOf[string, []string](), dataLen: xsync.NewMapOf[int64, string](), totalLen: xsync.NewCounter()}
)

func (u *useragent) Get(key string) []string {
	data, _ := u.data.Load(key)
	return data
}

func (u *useragent) GetAll() map[string][]string {
	var data = make(map[string][]string)
	u.data.Range(func(key string, value []string) bool {
		data[key] = value
		return true
	})
	return data
}

func (u *useragent) GetRandom(key string) string {
	browser := u.Get(key)
	length := len(browser)
	if length <= 0 {
		return ""
	}

	n := rand.IntN(length)
	if n > length {
		return ""
	}

	return browser[n]
}

func (u *useragent) GetAllRandom() string {
	if val, ok := u.dataLen.Load(rand.Int64N(u.totalLen.Value())); ok {
		return u.GetRandom(val)
	}

	return ""
}

func (u *useragent) Set(key, value string) {
	data, ok := u.data.Load(key)
	if ok {
		data = append(data, value)
	} else {
		data = []string{value}
	}

	u.data.Store(key, data)
	u.resetLen()
}

func (u *useragent) SetData(data map[string][]string) {
	for k, v := range data {
		u.data.Store(k, v)
	}
	u.resetLen()
}

func (u *useragent) resetLen() {
	u.dataLen.Clear()
	u.totalLen.Reset()
	u.totalLen.Add(0)

	u.data.Range(func(key string, value []string) bool {
		u.dataLen.Store(u.totalLen.Value(), key)
		u.totalLen.Inc()
		return true
	})
}
