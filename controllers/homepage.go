package controllers

func (c *Client) RenderHomepage() (string, interface{}, error) {
	return "home.gotmpl", nil, nil
}