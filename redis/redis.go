package redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/sirsean/textback/config"
)

var Pool *redigo.Pool

func Connect() {
	Pool = redigo.NewPool(
		func() (redigo.Conn, error) {
			host := config.Get().Redis.Host
			port := config.Get().Redis.Port
			c, err := redigo.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		3,
	)
}

func Close() {
	Pool.Close()
}
