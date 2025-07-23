package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/smnthjm08/go-social/internal/store"
)

var usernames = []string{
	"user1", "user2", "user3", "user4", "user5",
	"user6", "user7", "user8", "user9", "user10",
}

var titles = []string{
	"Hello World",
	"Getting Started with Go",
	"Understanding Interfaces",
	"Concurrency Basics",
	"REST API in Go",
	"Error Handling Tips",
	"Working with Databases",
	"Testing in Go",
	"Structs and Methods",
	"Go Routines Explained",
	"Channels in Go",
	"Dependency Management",
	"JSON in Go",
	"File I/O Basics",
	"Logging Best Practices",
	"Go Modules Intro",
	"HTTP Servers in Go",
	"Unit Testing Guide",
	"Pointers in Go",
	"Deploying Go Apps",
}

var contents = []string{
	"Welcome to my first blog post! Here I share my journey with Go.",
	"This post will help you get started with Go programming language.",
	"Let's dive into Go interfaces and how they make code flexible.",
	"Concurrency in Go is powerful and easy to use. Learn the basics here.",
	"Building a REST API in Go is straightforward. Follow these steps.",
	"Error handling in Go can be tricky. Here are some useful tips.",
	"Learn how to connect and work with databases in Go.",
	"Testing is crucial. This post covers unit testing in Go.",
	"Structs and methods are core to Go. Understand them in this post.",
	"Go routines make concurrency simple. Here's how they work.",
	"Channels are used for communication between Go routines.",
	"Manage your Go dependencies efficiently with these tips.",
	"Working with JSON in Go is easy. See how to encode and decode.",
	"File I/O is essential. Learn how to read and write files in Go.",
	"Logging helps debug and monitor apps. Best practices inside.",
	"Go modules simplify dependency management. Get started here.",
	"Set up an HTTP server in Go with just a few lines of code.",
	"Unit testing ensures code quality. A quick guide for Go.",
	"Pointers in Go can be confusing. This post makes them clear.",
	"Deploy your Go applications easily with these strategies.",
}

var tags = []string{
	"go", "programming", "tutorial", "web", "api", "database", "testing", "concurrency", "json", "deployment",
}

var comments = []string{
	"Great post!",
	"Thanks for sharing.",
	"Very informative.",
	"Awesome, keep it up!",
	"Nice explanation.",
}

func Seed(store store.Storage) error {
	ctx := context.Background()

	users := generateUsers(20)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return err
		}
	}

	posts := generatePosts(100, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return err
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return err
		}
	}

	log.Println("Seeding complete...")
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123456",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		var userID int64
		// Sscanf scans the argument string, storing successive space-separated values
		// into successive arguments as determined by the format. It returns the number of items
		// successfully parsed. Newlines in the input must match newlines in the format.
		fmt.Sscanf(user.ID, "%d", &userID)
		posts[i] = &store.Post{
			UserId:  userID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		var userID int64
		fmt.Sscanf(user.ID, "%d", &userID)
		cms[i] = &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  userID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
