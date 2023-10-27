package storage

import (
	"context"
	"github.com/luqus/s/types"
	"github.com/redis/go-redis/v9"
)

type PubSub interface {
	Publish(ctx context.Context, data *types.PublishLocationInput) error
	Subscribe(ctx context.Context, channelID string) (*types.Location, error)
}

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPubSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{
		client: client,
	}
}

func RedisInit(redisConnString,password string, db int) *redis.Client{
	client:=redis.NewClient(&redis.Options{
		Addr: redisConnString,
		Password: password,
		DB: db,
	})
	
	return client
}


func (r *RedisPubSub) Publish(ctx context.Context, data *types.PublishLocationInput) error {
	err := r.client.Publish(ctx,data.ChannelID, data.Location)
	if err != nil {
		return err.Err()
	}
	
	return nil
}


func (r*RedisPubSub) Subscribe(ctx context.Context,channelID string) (*types.Location,error) {
	location,err := r.client.Subscribe(ctx, channelID).Receive(ctx)
	if err != nil {
		return nil,err
	}
	
	return location.(*types.Location), nil
}