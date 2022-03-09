package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/lovemew67/public-misc/cornerstone"
)

const (
	opPing   = "PING"
	opPong   = "PONG"
	opSet    = "SET"
	opGet    = "GET"
	opDel    = "DEL"
	opOk     = "OK"
	opExists = "EXISTS"
)

type Config struct {
	Host    string
	RedisDB int
}

type Pool struct {
	ctx  cornerstone.Context
	pool *redis.Pool
	conf *Config
}

func NewPool(conf *Config) (*Pool, error) {
	funcName := "[redis][NewPool]"

	ctx := cornerstone.NewContext()
	host := conf.Host
	db := conf.RedisDB
	timeout := 10 * time.Minute
	maxIdle := 5
	maxActive := 20
	idleTimeout := 3 * time.Minute
	cornerstone.Infof(ctx, "%s init redis with config in ctx", funcName)

	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			cornerstone.Debugf(ctx, "%s redis dialing, host: %s", funcName, host)
			c, err := redis.Dial("tcp", host,
				redis.DialConnectTimeout(timeout),
				redis.DialWriteTimeout(timeout),
				redis.DialReadTimeout(timeout),
				redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do(opPing)
			return err
		},
	}

	result := &Pool{
		ctx,
		pool,
		conf,
	}
	err := result.Ping()
	if err != nil {
		result = nil
	}

	return result, err
}

func (p *Pool) Ping() error {
	c := p.pool.Get()
	defer c.Close()
	resp, err := redis.String(c.Do(opPing))
	if err != nil {
		return err
	}
	if resp != opPong {
		return fmt.Errorf("err resp:%s", resp)
	}
	return nil
}

func (p *Pool) Set(key string, value interface{}) error {
	c := p.pool.Get()
	defer c.Close()
	resp, err := redis.String(c.Do(opSet, key, value))
	if err != nil {
		return err
	}
	if resp != opOk {
		return fmt.Errorf("err resp:%s", resp)
	}
	return nil
}

func (p *Pool) GetBytes(key string) (value []byte, err error) {
	c := p.pool.Get()
	defer c.Close()
	value, err = redis.Bytes(c.Do(opGet, key))
	return
}

func (p *Pool) Delete(key string) error {
	c := p.pool.Get()
	defer c.Close()
	resp, err := redis.Int(c.Do(opDel, key))
	if err != nil {
		return err
	}
	if resp != 1 {
		return fmt.Errorf("err resp:%d", resp)
	}
	return nil
}

func (p *Pool) Exists(key string) (value bool, err error) {
	c := p.pool.Get()
	defer c.Close()
	value, err = redis.Bool(c.Do(opExists, key))
	return
}
