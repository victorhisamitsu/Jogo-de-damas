package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SearchRedis[T any](ctx context.Context, client *redis.Client, key string, chave string) (*T, error) {
	// Utilize um hash do Redis para salvar o UUID e a data de acesso do usuário
	concatenado := chave + ":" + key + ":"
	resultado, err := client.Get(ctx, concatenado).Result()
	if err != nil {
		fmt.Println("deu bem ruim")
		return nil, err
	}
	if resultado == "" {
		return nil, errors.New("erro na requisição")
	}
	var respostas T
	err = json.Unmarshal([]byte(resultado), &respostas)
	if err != nil {
		return nil, err
	}
	return &respostas, nil
}
