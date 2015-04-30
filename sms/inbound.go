package sms

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/sirsean/textback/model"
	"github.com/sirsean/textback/redis"
	"github.com/sirsean/textback/tz"
	"log"
	"net/http"
	"time"
)

var postDecoder = schema.NewDecoder()

func Inbound(w http.ResponseWriter, r *http.Request) {
	type Form struct {
		From      string
		To        string
		Body      string
		FromState string
	}

	r.ParseForm()
	form := new(Form)
	postDecoder.Decode(form, r.PostForm)

	log.Printf("received message")
	log.Printf("%v", form)

	loc := tz.ByState(form.FromState)
	log.Printf("%v", loc)

	command := model.NewCommandParseWithPhone(form.Body, form.From, loc)

	log.Printf("%v", command)

	commandJson, _ := json.Marshal(command)

	c := redis.Pool.Get()
	defer c.Close()

	c.Do("RPUSH", command.WhenKey(), commandJson)

	w.Header().Set("Content-Type", "text/xml")
	w.Write([]byte(fmt.Sprintf(`
		<Response>
			<Message>Okay! I'll text you at %v</Message>
		</Response>
	`, command.When.Format(time.RFC822))))
}
