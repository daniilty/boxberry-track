package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"path"

	"github.com/daniilty/boxberry-track/pkg/response"
)

// SearchResult - search response.
type SearchResult struct {
	TrackID                string `json:"track_id"`
	NameIM                 string `json:"NameIM"`
	ProgramNumber          string `json:"ProgramNumber"`
	OrderID                string `json:"order_id"`
	Weight                 string `json:"Weight"`
	DeliveryType           string `json:"delivery_type"`
	ForingParcel           string `json:"ForingParcel"`
	PointCity              string `json:"point_city"`
	PointAddress           string `json:"point_address"`
	PointPhone             string `json:"point_phone"`
	Code                   string `json:"Code"`
	DeliveryDate           string `json:"delivery_date"`
	ImCode                 string `json:"imCode"`
	IssueType              string `json:"issueType"`
	PayType                string `json:"payType"`
	IssueCondition         string `json:"issueCondition"`
	Status                 string `json:"status"`
	DeliveryIntervalStart  string `json:"deliveryIntervalStart"`
	DeliveryIntervalFin    string `json:"deliveryIntervalFin"`
	Sum                    string `json:"sum"`
	ShouldChangeOdp        string `json:"shouldChangeOdp"`
	AllowCDChange          string `json:"allowCDChange"`
	AllowStorageDateChange string `json:"allowStorageDateChange"`
}

// TrackResult - track response.
type TrackResult struct {
	Statuses []*Status `json:"Statuses"`
	Result   bool      `json:"result"`
}

// Status - track result status.
type Status struct {
	DateTime        string `json:"date_time"`
	Name            string `json:"name"`
	NameSiteForeign string `json:"name_site_foreign"`
	Status          string `json:"status"`
}

// Search - search by track number.
func (c *client) Search(ctx context.Context, orderNum string) ([]*SearchResult, error) {
	searchURL := *c.baseURL
	searchURL.Path = path.Join(searchURL.Path, "api/v1/tracking/order/get")
	searchURL.RawQuery = url.Values{
		"searchId": {
			orderNum,
		},
	}.Encode()

	resp, err := c.httpClient.Get(searchURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get: %s: %w", searchURL.String(), err)
	}
	defer resp.Body.Close()

	err = response.Validate(resp)
	if err != nil {
		return nil, err
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	res := []*SearchResult{}

	err = json.Unmarshal(bb, &res)

	return res, err
}

// Track - track by trackID.
func (c *client) Track(ctx context.Context, trackID string) (*TrackResult, error) {
	searchURL := *c.baseURL
	searchURL.Path = path.Join(searchURL.Path, "api/v1/tracking/status/get")
	searchURL.RawQuery = url.Values{
		"trackId": {
			trackID,
		},
	}.Encode()

	resp, err := c.httpClient.Get(searchURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get: %s: %w", searchURL.String(), err)
	}
	defer resp.Body.Close()

	err = response.Validate(resp)
	if err != nil {
		return nil, err
	}

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	res := &TrackResult{}

	err = json.Unmarshal(bb, res)

	return res, err
}
