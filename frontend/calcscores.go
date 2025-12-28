package frontend

import (
	"pastacat/rabbitserver/database"
	"sort"
)

type id struct {
	id string `db:"id"`
}

type Score struct {
	Id    string `json:"id" db:"id"`
	Score int    `json:"score" db:"score"`
}

func CalcScores() []Score {
	ids := make([]id, 0)
	database.DB.Select(&ids, "SELECT DISTINCT id FROM scores")
	scores := make([]Score, 0)
	for _, v := range ids {
		database.DB.Select(&scores, "SELECT MAX(score) FROM scores where id = $1", v.id)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i].Score > scores[j].Score })
	return scores
}
