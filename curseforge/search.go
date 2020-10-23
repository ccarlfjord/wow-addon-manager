package curseforge

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Search searches for addon on curseforge
func Search(name string) ([]Addon, error) {
	u, err := url.Parse(baseURL)
	u.Path += "addon/search"
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("gameID", gameID)
	q.Add("searchFilter", name)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

	req, err := get(u.String())
	if err != nil {
		return nil, err
	}
	res := make([]Addon, 0)
	if err := json.Unmarshal(req, &res); err != nil {
		return res, err
	}
	return res, nil
}
