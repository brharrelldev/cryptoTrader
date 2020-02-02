package general

import (
	"encoding/json"
	"fmt"
	"github.com/brharrelldev/crytoTrader/config"
	"github.com/brharrelldev/crytoTrader/lib"
	"net/http"
)

const (
	endpointPing         = "/api/v3/ping"
	endpointTime         = "/api/v3/time"
	EndpointExchangeInfo = "/api/v3/exchangeInfo"
)

type General struct {
	baseURL         string
	pingRequestFunc func() (*pingResponse, error)
}

type checkServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type pingResponse struct {
}

//General API object
func NewGeneralAPI(c *config.Config) (*General, error) {

	return &General{
		baseURL: c.BaseURL,
	}, nil

}

//GetPing mirrors the Binance API.  This checks for liveness of the server.  It usually returns an empty response
func (g *General) GetPing() (*pingResponse, error) { //test comment
	// will be used elsewhere in application

	var pingResp *pingResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, endpointPing), nil)
	if err != nil {
		return nil, fmt.Errorf("error building new request object")
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred when attempting to build request %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&pingResp); err != nil {
		return nil, fmt.Errorf("error decoding new json value %v", err)
	}

	return pingResp, nil
}

// will be used elsewhere in application
func (g *General) CheckServiceTime() (*checkServerTimeResponse, error) {
	// will be used elsewhere in application
	var checkServerResponse *checkServerTimeResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, endpointTime), nil)
	if err != nil {
		return nil, fmt.Errorf("error building request %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&checkServerResponse); err != nil {
		return nil, fmt.Errorf("error decoding response %v", err)
	}

	return checkServerResponse, nil

}

// used for formatting to json format for display
func (p *pingResponse) ToJson() (string, error) {
	// will be used elsewhere in application

	return lib.ToJson(p)

}

// used to format to json format to relay information back to user
func (cs *checkServerTimeResponse) ToJson() (string, error) {
	// will be used elsewhere in application
	return lib.ToJson(cs)

}
