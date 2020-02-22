package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":6379",
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

func MarshalBinary(users []User) ([]byte, error) {
	return json.Marshal(users)
}

func UnmarshalBinary(users string) []User {
	var _user []User

	if err := json.Unmarshal([]byte(users), &_user); err != nil {
		fmt.Println("Error: ", err)
	}

	return _user
}

func RedisSet(key string, value []User) {
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

func RedisGet(key string) []User {
	client, _ := NewRedisClient()

	value, err := client.Get(key).Result()
	if err != nil {
		//panic(err)
		return nil
	} else {
		fmt.Println(key, UnmarshalBinary(value))
		return UnmarshalBinary(value)
	}

}
