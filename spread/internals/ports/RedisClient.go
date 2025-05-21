package ports

type RedisClient interface {
	MakeRedisClient() interface{}
}
