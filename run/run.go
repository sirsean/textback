package run

import (
	"encoding/json"
	"github.com/sirsean/textback/model"
	"github.com/sirsean/textback/redis"
	"github.com/sirsean/textback/sms"
	"log"
	"time"
)

func Start() {
	ticker := time.Tick(20 * time.Second)
	tick()
	for {
		select {
		case <-ticker:
			tick()
		}
	}
}

func tick() {
	c := redis.Pool.Get()
	defer c.Close()

	key := model.FormatKey(time.Now())

	obj, _ := c.Do("LPOP", key)
	for obj != nil {
		var command model.Command
		json.Unmarshal(obj.([]byte), &command)
		err := sms.Send(command.Phone, command.Message)
		if err != nil {
			log.Printf("failed to send command %v %v", command, err)
		}
		obj, _ = c.Do("LPOP", key)
	}
}
