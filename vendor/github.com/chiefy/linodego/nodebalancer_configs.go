package linodego

import (
	"fmt"

	"github.com/go-resty/resty"
)

type NodeBalancerConfig struct {
	ID             int
	Port           int
	Protocol       string
	Algorithm      string
	Stickiness     string
	Check          string
	CheckInterval  int                     `json:"check_interval"`
	CheckAttempts  int                     `json:"check_attempts"`
	CheckPath      string                  `json:"check_path"`
	CheckBody      string                  `json:"check_body"`
	CheckPassive   bool                    `json:"check_passive"`
	CipherSuite    string                  `json:"cipher_suite"`
	NodeBalancerID int                     `json:"nodebalancer_id"`
	SSLCommonName  string                  `json:"ssl_commonname"`
	SSLFingerprint string                  `json:"ssl_fingerprint"`
	SSLCert        string                  `json:"ssl_cert"`
	SSLKey         string                  `json:"ssl_key"`
	NodesStatus    *NodeBalancerNodeStatus `json:"nodes_status"`
}

type NodeBalancerNodeStatus struct {
	Up   int
	Down int
}

// NodeBalancerConfigsPagedResponse represents a paginated NodeBalancerConfig API response
type NodeBalancerConfigsPagedResponse struct {
	*PageOptions
	Data []*NodeBalancerConfig
}

// endpoint gets the endpoint URL for NodeBalancerConfig
func (NodeBalancerConfigsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.NodeBalancerConfigs.endpointWithID(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends NodeBalancerConfigs when processing paginated NodeBalancerConfig responses
func (resp *NodeBalancerConfigsPagedResponse) appendData(r *NodeBalancerConfigsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of NodeBalancerConfig
func (NodeBalancerConfigsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(NodeBalancerConfigsPagedResponse{})
}

// ListNodeBalancerConfigs lists NodeBalancerConfigs
func (c *Client) ListNodeBalancerConfigs(nodebalancerID int, opts *ListOptions) ([]*NodeBalancerConfig, error) {
	response := NodeBalancerConfigsPagedResponse{}
	err := c.listHelperWithID(&response, nodebalancerID, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *NodeBalancerConfig) fixDates() *NodeBalancerConfig {
	return v
}

// GetNodeBalancerConfig gets the template with the provided ID
func (c *Client) GetNodeBalancerConfig(nodebalancerID int, configID int) (*NodeBalancerConfig, error) {
	e, err := c.NodeBalancerConfigs.endpointWithID(nodebalancerID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := c.R().SetResult(&NodeBalancerConfig{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerConfig).fixDates(), nil
}
