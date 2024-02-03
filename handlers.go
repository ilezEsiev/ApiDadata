package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"io/ioutil"
	"log"
	"net/http"
)

// @title dadata

// AddressSearchHandler обрабатывает запрос на поиск адресов.
// @Summary Поиск адресов с использованием Dadata API.
// @Description Обрабатывает запрос с параметрами поиска адресов и возвращает результат.
// @ID address-search-handler
// @Accept json
// @Produce json
// @Param request body RequestAddressSearch true "Параметры запроса"
// @Success 200 {object} ResponseAddress "Успешный ответ"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка при запросе к Dadata API"
// @Router /api/search/address [post]
func AddressSearchHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestAddressSearch

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	api := dadata.NewCleanApi(client.WithCredentialProvider(&keys))
	addresses, err := api.Address(context.Background(), req.Query)
	if err != nil {
		http.Error(w, "Error querying Dadata API", http.StatusInternalServerError)
		return
	}
	var convertedAddresses []*Address
	for _, a := range addresses {
		convertedAddresses = append(convertedAddresses, &Address{
			Country:      a.Country,
			Region:       a.Region,
			Area:         a.Area,
			City:         a.City,
			CityDistrict: a.CityDistrict,
			Settlement:   a.Settlement,
			Street:       a.Street,
			House:        a.House,
			Block:        a.Block,
			Flat:         a.Flat,
			PostalBox:    a.PostalBox,
			Timezone:     a.Timezone,
		})
	}
	responseData := ResponseAddress{Addresses: convertedAddresses}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Fatal(err)
		return
	}

}

// AddressGeocodeHandler обрабатывает запрос на геокодирование адреса с использованием Dadata API.
// @Summary Геокодирование адреса.
// @Description Обрабатывает запрос на геокодирование адреса и возвращает результат.
// @ID address-geocode-handler
// @Accept json
// @Produce json
// @Param request body RequestAddressGeocode true "Параметры запроса"
// @Success 200 {object} Address "Успешный ответ"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка при запросе к Dadata API"
// @Router /api/geocode/address [post]
func AddressGeocodeHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"

	req := RequestAddressGeocode{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Token 007233085e9e9f5e7fa5871f0828d87ff737f0bf")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	Adresses := Address{}
	json.Unmarshal(responseBody, &Adresses)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Fatal(err)
	}
}
