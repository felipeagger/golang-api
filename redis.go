package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.15.99:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		println("Error No Redis")
		return nil, err
	}

	return client, nil
}

func MarshalBinary(user User) ([]byte, error) {
	return json.Marshal(user)
}

func UnmarshalBinary(user string) User {
	var _user User

	if err := json.Unmarshal([]byte(user), &_user); err != nil {
		fmt.Println("Error: ", err)
	}

	return _user
}

func RedisSet(key string, value User) {
	client, _ := NewRedisClient()

	json, error := MarshalBinary(value)

	if error != nil {
		panic(error)
	}

	err := client.Set(key, string(json), 0).Err()

	if err != nil {
		panic(err)
	}
}

func RedisGet(key string) User {
	client, _ := NewRedisClient()

	value, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(key, UnmarshalBinary(value))
	return UnmarshalBinary(value)
}
