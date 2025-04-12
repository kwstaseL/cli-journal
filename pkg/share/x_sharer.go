package share

import (
	"fmt"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

const TweetCharLimiter = 280

type XSharer struct {
	// some creds?
}

type Tweet struct {
	Content string
}

func (t *XSharer) Share(note model.Note) error {
	tweet := formatNoteForTweet(note)
	return sendTweet(tweet)
}

func formatNoteForTweet(note model.Note) Tweet {
	content := fmt.Sprintf("%s\n%s\nCategory: %s\nTags: %s", 
		note.Header, note.Body, note.Category, note.Tags)
	
	if len(content) > TweetCharLimiter {
		content = content[:277] + "..."
	}

	return Tweet{Content: content}
}

func sendTweet(tweet Tweet) error {
	// Mock
	fmt.Println("Sending tweet:")
	fmt.Println(tweet.Content)
	return nil
}