package cornerstone

import (
	"context"
	"sync"
	"time"
)

type Context struct {
	c context.Context
	l *sync.RWMutex
	m *sync.Map
}

func NewContext() (ctx Context) {
	ctx = Context{
		c: context.Background(),
		l: &sync.RWMutex{},
		m: &sync.Map{},
	}
	return
}

// context native

func (ctx *Context) Done() (done <-chan struct{}) {
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	done = ctx.c.Done()
	return
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	deadline, ok = ctx.c.Deadline()
	return
}

func (ctx *Context) Value(key interface{}) (result interface{}) {
	keyS, ok := key.(string)
	if !ok {
		return
	}
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	result, _ = ctx.m.Load(keyS)
	return
}

func (ctx *Context) Err() (err error) {
	// TODO: implement this
	return
}

// map

func (ctx *Context) GetAllMap() (result map[string]interface{}) {
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	result = map[string]interface{}{}
	ctx.m.Range(func(k, v interface{}) bool {
		if kString, ok := k.(string); ok {
			result[kString] = v
		}
		return true
	})
	return
}

func (ctx *Context) Set(key string, value interface{}) {
	ctx.l.Lock()
	defer ctx.l.Unlock()
	ctx.m.Store(key, value)
}

func (ctx *Context) IsSet(key string) (ok bool) {
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	_, ok = ctx.m.Load(key)
	return
}

func (ctx *Context) value(key string) (result interface{}, ok bool) {
	ctx.l.RLock()
	defer ctx.l.RUnlock()
	result, ok = ctx.m.Load(key)
	return
}

func (ctx *Context) Get(key string) (result interface{}, ok bool) {
	result, ok = ctx.value(key)
	return
}

func (ctx *Context) GetBool(key string) (result bool, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(bool)
	return
}

func (ctx *Context) GetFloat64(key string) (result float64, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(float64)
	return
}

func (ctx *Context) GetInt(key string) (result int, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(int)
	return
}

func (ctx *Context) GetIntSlice(key string) (result []int, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.([]int)
	return
}

func (ctx *Context) GetString(key string) (result string, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(string)
	return
}

func (ctx *Context) GetStringSlice(key string) (result []string, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.([]string)
	return
}

func (ctx *Context) GetStringMap(key string) (result map[string]interface{}, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(map[string]interface{})
	return
}

func (ctx *Context) GetStringMapString(key string) (result map[string]string, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(map[string]string)
	return
}

func (ctx *Context) GetTime(key string) (result time.Time, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(time.Time)
	return
}

func (ctx *Context) GetDuration(key string) (result time.Duration, ok bool) {
	resultI, ok := ctx.value(key)
	if resultI == nil {
		return
	}
	result, ok = resultI.(time.Duration)
	return
}

// directive

func (ctx *Context) SetCancel() (cancelFunc context.CancelFunc) {
	ctx.l.Lock()
	defer ctx.l.Unlock()
	ctx.c, cancelFunc = context.WithCancel(ctx.c)
	return
}

func (ctx *Context) SetTimeout(timeout time.Duration) (cancelFunc context.CancelFunc) {
	ctx.l.Lock()
	defer ctx.l.Unlock()
	ctx.c, cancelFunc = context.WithTimeout(ctx.c, timeout)
	return
}

func (ctx *Context) SetDeadline(deadline time.Time) (ok bool, cancelFunc context.CancelFunc) {
	if deadline.Before(time.Now()) {
		return
	}
	ctx.l.Lock()
	defer ctx.l.Unlock()
	ok = true
	ctx.c, cancelFunc = context.WithDeadline(ctx.c, deadline)
	return
}

// copy

func copySyncMap(src *sync.Map) (dst *sync.Map) {
	dst = &sync.Map{}
	src.Range(func(k, v interface{}) bool {
		vSyncMap, ok := v.(sync.Map)
		if ok {
			dst.Store(k, copySyncMap(&vSyncMap))
		} else {
			dst.Store(k, v)
		}
		return true
	})
	return
}

func (ctx *Context) CopyContext() (result Context) {
	result = Context{
		c: context.Background(),
		l: &sync.RWMutex{},
		m: copySyncMap(ctx.m),
	}
	return
}

func (ctx *Context) WithCancel() (result Context, cancelFunc context.CancelFunc) {
	new, cancelFunc := context.WithCancel(ctx.c)
	result = Context{
		c: new,
		l: &sync.RWMutex{},
		m: copySyncMap(ctx.m),
	}
	return
}

func (ctx *Context) WithTimeout(timeout time.Duration) (result Context, cancelFunc context.CancelFunc) {
	new, cancelFunc := context.WithTimeout(ctx.c, timeout)
	result = Context{
		c: new,
		l: &sync.RWMutex{},
		m: copySyncMap(ctx.m),
	}
	return
}

func (ctx *Context) WithDeadline(deadline time.Time) (ok bool, result Context, cancelFunc context.CancelFunc) {
	if deadline.Before(time.Now()) {
		return
	}
	ok = true
	new, cancelFunc := context.WithDeadline(ctx.c, deadline)
	result = Context{
		c: new,
		l: &sync.RWMutex{},
		m: copySyncMap(ctx.m),
	}
	return
}
