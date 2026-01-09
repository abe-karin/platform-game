package service

import (
	"platform-game/model"
	"sort"
) 
func BuildRanking(scores []model.Score) []model.Ranking {
rankingMap := make(map[int]int)
for _, score := range scores { rankingMap[score.UserID] += score.Points } 
ranking := make([]model.Ranking, 0, len(rankingMap)) 
for userID, points := range rankingMap { ranking = append(ranking, model.Ranking {UserID: userID, Points: points, })}
sort.Slice(ranking, func(i,j int) bool {return ranking[i].Points > ranking[j].Points}) 
return ranking}
