package repository

import (
	"encoding/json"
	"fmt"
	"moview/src/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MovieRepository interface {
	GetMoviewByID(id int) (*models.Movie, error)
	GetMovies() ([]models.Movie, error)
	FetchMovies(date string) error
}

type movieRepository struct {
	DB *gorm.DB
}

type Movie struct {
	Rank    string	`json:"rank"`
	MovieNm string	`json:"movieNm"`
	MovieCd string	`json:"movieCd"`
	AudiAcc string  `json:"audiAcc"`
}

type MovieResponse struct {
	BoxOfficeResult struct {
		DailyBoxOfficeList []Movie `json:"dailyBoxOfficeList"`
	} `json:"boxOfficeResult"`
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		DB: db,
	}
}

func (r *movieRepository) GetMoviewByID(id int) (*models.Movie, error) {
	var movie models.Movie
	if err := r.DB.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepository) GetMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.DB.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) FetchMovies(date string) error {
	apiKey := os.Getenv("MOVIE_API_KEY")
	baseUrl := "http://www.kobis.or.kr/kobisopenapi/webservice/rest/boxoffice/searchDailyBoxOfficeList.json"
	url := fmt.Sprintf("%s?key=%s&targetDt=%s", baseUrl, apiKey, date)
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var movieResponse MovieResponse
	if err := json.NewDecoder(response.Body).Decode(&movieResponse); err != nil {
		return err
	}
	if len(movieResponse.BoxOfficeResult.DailyBoxOfficeList) > 0 {
		for _, movieData := range movieResponse.BoxOfficeResult.DailyBoxOfficeList {
			rank, err := strconv.Atoi(movieData.Rank)
			if err != nil {
				return err
			}
			movieCd, err := strconv.Atoi(movieData.MovieCd)
			if err != nil {
				return err
			}
			audiAcc, err := strconv.Atoi(movieData.AudiAcc)
			if err != nil {
				return err
			}
			result := r.DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "movie_cd"}},
				DoUpdates: clause.AssignmentColumns([]string{"title", "audience", "rank", "UpdatedAt"}),
			}).Create(&models.Movie{
				MovieCd:   movieCd,
				Title:     movieData.MovieNm,
				Audience:  audiAcc,
				Rank:      rank,
				CreatedAt: time.Now(),
			})
			if result.Error != nil {
				return result.Error
			}
		}
	}

	return nil
}