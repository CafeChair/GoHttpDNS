package main

import (
	"github.com/hoisie/redis"
)

/*
直接从redis中获取解析结果(首选)
*/
func resolveFromRedis(domain,address string) (string,error) {
	pool := newRedisPool(address)
	conn := pool.Get()
	defer conn.Close()
	res,err := redis.String(conn.Do("GET",domain))
	if err != nil {
		return "", err
	}
	return res,nil
}

/*
解析到的结果存到redis中
*/
func storageToRedis(domain,ips,address string) (bool,error) {
	pool := newRedisPool(address)
	conn := pool.Get()
	defer conn.Close()
	ok,err := redis.String(conn.Do("SET",domain,ips))
	if err != nil {
		return false,err
	}
	return true,nil
}

//初始化redis pool
func newRedisPool(address string) *redis.Pool {
	RedisPool := &redis.Pool{
		// Maximum number of idle connections in the pool
		MaxIdle:10,
		//Dial is for creating and configure connection
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp",address)
			if err != nil {
				return nil,err
			}
			return conn,err
		},
	}
	return RedisPool
}