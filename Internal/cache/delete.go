package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func DeleteRedis(ctx context.Context, client *redis.Client, key string, chave string) error {
	// Utilize um hash do Redis para salvar o UUID e a data de acesso do usuário
	concatenado := chave + ":" + key
	resposta, err := client.Del(ctx, concatenado).Result()
	if err != nil {
		return err
	}
	if resposta == 0 {
		return errors.New("nenhum dado apagado")
	}
	fmt.Printf("int: %v\n", resposta)
	return nil
}
