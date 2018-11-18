package utils

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

func JoinUrl(base string, new ...interface{}) string {
	u, err := url.Parse(base)

	if err != nil {
		log.Fatal(err)
	}

	for _, token := range new {
		u.Path = path.Join(u.Path, fmt.Sprint(token))
	}

	s := u.String()
	log.Printf(s)

	return s
}
