package curseforge

import (
	"encoding/json"
	"net/url"
)

func SearchByFingerprints(fingerPrint []uint32) (*FingerPrintSearchResult, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.Path += "fingerprint"

	d, err := json.Marshal(fingerPrint)
	req, err := post(u.String(), d)
	if err != nil {
		return nil, err
	}
	res := new(FingerPrintSearchResult)
	if err := json.Unmarshal(req, res); err != nil {
		return res, err
	}
	return res, nil
}
