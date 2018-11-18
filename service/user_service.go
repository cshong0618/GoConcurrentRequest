package service

import (
	"encoding/json"
	. "gojsonplaceholder/model"
	"gojsonplaceholder/utils"
	"log"
	"net/http"
	"strconv"
)

type UserData struct {
	User User 	`json:"user"`
	Todos []Todo `json:"todos"`
	Posts []Post `json:"posts"`
}

var baseUrl = "https://jsonplaceholder.typicode.com/"


func getUser(userId int, ch chan User) {
	url := utils.JoinUrl(baseUrl, "users", userId)
	response, err := http.Get(url)

	if err != nil {
		log.Print(err)
		ch <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	err = json.NewDecoder(response.Body).Decode(&user)

	if err != nil {
		log.Print(err)
		user = User{}
	}

	log.Printf("got user")
	ch <- user
}

func getTodos(userId int, ch chan []Todo) {
	url := utils.JoinUrl(baseUrl, "todos") +
		"/?userId=" + strconv.Itoa(userId)
	response, err := http.Get(url)

	if err != nil {
		log.Print(err)
		ch <- []Todo{}
		return
	}
	defer response.Body.Close()

	var todos []Todo
	err = json.NewDecoder(response.Body).Decode(&todos)

	if err != nil {
		log.Print(err)
	}

	log.Printf("got todo")
	ch <- todos
}

func getComments(userId int, ch chan []Post) {
	url := utils.JoinUrl(baseUrl, "posts") +
		"/?userId=" + strconv.Itoa(userId)
	response, err := http.Get(url)

	if err != nil {
		log.Print(err)
		ch <- []Post{}
		return
	}
	defer response.Body.Close()

	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)

	if err != nil {
		log.Print(err)
	}

	log.Printf("got posts")
	ch <- posts
}

func GetUserData(userId int) UserData{
	var userData UserData

	chUser := make(chan User)
	chTodos := make(chan []Todo)
	chPosts := make(chan []Post)

	go getUser(userId, chUser)
	go getTodos(userId, chTodos)
	go getComments(userId, chPosts)

	userData.User = <- chUser
	userData.Todos = <- chTodos
	userData.Posts = <- chPosts

	return userData
}