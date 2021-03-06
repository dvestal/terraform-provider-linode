package linodego

import (
	"fmt"

	"github.com/go-resty/resty"
)

// LongviewClient represents a LongviewClient object
type LongviewClient struct {
	ID int
	// UpdatedStr string `json:"updated"`
	// Updated *time.Time `json:"-"`
}

// LongviewClientsPagedResponse represents a paginated LongviewClient API response
type LongviewClientsPagedResponse struct {
	*PageOptions
	Data []*LongviewClient
}

// endpoint gets the endpoint URL for LongviewClient
func (LongviewClientsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.LongviewClients.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends LongviewClients when processing paginated LongviewClient responses
func (resp *LongviewClientsPagedResponse) appendData(r *LongviewClientsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of LongviewClient
func (LongviewClientsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(LongviewClientsPagedResponse{})
}

// ListLongviewClients lists LongviewClients
func (c *Client) ListLongviewClients(opts *ListOptions) ([]*LongviewClient, error) {
	response := LongviewClientsPagedResponse{}
	err := c.listHelper(&response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *LongviewClient) fixDates() *LongviewClient {
	// v.Created, _ = parseDates(v.CreatedStr)
	// v.Updated, _ = parseDates(v.UpdatedStr)
	return v
}

// GetLongviewClient gets the template with the provided ID
func (c *Client) GetLongviewClient(id string) (*LongviewClient, error) {
	e, err := c.LongviewClients.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, id)
	r, err := c.R().SetResult(&LongviewClient{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LongviewClient).fixDates(), nil
}
