package frontend

import (
	"fmt"
	"html/template"
	"pastacat/rabbitserver/database"
	"slices"
	"sort"
	"time"
)

type Score struct {
	Id    string `json:"id" db:"id"`
	Score int    `json:"score" db:"MAX(score)"`
}

type ScoreTimestamp struct {
	Timestamp string `db:"created"`
	Score     int    `db:"score"`
}

func CalcScores() []Score {
	ids := make([]string, 0)
	database.DB.Select(&ids, "SELECT DISTINCT id FROM scores ORDER BY score DESC, created ASC")
	scores := make([]Score, 0)
	for _, v := range ids {
		temp := make([]int, 0)
		database.DB.Select(&temp, "SELECT DISTINCT score FROM scores where id = $1", v)
		scores = append(scores, Score{v, slices.Max(temp)})
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i].Score > scores[j].Score })
	return scores
}

func GenScoreList() template.HTML {
	scores := CalcScores()
	result := ""
	for _, v := range scores {
		result = result + "<li>" + v.Id + ": " + fmt.Sprint(v.Score) + "</li>\n" //We already checked usernames in the middleware this is safe
	}
	return template.HTML(result)
}

func CalcScoresUser(user string) []ScoreTimestamp {
	scores := make([]ScoreTimestamp, 0)
	database.DB.Select(&scores, "SELECT score, created FROM scores WHERE id = $1 ORDER BY created ASC", user)
	//sort.Slice(scores, func(i, j int) bool { return scores[i].Score > scores[j].Score })
	return scores
}

func GenScoreListUser(user string) template.HTML {
	scores := CalcScoresUser(user)
	result := ""
	for _, v := range scores {
		t, _ := time.Parse(time.RFC3339, v.Timestamp)
		result = result + "<li>" + fmt.Sprint(v.Score) + " (" + fmt.Sprint(t.UTC().Format(time.UnixDate)) + ")" + "</li>\n" //We already checked usernames in the middleware this is safe
	}
	return template.HTML(result)
}
