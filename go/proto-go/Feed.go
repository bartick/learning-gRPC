package feed

import (
	"context"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	FeedServiceServer
}

func (s *Server) GetFeed(ctx context.Context, in *FeedRequest) (*FeedReply, error) {
	return &FeedReply{
		FeedId:          1,
		FeedTitle:       "A Test Feed",
		FeedContent:     "This is the feed content that you don't need",
		FeedPublishTime: time.Now().GoString(),
		FeedAuthor:      "bartick",
	}, nil
}

func (s *Server) GetAllFeed(ctx context.Context, in *emptypb.Empty) (*AllFeeds, error) {
	return &AllFeeds{
		Feeds: []*FeedReply{
			{
				FeedId:          1,
				FeedTitle:       "A Test Feed",
				FeedContent:     "This is the feed content that you don't need",
				FeedPublishTime: time.Now().GoString(),
				FeedAuthor:      "bartick",
			},
			{
				FeedId:          2,
				FeedTitle:       "A Test Feed 2",
				FeedContent:     "This is the feed content that you don't need",
				FeedPublishTime: time.Now().GoString(),
				FeedAuthor:      "bartick",
			},
		},
	}, nil
}

func (s *Server) PostFeed(ctx context.Context, in *FeedPost) (*FeedReply, error) {
	return &FeedReply{
		FeedId:          1,
		FeedTitle:       "A Test Feed",
		FeedContent:     "This is the feed content that you don't need",
		FeedPublishTime: time.Now().GoString(),
		FeedAuthor:      "bartick",
	}, nil
}

func (s *Server) DeleteFeed(ctx context.Context, in *FeedRequest) (*FeedSuccess, error) {
	return &FeedSuccess{
		Message: "success",
	}, nil
}
