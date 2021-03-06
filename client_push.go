package jpush

import (
	"encoding/json"
	"strconv"
	"bytes"
)

func (c *Client) Push(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) UndoPush(msgID int64) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/" + string(msgID)
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GetCidPool(count int, cidType string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/cid?"
	if count > 0 {
		link += "count=" + strconv.Itoa(count)
	}
	if cidType != "" {
		link += "type=" + cidType
	}
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GroupPush(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/grouppush"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), true)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) Validate(req *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/validate"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
