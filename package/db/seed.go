package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/thejus-r/social/package/store"
)

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error seeding users", err.Error())
		}
	}

	posts := generatePosts(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error seeding posts", err.Error())
		}
	}

	comments := generateComments(users, posts)
	for _, post := range comments {
		if err := store.Comments.Create(ctx, post); err != nil {
			log.Println("Error seeding comments", err.Error())
		}
	}
}

func generateComments(users []*store.User, posts []*store.Post) []*store.Comment {
	var comments []*store.Comment
	noOfUsers := len(users) - 1

	for i := 0; i < len(posts)-1; i++ {
		noOfComments := getRandomNumber(0, 3)
		for j := 0; j < noOfComments; j++ {
			userID := getRandomNumber(1, noOfUsers)
			temp := &store.Comment{
				PostID: posts[i].ID,
				UserID: users[userID].ID,
			}
			comments = append(comments, temp)
		}
	}
	return comments
}

func generateUsers(num int) []*store.User {

	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: fmt.Sprintf("%s%d", usernames[i], i),
			Email:    fmt.Sprintf("%s%d@mail.com", strings.ToLower(usernames[i]), i),
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	noOfUsers := len(users)

	for i := 0; i < num; i++ {
		posts[i] = &store.Post{
			Title:   titles[i],
			Content: content,
			Tags:    []string{tags[getRandomNumber(0, 9)], tags[getRandomNumber(0, 9)]},
			UserId:  users[getRandomNumber(0, noOfUsers)].ID,
		}
	}
	return posts
}

func getRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
