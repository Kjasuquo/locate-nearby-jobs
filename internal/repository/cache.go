package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kjasuquo/jobslocation/internal/model"
	"log"
	"os"
	"time"
)

type Redis struct {
	cache *redis.Client
}

func NewRedis() RedisRepo {

	var opts *redis.Options

	if os.Getenv("LOCAL") == "true" {
		redisAddr := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))
		opts = &redis.Options{
			Addr:     redisAddr,
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	} else {
		HerokuOpts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}

		opts = HerokuOpts

	}

	rdb := redis.NewClient(opts)

	return &Redis{cache: rdb}
}

func (r *Redis) Get(ctx context.Context, title string) ([]model.Jobs, error) {
	var job []model.Jobs
	val, err := r.cache.Get(ctx, title).Result()
	if err != nil {
		log.Printf("error getting cache from redis: %v\n", err)
		return nil, err
	}

	//cache hit
	fmt.Println("cache hit")
	err = json.Unmarshal(bytes.NewBufferString(val).Bytes(), &job)
	if err != nil {
		log.Printf("cannot unmarshall cache: %v\n", err)
		return nil, err
	}

	return job, nil
}

func (r *Redis) Set(ctx context.Context, title string, job []model.Jobs) error {
	fmt.Println("cache miss")
	byteJob, err := json.Marshal(job)
	if err != nil {
		log.Printf("cannot marshall job: %v\n", err)
		return err
	}

	//set redis with the value
	err = r.cache.Set(ctx, title, bytes.NewBuffer(byteJob).Bytes(), time.Second*15).Err()
	if err != nil {
		log.Printf("cannot set cache: %v\n", err)
		return err
	}

	return nil
}
