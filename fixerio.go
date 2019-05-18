package fixerio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fixerio"
	"net/http"
	"strings"
	"time"
)

/*
	Package Fixerio provides a simple interface to the
	fixer.io API, a service for currency exchange rates.
*/

const (
	baseURL = "data.fixer.io/api"
)

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

// rates is list of couples (currency, exchange_rate)
type rates map[string]float32

type apiError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

// ClientResponse is the object returned by the service.
type ClientResponse struct {
	Date  time.Time
	Base  string
	Rates rates
}

// NewRequest Initializes fixerio.
func NewRequest() *Request {
	return &Request{
		httpClient: *http.DefaultClient,
		base:       fixerio.EUR.String(),
		protocol:   "http",
		accessKey:  "",
		date:       "",
		symbols:    make([]string, 0),
	}
}

// Client sets a custom http client
func (f *Request) Client(httpClient *http.Client) {
	if httpClient != nil {
		f.httpClient = *httpClient
	}
}

// AccessKey sets access key
func (f *Request) AccessKey(accessKey string) {
	f.accessKey = accessKey
}

// Base sets base currency
func (f *Request) Base(currency string) {
	f.base = currency
}

// Secure defines whether to make the connection secure
func (f *Request) Secure(secure bool) {
	if secure {
		f.protocol = "https"
	} else {
		f.protocol = "http"
	}
}

// Symbols contains the currencies that should be returned
func (f *Request) Symbols(currencies ...string) {
	f.symbols = currencies
}

// Historical specifies a date from which to retrieve the records
func (f *Request) Historical(date time.Time) {
	f.date = date.Format("2006-01-02")
}

// GetRates is the main function to retrieve the rates data
func (f *Request) GetRates() (ClientResponse, error) {
	url := f.GetURL()
	response, err := f.makeRequest(url)

	if err != nil {
		return ClientResponse{}, err
	}

	return response, nil
}

// GetURL constructs the url and insert parameters for the API request
// example: https://data.fixer.io/api/2019-05-07?access_key=...
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
		return ClientResponse{}, errors.New("In making request: could not connect to server")
	}

	defer body.Body.Close()

	err = json.NewDecoder(body.Body).Decode(&response)

	if err != nil || !response.Success {
		return ClientResponse{}, errors.New("In making request: could not parse response or response was invalid")
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
