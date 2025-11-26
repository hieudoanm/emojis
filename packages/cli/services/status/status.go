package status

import (
	"encoding/json"
	"fmt"
	"log"

	"status-cli/utils/requests"
	"time"
)

// Page represents the status page info
type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

// Status represents the page status
type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

// Response wraps page and status
type Response struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

// GetStatus fetches status from a URL
func GetStatus(url string, debug bool) (Response, error) {
	responseByte, getError := requests.Get(url, requests.Options{Debug: debug})
	if getError != nil {
		return Response{}, getError
	}
	var response Response
	jsonError := json.Unmarshal(responseByte, &response)
	if jsonError != nil {
		return Response{}, jsonError
	}
	return response, nil
}

// PrintFullStatus prints status with borders, color, and timestamp
func PrintFullStatus(url string, debug bool) {
	resp, err := GetStatus(url, debug)
	if err != nil {
		log.Printf("[%s] \033[31mError:\033[0m %v\n", time.Now().Format(time.RFC3339), err)
		return
	}

	// Colors
	const (
		ColorReset  = "\033[0m"
		ColorBlue   = "\033[34m"
		ColorGreen  = "\033[32m"
		ColorYellow = "\033[33m"
		ColorCyan   = "\033[36m"
	)

	timestamp := time.Now().Format(time.RFC3339)
	border := fmt.Sprintf("%s==================== STATUS PAGE ==================== [%s]%s", ColorBlue, timestamp, ColorReset)

	fmt.Println(border)
	fmt.Printf("%sPage Name    :%s %s\n", ColorCyan, ColorReset, resp.Page.Name)
	fmt.Printf("%sPage ID      :%s %s\n", ColorCyan, ColorReset, resp.Page.Id)
	fmt.Printf("%sURL          :%s %s\n", ColorCyan, ColorReset, resp.Page.Url)
	fmt.Printf("%sTime Zone    :%s %s\n", ColorCyan, ColorReset, resp.Page.TimeZone)
	fmt.Printf("%sUpdated At   :%s %s\n", ColorCyan, ColorReset, resp.Page.UpdatedAt)
	fmt.Printf("%sIndicator    :%s %s\n", ColorGreen, ColorReset, resp.Status.Indicator)
	fmt.Printf("%sDescription  :%s %s\n", ColorYellow, ColorReset, resp.Status.Description)
	fmt.Println(border)
}

// PrintDescriptiveStatus prints only the descriptive status with timestamp
func PrintDescriptiveStatus(name string, url string, debug bool) {
	resp, err := GetStatus(url, debug)
	if err != nil {
		log.Printf("[%s] \033[31mError:\033[0m %v\n", time.Now().Format(time.RFC3339), err)
		return
	}
	const (
		ColorYellow = "\033[33m"
		ColorReset  = "\033[0m"
	)
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] %s%s : %s%s\n", timestamp, ColorYellow, name, resp.Status.Description, ColorReset)
}
