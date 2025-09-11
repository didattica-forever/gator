package main

import (
	"context"
	"fmt"

	"github.com/didattica-forever/gator/internal/database"
)

func handlerListFeeds(s *state, cmd command) error {
	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	printFeedDetail(s, feeds)


	return nil
}

func printFeedDetail(s *state, feeds []database.Feed) {
	for _, feed := range(feeds) {
		fmt.Printf("* ID:            %s\n", feed.ID)
		fmt.Printf("* Created:       %v\n", feed.CreatedAt)
		fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)

		user, err := s.db.GetUserById(context.Background(), feed.UserID)

		username := "Unknown"
		if err == nil {
			 username = user.Name
		}

		fmt.Printf("* UserID:        %s\n", feed.UserID)		
		fmt.Printf("* UserName:      %s\n", username)		
	}

}
