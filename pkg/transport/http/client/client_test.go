package client

import (
	"fmt"
	"testing"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

type Case struct {
	label          string
	input          string
	expectedOutput *domain.Message
}

func TestParseData(t *testing.T) {
	testCases := []Case{
		{
			label: "testing tweets parsing",
			input: ``,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Favorites: toIntPointer(1),
					Retweets:  toIntPointer(1),
					Timestamp: 23867283749,
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {
		fmt.Println(tc.label)
	}
}
