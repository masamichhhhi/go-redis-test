package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Connection() redis.Conn {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

type Data struct {
	Key   string
	Value string
}

func Set(key, value string, c redis.Conn) string {
	res, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		panic(err)
	}
	return res
}

func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}

// 複数登録
func Mset(datas []Data, c redis.Conn) {
	var query []interface{}
	for _, v := range datas {
		query = append(query, v.Key, v.Value)
	}
	fmt.Println(query)

	c.Do("MSET", query...)
}

// 複数取得
func Mget(keys []string, c redis.Conn) []string {
	var query []interface{}
	for _, v := range keys {
		query = append(query, v)
	}
	fmt.Println("MGET query:", query)

	res, err := redis.Strings(c.Do("MGET", query...))
	if err != nil {
		panic(err)
	}
	return res
}

// TTLの設定
func Expire(key string, ttl int, c redis.Conn) {
	c.Do("EXPIRE", key, ttl)
}

func main() {
	c := Connection()
	defer c.Close()

	res_set := Set("sample-key", "sample-value", c)
	fmt.Println(res_set)

	res_get := Get("sample-key", c)
	fmt.Println(res_get)

	datas := []Data{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
	}
	Mset(datas, c)

	keys := []string{"key1", "key2"}
	res_mget := Mget(keys, c)
	fmt.Println(res_mget)

	Expire("key1", 10, c)
}
