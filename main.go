package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	defer conn.Close()
	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	checkErr(err)
	title, err := redis.String(conn.Do("HGET", "album:2", "title"))
	checkErr(err)
	fmt.Println(title)

}
