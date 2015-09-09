package redis

import (
	redigo "github.com/garyburd/redigo/redis"
)

var Pool *redigo.Pool

func Connect() {
	Pool = redigo.NewPool(
		func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", "redis")
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
