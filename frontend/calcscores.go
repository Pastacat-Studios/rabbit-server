package frontend

import (
	"fmt"
	"html/template"
	"pastacat/rabbitserver/database"
	"slices"
	"sort"
)

type Score struct {
	Id    string `json:"id" db:"id"`
	Score int    `json:"score" db:"MAX(score)"`
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
		result = result + "<li>" + v.Id + ": " + fmt.Sprint(v.Score) + "</li>\n" //We already checked usernames in the connect function this is safe
	}
	fmt.Println(result)
	return template.HTML(result)
}
