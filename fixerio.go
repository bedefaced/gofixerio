package fixerio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

/*
	Package Fixerio provides a simple interface to the
	fixer.io API, a service for currency exchange rates.
*/

// Request Holds the request parameters.
type Request struct {
	httpClient http.Client
	base       string
	protocol   string
	accessKey  string
	date       string
	symbols    []string
}

// Response JSON response object.
type Response struct {
	Success   bool     `json:"success"`
	Timestamp int      `json:"timestamp"`
	Base      string   `json:"base"`
	Date      string   `json:"date"`
	Rates     rates    `json:"rates"`
	Error     apiError `json:"error"`
}

type rates map[string]float32
type apiError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

// ClientResponse Response object returned by the service.
type ClientResponse struct {
	Date  time.Time
	Base  string
	Rates rates
}

const baseURL = "data.fixer.io/api"

// New Initializes fixerio.
func New() *Request {
	return &Request{
		httpClient: *http.DefaultClient,
		base:       EUR.String(),
		protocol:   "https",
		accessKey:  "",
		date:       "",
		symbols:    make([]string, 0),
	}
}

func (f *Request) Client(httpClient *http.Client) {
	if httpClient != nil {
		f.httpClient = *httpClient
	}
}

// AccessKey Sets access key.
func (f *Request) AccessKey(accessKey string) {
	f.accessKey = accessKey
}

// Sets base currency.
func (f *Request) Base(currency string) {
	f.base = currency
}

// Secure Make the connection secure or not by setting the secure argument to true or false.
func (f *Request) Secure(secure bool) {
	if secure {
		f.protocol = "https"
	} else {
		f.protocol = "http"
	}
}

// Symbols List of currencies that should be returned.
func (f *Request) Symbols(currencies ...string) {
	f.symbols = currencies
}

// Historical Specify a date in the past to retrieve historical records.
func (f *Request) Historical(date time.Time) {
	f.date = date.Format("2006-01-02")
}

// GetRates Retrieve the exchange rates.
func (f *Request) GetRates() (ClientResponse, error) {
	url := f.GetURL()
	response, err := f.makeRequest(url)

	if err != nil {
		return ClientResponse{}, err
	}

	return response, nil
}

// GetURL Formats the URL correctly for the API Request.
func (f *Request) GetURL() string {
	var url bytes.Buffer

	url.WriteString(f.protocol)
	url.WriteString("://")
	url.WriteString(baseURL)
	url.WriteString("/")

	if f.date == "" {
		url.WriteString("latest")
	} else {
		url.WriteString(f.date)
	}

	url.WriteString("?base=")
	url.WriteString(string(f.base))

	if f.accessKey != "" {
		url.WriteString("&access_key=")
		url.WriteString(string(f.accessKey))
	}

	if len(f.symbols) >= 1 {
		url.WriteString("&symbols=")
		url.WriteString(strings.Join(f.symbols, ","))
	}

	return url.String()
}

func (f *Request) makeRequest(url string) (ClientResponse, error) {
	var response Response
	body, err := f.httpClient.Get(url)

	if err != nil {
		return ClientResponse{}, errors.New("Couldn't connect to server")
	}

	defer body.Body.Close()

	err = json.NewDecoder(body.Body).Decode(&response)

	if err != nil {
		return ClientResponse{}, fmt.Errorf("Couldn't parse Response %v", err)
	}

	var sResponse = ClientResponse{
		Date:  time.Unix(int64(response.Timestamp), 0).UTC(),
		Rates: response.Rates,
		Base:  response.Base,
	}

	if response.Success == false {
		return sResponse, errors.New(response.Error.Type)
	}

	return sResponse, nil
}
