package status

import (
	"encoding/json"
	"fmt"

	http "status-cli/utils/http"
)

// Page ...
type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

// Status ...
type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

// Response ...
type Response struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

// GetStatus ...
func GetStatus(url string) (Response, error) {
	responseByte, getError := http.Get(url, http.Options{})
	if getError != nil {
		fmt.Println("getError: ", getError)
		return Response{}, getError
	}
	// Parse response
	var response Response
	jsonError := json.Unmarshal(responseByte, &response)
	if jsonError != nil {
		fmt.Println("jsonError: ", jsonError)
		return Response{}, jsonError
	}

	return response, nil
}

// PrintFullStatus ...
func PrintFullStatus(url string) {
	response, err := GetStatus(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// Parse response
	fmt.Printf("ID          : %s\n", response.Page.Id)
	fmt.Printf("URL         : %s\n", response.Page.Url)
	fmt.Printf("Name        : %s\n", response.Page.Name)
	fmt.Printf("Time Zone   : %s\n", response.Page.TimeZone)
	fmt.Printf("Updated At  : %s\n", response.Page.UpdatedAt)
	fmt.Printf("Indicator   : %s\n", response.Status.Indicator)
	fmt.Printf("Description : %s\n", response.Status.Description)
}

// PrintDescriptiveStatus ...
func PrintDescriptiveStatus(name string, url string) {
	response, err := GetStatus(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("%s : %s\n", name, response.Status.Description)
}
