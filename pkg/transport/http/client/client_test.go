package client_test

import (
	"fmt"
	"testing"

	"github.com/apelletant/upfluence-tt/pkg/domain"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/client"
)

type Case struct { //nolint:govet
	label          string
	input          string
	expectedOutput *domain.Message
}

func TestParseData(t *testing.T) {
	testCases := []Case{
		{
			label: "testing tweets parsing",
			input: `{
				"tweet": {
					"id": 996016728,
					"content": "@UbisoftSupport LET ME PLAY SIEGE!!!! \u0026gt;:C",
					"retweets": 34,
					"favorites": 102,
					"timestamp": 1722386416,
					"post_id": "1818446548289675280",
					"is_retweet": false,
					"comments": 5
				}
    	}`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Favorites: client.ToIntPointer(102),
					Retweets:  client.ToIntPointer(34),
					Timestamp: 1722386416,
					Comments:  client.ToIntPointer(5),
				},
				Err: nil,
			},
		},
		{
			label: "testing youtube parsing",
			input: `{
				"youtube_video": {
					"id": 220984035,
					"name": "Pineapple \u0026 Kimchi Freid Rice topped with Teriyaki Steak",
					"description": "Pineapple \u0026 Kimchi Freid Rice topped with Teriyaki Steak ðŸ”¥ðŸðŸ¥© \n\nThis recipe right here is easy to make, packed with flavor and can be done in under 30 minutes! Do yourself a favor and give this one a try for your next dinner or meal prep!ðŸ’ªðŸ»\n\nKnife from @excalibladesknives ðŸ”ª \n12 in Wok from @hexclad \n\nFULL RECIPE â¬‡ï¸\n\nINGREDIENTS \n- ~2 lbs steak of choice cut into bite size cubes \n(I used top sirloin)\n- Extra virgin olive oil\n- @lanesbbq Q-Nami Seasoning\n- @primalkitchenfoods Teriyaki Sauce\n\nPineapple \u0026 Kimchi Fried Rice:\n- @fourthandheart garlic ghee \n- 1 small white onion diced\n- ~3/4 cup bottom of green onions sliced\n- 2 eggs \n- ~1.5 cup pineapple diced \n- ~1 cup @clevelandkitchen kimchi \n- ~4 cups cooked jasmine rice\n- ~1 tbsp Toasted sesame oil \n- ~2-3 tbsp Coconut aminos\n- ~1/3 cup @trybachans Japanese BBQ Sauce\n- ~1/4 cup @primalkitchenfoods Teriyaki Sauce \n- ~1-2 tbsp Sriracha \n- Seasonings: Salt, pepper and garlic powder to taste. \n\nToppings:\n- Sesame seeds  \n- Sliced green onions\n\nNote(s): All measurements are estimated, adjust ingredients to your liking. \n\nDIRECTIONS \n1.  Add cubed steak to a large bowl, drizzle in olive oil and season generously with the Lane's Q-Nami seasoning. Mix till well combined and set aside.\n2. Get a large skillet or wok going over medium heat. Add in 2-3 tbsp of ghee along with both onions. Let saute for a few minutes then add in the eggs. Scramble and let this cook then mix them in. Next mix in the pineapple and kimchi then add in the rice along with everything else. Mix till well combined and let fry for a few minutes. Season to taste. \n3. Add 1-2 tbsp of ghee to a skillet over medium-high heat then sear the steak for 2-3 minutes per side. Once almost done, lower the heat and add in your preferred amount of teriyaki sauce. Mix and let simmer till sauce thickens. \n4. Plate it up \u0026 ENJOY! ðŸ™ŒðŸ»\nâ€¢\nâ€¢\nâ€¢\nâ€¢\n#explore #cooking #bhfyp #food #easyrecipe #simple #foodporn #healthy #tasty #shorts #healthyrecipes #calwillcookit #diet #healthylifestyle #mealprep #gains #weightloss #steak #fitfood #highprotein",
					"link": "https://www.youtube.com/watch?v=q-mdF1Uaqkk",
					"views": 5289,
					"comments": 0,
					"likes": 175,
					"dislikes": 0,
					"timestamp": 1728510897,
					"post_id": "q-mdF1Uaqkk"
		    	}
		  }`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Timestamp: 1728510897,
					Comments:  client.ToIntPointer(0),
					Likes:     client.ToIntPointer(175),
				},
				Err: nil,
			},
		},
		{
			label: "testing pin parsing",
			input: `{
				"pin": {
					"id": 103806200,
					"title": "Easy Pineapple Whip Ice Cream Recipe",
					"description": "Pineapple Whip Ice Cream, anyone? Only 3 ingredients for this creamy dream:\n\nðŸ¥› 2 cups milk of choice (we used lactose-free milk)\nðŸ® 2 tsp vanilla sugar-free pudding mix\nðŸŒŸ 2 scoops SparkÂ® Pineapple Whip\n\nBlend, freeze and enjoy a guilt-free treat! Perfect for summer days! â˜€ï¸ðŸâœ¨",
					"links": "https://www.advocare.com/advocare-spark/A7932.html",
					"likes": 234,
					"comments": 198,
					"saves": 0,
					"repins": 1,
					"timestamp": 1723152324,
					"post_id": "81135230780683268"
				}
		  }`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Timestamp: 1723152324,
					Comments:  client.ToIntPointer(198),
					Likes:     client.ToIntPointer(234),
				},
				Err: nil,
			},
		},
		{
			label: "testing article parsing",
			input: `{
				"article": {
					"id": 396210547,
					"title": "Introducing Your Irish American Heritage To Friends",
					"description": "\u003ca href=\"https://www.irishamericanmom.com/introducing-your-irish-american-heritage-to-friends/\" title=\"Introducing Your Irish American Heritage To Friends\" rel=\"nofollow\"\u003e\u003cimg width=\"300\" height=\"300\" src=\"https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-300x300.jpg\" class=\"webfeedsFeaturedVisual wp-post-image\" alt=\"American and Irish flags with a text banner\" style=\"display: block; margin: auto; margin-bottom: 30px;max-width: 100%;\" link_thumbnail=\"1\" decoding=\"async\" fetchpriority=\"high\" srcset=\"https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-300x300.jpg 300w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-1024x1024.jpg 1024w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-90x90.jpg 90w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-768x768.jpg 768w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-500x500.jpg 500w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-360x360.jpg 360w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-720x720.jpg 720w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-180x180.jpg 180w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-96x96.jpg 96w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends-150x150.jpg 150w, https://www.irishamericanmom.com/wp-content/uploads/2024/09/Irish-American-Heritage-Sharing-Your-Irish-Roots-with-Friends.jpg 1200w\" sizes=\"(max-width: 300px) 100vw, 300px\" data-pin-description=\"Explore ways to share information about your Irish heritage and roots with friends and family.\" data-pin-title=\"Irish American Heritage Sharing Your Irish Roots With Friends\" data-pin-url=\"https://www.irishamericanmom.com/introducing-your-irish-american-heritage-to-friends/?tp_image_id=54654\" /\u003e\u003c/a\u003e\u003cp\u003eLearning about the depth of your Irish heritage and family ancestry is a big deal for most people. Irish Americans often feel a deep link to the land of their ancestors. Studying your family tree and learning about your family members from the Emerald Isle starts a journey of discovery to find a whole new...\u003c/p\u003e\n\u003cp\u003e\u003ca class=\"more-link\" href=\"https://www.irishamericanmom.com/introducing-your-irish-american-heritage-to-friends/\"\u003eRead More\u003c/a\u003e\u003c/p\u003e\nThe post \u003ca href=\"https://www.irishamericanmom.com/introducing-your-irish-american-heritage-to-friends/\"\u003eIntroducing Your Irish American Heritage To Friends\u003c/a\u003e first appeared on \u003ca href=\"https://www.irishamericanmom.com\"\u003eIrish American Mom\u003c/a\u003e.",
					"timestamp": 1726842063,
					"url": "https://www.irishamericanmom.com/introducing-your-irish-american-heritage-to-friends/",
					"content": ""
				}
		  }`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Timestamp: 1726842063,
				},
				Err: nil,
			},
		},
		{
			label: "testing tiktok parsing",
			input: `{
				"tiktok_video": {
					"id": 366467907,
					"name": "",
					"thumbnail_url": "https://p16-sign-sg.tiktokcdn.com/obj/tos-alisg-p-0037/osoj1vICMfyhlegCIxiPPoAQGeArLDCm3C8MGg?lk3s=81f88b70\u0026x-expires=1730811600\u0026x-signature=MiWrD7ZJRVFAgOcGmdjvyeODaAs%3D\u0026shp=81f88b70\u0026shcp=-",
					"link": "https://www.tiktok.com/@dyh8iyyvouy/video/7376399341298273554",
					"comments": 1,
					"likes": 42,
					"timestamp": 1717451821,
					"post_id": "7376399341298273554",
					"plays": 163,
					"shares": 0
				}
		  }`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Timestamp: 1717451821,
					Comments:  client.ToIntPointer(1),
					Likes:     client.ToIntPointer(42),
				},
				Err: nil,
			},
		},
		{
			label: "testing instagram parsing",
			input: `{
				"instagram_media": {
					"id": 2015929896,
					"text": "Makeup tonos ðŸ¤Ž marrones y Dorados ðŸ¤Ž",
					"link": "",
					"type": "video",
					"location_name": "",
					"likes": 23,
					"comments": 1,
					"timestamp": 1664749733,
					"post_id": "CjOmELToNoK",
					"views": 90,
					"mtype": 3,
					"thumbnail_url": "https://scontent-iad3-2.cdninstagram.com/v/t51.2885-15/310300342_775891656850815_2199132035482523974_n.jpg?stp=dst-jpg_e15\u0026_nc_ht=scontent-iad3-2.cdninstagram.com\u0026_nc_cat=103\u0026_nc_ohc=grmBnRX6E2IQ7kNvgF55Qd1\u0026_nc_gid=8da641c838aa4db487387ea8dacf7a95\u0026edm=AOQ1c0wBAAAA\u0026ccb=7-5\u0026oh=00_AYD6pgsZyOvvNt6SQVrhDzchmIqENYsQp4rxxWD3VaMzvQ\u0026oe=672D4812\u0026_nc_sid=8b3546",
					"hidden_likes": false,
					"plays": 0
				}
		  }`,
			expectedOutput: &domain.Message{
				Data: &domain.MsgData{
					Timestamp: 1664749733,
					Comments:  client.ToIntPointer(1),
					Likes:     client.ToIntPointer(23),
				},
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {
		res := client.ExtractMessage(tc.input)

		if ok := matchMsg(tc.expectedOutput, res); !ok {
			t.Fail()
		}
	}
}

func matchMsg(expected, input *domain.Message) bool {
	if expected.Err != input.Err {
		fmt.Printf("expected err = %v got %v\n", expected.Err, input.Err)

		return false
	}

	return matchData(expected.Data, input.Data)
}

func matchData(expected, input *domain.MsgData) bool {
	if expected.Likes != nil && *expected.Likes != *input.Likes {
		fmt.Printf("expected like = %d got %d\n", *expected.Likes, *input.Likes)

		return false
	}

	if expected.Retweets != nil && *expected.Retweets != *input.Retweets {
		fmt.Printf("expected retweets = %d got %d\n", *expected.Retweets, *input.Retweets)

		return false
	}

	if expected.Comments != nil && *expected.Comments != *input.Comments {
		fmt.Printf("expected like = %d got %d\n", *expected.Likes, *input.Likes)

		return false
	}

	if expected.Favorites != nil && *expected.Favorites != *input.Favorites {
		fmt.Printf("expected like = %d got %d\n", *expected.Likes, *input.Likes)

		return false
	}

	if expected.Timestamp != input.Timestamp {
		fmt.Printf("expected like = %d got %d\n", expected.Timestamp, input.Timestamp)

		return false
	}

	return true
}
