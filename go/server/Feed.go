package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	feed "github.com/bartick/learning-gRPC/go/proto-go"
	_ "github.com/mattn/go-sqlite3"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("sqlite3", "./../../database.sqlite")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS feeds (id INTEGER PRIMARY KEY, title TEXT, content TEXT, author TEXT, createdAt TEXT, updatedAt TEXT)")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS sqlite_sequence (name TEXT, seq INTEGER)")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	fmt.Println("Connected to database")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nClosing Database...")
		db.Close()
		os.Exit(0)
	}()

}

type Server struct {
	feed.FeedServiceServer
}

func (s *Server) GetFeed(ctx context.Context, in *feed.FeedRequest) (*feed.FeedReply, error) {

	if in.FeedId == "" {
		return nil, fmt.Errorf("FeedId is required in the request")
	}

	feedId, err := strconv.Atoi(in.FeedId)
	if err != nil {
		return nil, err
	}

	dataRow := db.QueryRow("SELECT * FROM feeds WHERE id = ?", feedId)

	var (
		id        int32
		title     string
		content   string
		author    string
		createdAt string
		updatedAt string
	)

	err = dataRow.Scan(&id, &title, &content, &author, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	return &feed.FeedReply{
		FeedId:          id,
		FeedTitle:       title,
		FeedContent:     content,
		FeedPublishTime: createdAt,
		FeedAuthor:      author,
	}, nil
}

func (s *Server) GetAllFeeds(ctx context.Context, in *emptypb.Empty) (*feed.AllFeeds, error) {

	rows, err := db.Query("SELECT * FROM feeds")
	if err != nil {
		return nil, err
	}

	var (
		id        int32
		title     string
		content   string
		author    string
		createdAt string
		updatedAt string
	)
	feedToSend := make([]*feed.FeedReply, 0)
	for rows.Next() {
		err = rows.Scan(&id, &title, &content, &author, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		feedToSend = append(feedToSend, &feed.FeedReply{
			FeedId:          id,
			FeedTitle:       title,
			FeedContent:     content,
			FeedPublishTime: createdAt,
			FeedAuthor:      author,
		})
	}

	return &feed.AllFeeds{
		Feeds: feedToSend,
	}, nil
}

func (s *Server) PostFeed(ctx context.Context, in *feed.FeedPost) (*feed.FeedReply, error) {

	if in.FeedTitle == "" {
		return nil, fmt.Errorf("FeedTitle is required in the request")
	}
	if in.FeedContent == "" {
		return nil, fmt.Errorf("FeedContent is required in the request")
	}
	if in.FeedAuthor == "" {
		return nil, fmt.Errorf("FeedAuthor is required in the request")
	}

	row := db.QueryRow("SELECT seq FROM sqlite_sequence WHERE name = 'feeds'")
	var feedId int32
	err = row.Scan(&feedId)
	if err != nil {
		return nil, err
	}

	_, err := db.Exec("INSERT INTO feeds (id, title, content, author, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)", feedId, in.FeedTitle, in.FeedContent, in.FeedAuthor, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		return nil, err
	}

	_, _ = db.Exec("UPDATE sqlite_sequence SET seq = ? WHERE name = 'feeds'", feedId+1)

	return &feed.FeedReply{
		FeedId:          feedId,
		FeedTitle:       in.FeedTitle,
		FeedContent:     in.FeedContent,
		FeedPublishTime: time.Now().Format("2006-01-02 15:04:05"),
		FeedAuthor:      in.FeedAuthor,
	}, nil
}

func (s *Server) DeleteFeed(ctx context.Context, in *feed.FeedRequest) (*feed.FeedSuccess, error) {

	if in.FeedId == "" {
		return nil, fmt.Errorf("FeedId is required in the request")
	}

	feedId, err := strconv.Atoi(in.FeedId)
	if err != nil {
		return nil, err
	}

	row, err := db.Exec("DELETE FROM feeds WHERE id = ?", feedId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no feed found with id: %d to delete", feedId)
	}

	return &feed.FeedSuccess{
		Message: "success",
	}, nil
}
