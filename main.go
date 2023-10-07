package chatroom

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func main() {
	//redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CHAT_ROOM_REDIS"),
		Password: "",
		DB:       1,
	})

	//redis.NewDialer()

	println("redis", rdb)
}
