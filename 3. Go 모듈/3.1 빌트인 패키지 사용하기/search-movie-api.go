package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const APIKEY = "193ef3a"

type MovieInfo struct {
	Title      string `json:"Title"` // JSON 문자열을 구조체로 마샬할 때, 매핑할 키 값을 알려줌
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
}

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}

func SearchByName(name string) (*MovieInfo, error) {
	parms := url.Values{} // 쿼리 파라미터를 처리하기 위한 map. 안전한 url 인코딩을 지원
	parms.Set("apikey", APIKEY)
	parms.Set("t", name) // t는 API 규칙에 따른 것
	siteURL := "https://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	mi := &MovieInfo{}                          // 빈 구조체 생성 및 참조를 변수에 할당
	return mi, json.Unmarshal([]byte(body), mi) // 성공하면 값이 담길 것, 언마샬에서 실패한다면 에러가 담길 것
}

func SearchById(id string) (*MovieInfo, error) {
	parms := url.Values{} // 쿼리 파라미터를 처리하기 위한 map. 안전한 url 인코딩을 지원
	parms.Set("apikey", APIKEY)
	parms.Set("i", id) // t는 API 규칙에 따른 것
	siteURL := "https://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	mi := &MovieInfo{}                          // 빈 구조체 생성 및 참조를 변수에 할당
	return mi, json.Unmarshal([]byte(body), mi) // 성공하면 값이 담길 것, 언마샬에서 실패한다면 에러가 담길 것
}

var zeroValueMovieInfo MovieInfo

func main() {
	body, err := SearchById("tt3896198")
	if err != nil {
		fmt.Println(err)
	} else if *body == zeroValueMovieInfo {
		fmt.Println("No Results")
	} else {
		fmt.Println(body.Title)
	}
	body, err = SearchByName("Game of")
	if err != nil {
		fmt.Println(err)
	} else if *body == zeroValueMovieInfo {
		fmt.Println("No Results")
	} else {
		fmt.Println(body.Title)
	}
}
