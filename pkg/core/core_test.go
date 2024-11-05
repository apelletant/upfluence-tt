//nolint:dupl
package core_test

import (
	"maps"
	"testing"

	"github.com/apelletant/upfluence-tt/pkg/core"
	"github.com/apelletant/upfluence-tt/pkg/domain"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/client"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/clientmock"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/servermock"
)

type Case struct { //nolint:govet
	label           string
	msgs            []*domain.Message
	expectedOuput   map[string]int
	wantedDimension string
}

func TestRunQuery(t *testing.T) {
	server := servermock.New()

	testCases := []Case{
		{
			label:           "Testing likes",
			wantedDimension: "likes",
			msgs: []*domain.Message{
				{
					Data: &domain.MsgData{
						Likes:     client.ToIntPointer(20),
						Comments:  nil,
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1468897913,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     client.ToIntPointer(10),
						Comments:  nil,
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1264454550,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     client.ToIntPointer(30),
						Comments:  nil,
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1733010686,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     client.ToIntPointer(0),
						Comments:  nil,
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1274798714,
					},
					Err: nil,
				},
			},
			expectedOuput: map[string]int{
				"minimum_timestamp": 1264454550,
				"maximum_timestamp": 1733010686,
				"avg_likes":         15,
				"total_posts":       4,
			},
		},
		{
			label:           "Testing comments",
			wantedDimension: "comments",
			msgs: []*domain.Message{
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Comments:  client.ToIntPointer(20),
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1468897913,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Comments:  client.ToIntPointer(10),
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1264454550,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Comments:  client.ToIntPointer(30),
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1733010686,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Comments:  client.ToIntPointer(0),
						Favorites: nil,
						Retweets:  nil,
						Timestamp: 1274798714,
					},
					Err: nil,
				},
			},
			expectedOuput: map[string]int{
				"minimum_timestamp": 1264454550,
				"maximum_timestamp": 1733010686,
				"avg_comments":      15,
				"total_posts":       4,
			},
		},
		{
			label:           "Testing retweets",
			wantedDimension: "retweets",
			msgs: []*domain.Message{
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Favorites: nil,
						Comments:  nil,
						Retweets:  client.ToIntPointer(20),
						Timestamp: 1468897913,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Retweets:  client.ToIntPointer(10),
						Favorites: nil,
						Comments:  nil,
						Timestamp: 1264454550,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Retweets:  client.ToIntPointer(30),
						Favorites: nil,
						Comments:  nil,
						Timestamp: 1733010686,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Retweets:  client.ToIntPointer(0),
						Favorites: nil,
						Comments:  nil,
						Timestamp: 1274798714,
					},
					Err: nil,
				},
			},
			expectedOuput: map[string]int{
				"minimum_timestamp": 1264454550,
				"maximum_timestamp": 1733010686,
				"avg_retweets":      15,
				"total_posts":       4,
			},
		},
		{
			label:           "Testing favorites",
			wantedDimension: "favorites",
			msgs: []*domain.Message{
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Favorites: client.ToIntPointer(20),
						Comments:  nil,
						Retweets:  nil,
						Timestamp: 1468897913,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Favorites: client.ToIntPointer(10),
						Comments:  nil,
						Retweets:  nil,
						Timestamp: 1264454550,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Favorites: client.ToIntPointer(30),
						Comments:  nil,
						Retweets:  nil,
						Timestamp: 1733010686,
					},
					Err: nil,
				},
				{
					Data: &domain.MsgData{
						Likes:     nil,
						Favorites: client.ToIntPointer(0),
						Comments:  nil,
						Retweets:  nil,
						Timestamp: 1274798714,
					},
					Err: nil,
				},
			},
			expectedOuput: map[string]int{
				"minimum_timestamp": 1264454550,
				"maximum_timestamp": 1733010686,
				"avg_favorites":     15,
				"total_posts":       4,
			},
		},
	}

	for _, tc := range testCases {
		client := clientmock.New(tc.msgs)

		appDeps := &domain.Dependencies{
			Client: client,
			Server: server,
		}

		app, _ := core.NewApp(appDeps)

		// hard coded duration 'cause not tested and no incidence on the test resutl
		res, err := app.RunQuery(tc.wantedDimension, "10s")
		if err != nil {
			t.Fatal("error while running query", err)
		}

		if ok := maps.Equal(tc.expectedOuput, res); !ok {
			t.Fail()
		}
	}
}
