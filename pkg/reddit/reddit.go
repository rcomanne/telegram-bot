package reddit

import (
	"context"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"math/rand"
)

var ctx = context.Background()

func GetPostFromSubreddit(subreddit string) (*reddit.Post, error) {
	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "week",
	})
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.URL)
		fmt.Println(post.Body)
	}

	selected := posts[rand.Intn(len(posts))]

	return selected, err
}
