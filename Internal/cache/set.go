package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Cache *MemoryRedis

type MemoryRedis struct {
	Store *redis.Client
}

func NewCache(client *redis.Client) *MemoryRedis {
	Cache = &MemoryRedis{Store: client}
	return Cache
}

func (m *MemoryRedis) Insert(ctx context.Context, chave string, valor any, typeData string) {
	SaveDataRedis[any](ctx, m.Store, chave, valor, typeData)
}

func SaveDataRedis[T any](ctx context.Context, client *redis.Client, key string, value T, chave string) error {
	// Utilize um hash do Redis para salvar o UUID e a data de acesso do usuário
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	concatenado := chave + ":" + key
	fmt.Println(concatenado)
	err = client.Set(ctx, concatenado, string(valueByte), 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
