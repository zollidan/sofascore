package parser

import "github.com/zollidan/sofascore/models"

// ============ PARSER ============

func ParseGames(events []interface{}) []models.Game {
	games := make([]models.Game, 0, len(events))

	for _, event := range events {
		eventMap, ok := event.(map[string]interface{})
		if !ok {
			continue
		}

		game := parseGame(eventMap)
		games = append(games, game)
	}

	return games
}

func parseGame(eventMap map[string]interface{}) models.Game {
	game := models.Game{}

	// ID
	if id, ok := eventMap["id"].(float64); ok {
		game.ID = int(id)
	}

	// StartTimestamp
	if ts, ok := eventMap["startTimestamp"].(float64); ok {
		game.StartTimestamp = int64(ts)
	}

	// HomeTeam
	if homeTeam, ok := eventMap["homeTeam"].(map[string]interface{}); ok {
		if name, ok := homeTeam["name"].(string); ok {
			game.HomeTeamName = name
		}
	}

	// AwayTeam
	if awayTeam, ok := eventMap["awayTeam"].(map[string]interface{}); ok {
		if name, ok := awayTeam["name"].(string); ok {
			game.AwayTeamName = name
		}
	}

	// Scores
	game.HomeScore, game.HasScore = parseScore(eventMap["homeScore"])
	game.AwayScore, _ = parseScore(eventMap["awayScore"])

	// Status
	if status, ok := eventMap["status"].(map[string]interface{}); ok {
		if statusType, ok := status["type"].(string); ok {
			game.StatusType = statusType
		}
	}

	// Tournament
	if tournament, ok := eventMap["tournament"].(map[string]interface{}); ok {
		if name, ok := tournament["name"].(string); ok {
			game.TournamentName = name
		}
	}

	// Round
	if roundInfo, ok := eventMap["roundInfo"].(map[string]interface{}); ok {
		if round, ok := roundInfo["round"].(float64); ok {
			game.Round = int(round)
		}
	}

	return game
}

func parseScore(scoreData interface{}) (int, bool) {
	scoreMap, ok := scoreData.(map[string]interface{})
	if !ok {
		return 0, false
	}

	if current, ok := scoreMap["current"].(float64); ok {
		return int(current), true
	}

	return 0, false
}
