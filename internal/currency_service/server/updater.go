package server

import (
	"context"
	"encoding/json"
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
		if err != nil {
			panic("Error reading courses")
		}

		err = updater.updateCourses(&courses)
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(updater.Config.Period) * time.Second)
	}
}

func (updater *UpdaterService) updateCourses (courses *[]database.Course) error {
	for _, course := range *courses {
		apiUrl := updater.getApiUrl(course)
		body, err := updater.sendApiRequest(apiUrl)
		if err != nil {
			return err
		}

		dto, err := updater.unmarshalData(body)
		if err != nil {
			return err
		}

		err = updater.updateCourse(course, dto)
		if err != nil {
			return err
		}
	}

	return nil
}

func (updater *UpdaterService) getApiUrl(course database.Course) string {
	replacer := strings.NewReplacer("$1", course.Currency1, "$2", course.Currency2)
	apiUrl := replacer.Replace(updater.Config.UrlPattern)
	return apiUrl
}

func (updater *UpdaterService) sendApiRequest(apiUrl string) (*[]byte, error) {
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (updater *UpdaterService) unmarshalData (data *[]byte) (dto FcsApiDTO, err error){
	err = json.Unmarshal(*data, &dto)
	if err != nil {
		return FcsApiDTO{}, err
	}
	return dto, nil
}

func (updater *UpdaterService) updateCourse (course database.Course, dto FcsApiDTO) error {
	price, err := strconv.ParseFloat(dto.Response[0].Price, 32)
	if err != nil {
		return err
	}
	err = updater.DB.UpdateCourse(context.Background(), course.Currency1, course.Currency2, price, dto.Response[0].LastChanged)
	if err != nil {
		return err
	}
	return nil
}