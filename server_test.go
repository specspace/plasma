package plasma

import (
	"testing"
)

func TestStatusResponse_JSON(t *testing.T) {
	tt := []struct {
		statusResponse StatusResponse
		expectedJSON   string
	}{
		{
			statusResponse: StatusResponse{
				Version: Version{
					Name:           "1.16.4",
					ProtocolNumber: 754,
				},
				MOTD: "test",
				PlayersInfo: PlayersInfo{
					MaxPlayers: 20,
				},
			},
			expectedJSON: "{\"version\":{\"name\":\"1.16.4\",\"protocol\":754},\"players\":{\"max\":20,\"online\":0," +
				"\"sample\":null},\"description\":{\"text\":\"test\"},\"favicon\":\"\"}",
		},
		{
			statusResponse: StatusResponse{
				Version: Version{
					Name:           "1.16.4",
					ProtocolNumber: 754,
				},
				MOTD: "test",
				PlayersInfo: PlayersInfo{
					MaxPlayers:    20,
					PlayersOnline: 20,
					Players: []PlayerInfo{
						{
							Name: "Haveachin",
							UUID: "1234",
						},
					},
				},
			},
			expectedJSON: "{\"version\":{\"name\":\"1.16.4\",\"protocol\":754},\"players\":{\"max\":20,\"online\":20," +
				"\"sample\":[{\"name\":\"Haveachin\",\"id\":\"1234\"}]},\"description\":{\"text\":\"test\"}," +
				"\"favicon\":\"\"}",
		},
		{
			statusResponse: StatusResponse{
				Version: Version{
					Name:           "1.16.4",
					ProtocolNumber: 754,
				},
				MOTD: "test",
				PlayersInfo: PlayersInfo{
					MaxPlayers:    20,
					PlayersOnline: 20,
					Players: []PlayerInfo{
						{
							Name: "mb175",
							UUID: "175",
						},
					},
				},
				IconPath: "./.testfiles/test-icon-64x64.png",
			},
			expectedJSON: "{\"version\":{\"name\":\"1.16.4\",\"protocol\":754},\"players\":{\"max\":20,\"online\":20," +
				"\"sample\":[{\"name\":\"mb175\",\"id\":\"175\"}]},\"description\":{\"text\":\"test\"},\"favicon\":" +
				"\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAIAAAAlC+aJAAAAAXNSR0IArs4c6QAAAARnQU1BAACx" +
				"jwv8YQUAAAAJcEhZcwAAFiUAABYlAUlSJPAAAABeSURBVGhD7c9BCQAwDACxCemz/p3Nw2Qcg0AM5NzZrwnUBGoCNYGaQE2gJlAT" +
				"qAnUBGoCNYGaQE2gJlATqAnUBGoCNYGaQE2gJlATqAnUBGoCNYGaQE2gJlD7PDD7AFpP0Q+6dA8hAAAAAElFTkSuQmCC\"}",
		},
	}

	for _, tc := range tt {
		bb, err := tc.statusResponse.JSON()
		if err != nil {
			t.Error(err)
		}

		if tc.expectedJSON != string(bb) {
			t.Errorf("got %s; want %s", tc.expectedJSON, string(bb))
		}
	}
}
