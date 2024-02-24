package main

import (
	"time"

	"github.com/davifrjose/BOILERPLATE/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey: user.ApiKey,
	}
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url :     feed.Url, 
		UserID :    feed.UserID,
	}
}

func databasesFeedToFeed(dbFeebs []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbFeed := range dbFeebs {
		feeds = append(feeds, databaseFeedToFeed(dbFeed) )
	}
	return feeds
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowsToFeedFollows(feedFollow database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		FeedID: feedFollow.FeedID,
		UserID :    feedFollow.UserID,
	}
}

func databasesFeedFollowToFeedFollow(dbFeedFollows []database.FeedFollow) []FeedFollows {
	feedFollows := []FeedFollows{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowsToFeedFollows(dbFeedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	var description *string
	if post.Description.Valid {
		description = &post.Description.String
	}
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: description,
		PublishedAt: &post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}
