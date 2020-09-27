package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/RichardKhims/go_course/internal/currency_service/database"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UpdaterService struct {
	Config config.ApiConfig
	DB database.Database
}

type FcsApiDTO struct {
	Status   bool   `json:"status"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Response []struct {
		ID          string `json:"id"`
		Price       string `json:"price"`
		Change      string `json:"change"`
		ChgPer      string `json:"chg_per"`
		LastChanged string `json:"last_changed"`
		Symbol      string `json:"symbol"`
	} `json:"response"`
	Info struct {
		ServerTime  string `json:"server_time"`
		CreditCount int    `json:"credit_count"`
		T           string `json:"_t"`
	} `json:"info"`
}

func (updater *UpdaterService) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		courses, err := updater.DB.GetAllCourses(context.Background())
		fmt.Println(courses)
		if err != nil {
			panic("Error reading courses")
		}

		for _, course := range courses {
			fmt.Println(course)
			replacer := strings.NewReplacer("$1", course.Currency1, "$2", course.Currency2)
			apiUrl := replacer.Replace(updater.Config.UrlPattern)
			response, err := http.Get(apiUrl)
			if err != nil {
				panic("Invalid api url")
			}
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic("Invalid response body")
			}
			var dto FcsApiDTO
			err = json.Unmarshal(body, &dto)
			fmt.Println(dto)
			if err != nil {
				panic("Couldn't parse response")
			}
			price, err := strconv.ParseFloat(dto.Response[0].Price, 32)
			fmt.Println(price)
			if err != nil {
				panic("Invalid price format")
			}
			err = updater.DB.UpdateCourse(context.Background(), course.Currency1, course.Currency2, price)
			if err != nil {
				panic("Couldn't update db course")
			}
		}

		time.Sleep(time.Duration(updater.Config.Period) * time.Second)
	}
}