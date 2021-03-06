package linodego

import (
	"time"

	"github.com/go-resty/resty"
)

type Notification struct {
	UntilStr string `json:"until"`
	WhenStr  string `json:"when"`

	Label    string
	Message  string
	Type     string
	Severity string
	Entity   *NotificationEntity
	Until    *time.Time `json:"-"`
	When     *time.Time `json:"-"`
}

// NotificationEntity adds detailed information about the Notification.
// This could refer to the ticket that triggered the notification, for example.
type NotificationEntity struct {
	ID    int
	Label string
	Type  string
	URL   string
}

// NotificationsPagedResponse represents a paginated Notifications API response
type NotificationsPagedResponse struct {
	*PageOptions
	Data []*Notification
}

// endpoint gets the endpoint URL for Notification
func (NotificationsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Notifications.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends Notifications when processing paginated Notification responses
func (resp *NotificationsPagedResponse) appendData(r *NotificationsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Notifications
func (NotificationsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(NotificationsPagedResponse{})
}

// ListNotifications gets a collection of Notification objects representing important,
// often time-sensitive items related to the Account. An account cannot interact directly with
// Notifications, and a Notification will disappear when the circumstances causing it
// have been resolved. For example, if the account has an important Ticket open, a response
// to the Ticket will dismiss the Notification.
func (c *Client) ListNotifications(opts *ListOptions) ([]*Notification, error) {
	response := NotificationsPagedResponse{}
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
func (v *Notification) fixDates() *Notification {
	v.Until, _ = parseDates(v.UntilStr)
	v.When, _ = parseDates(v.WhenStr)
	return v
}
