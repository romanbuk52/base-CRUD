package web

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// store, err := NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
// if err != nil {
// 	panic(err)
// }
// defer store.Close()

var ctx = context.Background()

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.13.15:6379",
		Password: "Tratata101", // no password set
		DB:       0,            // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

// func MyHandler(w http.ResponseWriter, r *http.Request) {
// 	// Get a session. Get() always returns a session, even if empty.
// 	session, err := store.Get(r, "session-name")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set some session values.
// 	session.Values["foo"] = "bar"
// 	session.Values[42] = 43
// 	// Save it before we write to the response/return from the handler.
// 	err = session.Save(r, w)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
