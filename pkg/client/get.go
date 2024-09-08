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

// ParcelWithStatuses - main entity.
type ParcelWithStatuses struct {
	NameIM         string    `json:"NameIM"`
	ProgramNumber  string    `json:"ProgramNumber"`
	OrderID        string    `json:"order_id"`
	Weight         string    `json:"Weight"`
	DeliveryType   string    `json:"delivery_type"`
	PointAddress   string    `json:"point_address"`
	Code           string    `json:"Code"`
	DeliveryDate   string    `json:"delivery_date"`
	ImCode         string    `json:"imCode"`
	IssueType      string    `json:"issueType"`
	PayType        string    `json:"payType"`
	IssueCondition string    `json:"issueCondition"`
	Status         string    `json:"status"`
	StatusCode     string    `json:"status_code"`
	Sum            string    `json:"sum"`
	LastUpdateDate string    `json:"lastUpdateDate"`
	Statuses       []*Status `json:"Statuses"`
}

// Status - track result status.
type Status struct {
	DateTime        string `json:"date_time"`
	Name            string `json:"name"`
	NameSiteForeign string `json:"name_site_foreign"`
	Status          string `json:"status"`
}

// SearchResult - search response.
type SearchResult struct {
	ParcelWithStatuses []*ParcelWithStatuses `json:"parcel_with_statuses"`
	Pager              struct {
		TotalParcels   int `json:"total_parcels"`
		CurrentParcels int `json:"current_parcels"`
		Offset         int `json:"offset"`
		Limit          int `json:"limit"`
	} `json:"pager"`
}

// TrackResult - track response.
type TrackResult struct {
	Statuses []*Status `json:"Statuses"`
	Result   bool      `json:"result"`
}

// Search - search by track number.
func (c *client) Search(ctx context.Context, orderNum string) (*SearchResult, error) {
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

	res := &SearchResult{}

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
