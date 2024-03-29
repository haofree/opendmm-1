package opendmm

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/junzh0u/httpx"
)

func newDocumentInUTF8(url string, getfunc func(string) (*http.Response, error)) (*goquery.Document, error) {
	resp, err := getfunc(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code %d from %s", resp.StatusCode, url)
	}
	body, err := httpx.RespBodyInUTF8(resp)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(strings.NewReader(body))
}
