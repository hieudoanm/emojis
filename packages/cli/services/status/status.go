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

// PrintStatus ...
func PrintStatus(url string) {
	fmt.Println(url)
	responseByte, getError := http.Get(url, http.Options{})
	if getError != nil {
		fmt.Println("Error: ", getError)
		return
	}
	// Parse response
	var response Response
	jsonError := json.Unmarshal(responseByte, &response)
	if jsonError != nil {
		fmt.Println("Error: ", jsonError)
		return
	}
	fmt.Printf("ID          : %s\n", response.Page.Id)
	fmt.Printf("URL         : %s\n", response.Page.Url)
	fmt.Printf("Name        : %s\n", response.Page.Name)
	fmt.Printf("Time Zone   : %s\n", response.Page.TimeZone)
	fmt.Printf("Updated At  : %s\n", response.Page.UpdatedAt)
	fmt.Printf("Indicator   : %s\n", response.Status.Indicator)
	fmt.Printf("Description : %s\n", response.Status.Description)
}
