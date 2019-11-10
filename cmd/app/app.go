package main

/*
 *
 */
import (
	"encoding/json"
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	Msg struct {
		Data    string `json:"data"`
		Address string `json:"address"`
	}
)

var (
	redisClient *redis.Client
)

// устанавливает дискорд пользователя
func discordSet(w http.ResponseWriter, r *http.Request) {
	type reqFormat struct {
		UserId  string `json:"userId"`
		Address string `json:"address"`
	}
	var (
		jsonBody []byte
		err      error
		req      reqFormat
	)

	defer r.Body.Close()
	jsonBody, err = ioutil.ReadAll(r.Body)

	err = json.Unmarshal(jsonBody, &req)
	if err != nil {
		fmt.Println(err, string(jsonBody))
		fmt.Fprint(w, `wrong req format`)
		return
	}

	redisClient.Set(req.Address, req.UserId, time.Hour*10000)
	return
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	http.Handle("/events/", s)
	http.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		jsonBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Fprint(writer, `req body not json`)
			fmt.Println(err)
			return
		}

		var msg Msg
		err = json.Unmarshal(jsonBody, &msg)
		if err != nil {
			fmt.Fprint(writer, `req body wrong type`)
			fmt.Println(err)
			return
		}

		s.SendMessage("/events/"+msg.Address, sse.SimpleMessage(msg.Data))
	})
	http.HandleFunc("/discordSet", discordSet)
	http.HandleFunc("/getTokens", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(`https://etherscan.io/token/generic-tokenholder-inventory?` + r.URL.Query().Encode())
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, `api query err`)
			return
		}

		jsonBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("body read err", err)
			fmt.Fprint(w, `something wrong`)
			return
		}

		fmt.Fprint(w, string(jsonBody))
	})

	http.ListenAndServe(":3000", nil)
}
