package main

import (
	"time"

	"github.com/Gerard-007/bootdotdev/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json: "id"id"`
	Email     string    `json: "email"`
	Password  string    `json: "password"`
	Username  string    `json: "username"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	APIKey    string    `json: "api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json: "id"`
	Name      string    `json: "name"`
	Url       string    `json: "url"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	UserID    uuid.UUID `json: "user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt.Time,
		UpdatedAt: dbFeed.UpdatedAt.Time,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	FeedID    uuid.UUID `json: "feed_id"`
	UserID    uuid.UUID `json: "user_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt.Time,
		UpdatedAt: dbFeedFollow.UpdatedAt.Time,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
	}
}

func databaseFeedFollowToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID     `json: "id"`
	Title       string        `json: "title"`
	Description *string       `json: "description"`
	Url         string        `json: "url"`
	Published   time.Time     `json: "published"`
	FeedID      uuid.NullUUID `json: "feed_id"`
	CreatedAt   time.Time     `json: "created_at"`
	UpdatedAt   time.Time     `json: "updated_at"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description != "" {
		description = &dbPost.Description
	}
	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Description: description,
		Url:         dbPost.Url,
		Published:   dbPost.Published,
		FeedID:      dbPost.FeedID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}