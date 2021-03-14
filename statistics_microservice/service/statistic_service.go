package service

import (
	"SA-final/statistics_microservice/models"
)

func CalculateStatistic(userList []*models.User, bookList []*models.Book) []*models.Statistic {
	statistics := []*models.Statistic{}
	for _, book := range bookList {
		counter := 0
		statistic := &models.Statistic{
			BookName:   book.Name,
			Author:     book.Author,
			Percentage: 0,
		}
		for _, user := range userList {
			for _, id := range user.BookIds {
				if book.ID == id {
					counter += 1
				}
			}
		}
		statistic.Percentage = int(float32(counter) / float32(len(userList)) *100)
		statistics = append(statistics, statistic)
	}
	return statistics
}
