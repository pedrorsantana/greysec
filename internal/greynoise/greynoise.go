package greynoise

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type GreyNoiseResponse struct {
	Ip    string `json:"ip"`
	Noise bool   `json:"noise"`
	Code  string `json:"hex"`
}

type cached map[string]bool

var (
	once sync.Once
	cfg  Configuration
)

type Configuration struct {
	cfgKey      string
	cfgMaxCache int
	cacheBuf    cached
}

func New(key string, maxCache int) Configuration {
	once.Do(func() {
		cfg.cacheBuf = make(cached)
		cfg = Configuration{
			cfgKey:      key,
			cfgMaxCache: maxCache,
		}
	})
	return cfg
}

func (s *Configuration) inCache(ip string) (found bool) {
	if len(s.cacheBuf) >= s.cfgMaxCache {
		s.cacheBuf = make(cached)
	}
	return s.cacheBuf[ip] == true
}

func (s *Configuration) cache(ip string) {
	if s.cacheBuf == nil {
		s.cacheBuf = make(cached)
	}
	s.cacheBuf[ip] = true
}

func (s *Configuration) LookUpIP(ip string) (found bool, err error) {
	if ip == "" {
		return false, errors.New("Missing field ip.")
	}

	if s.inCache(ip) {
		return false, nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.greynoise.io/v2/noise/quick/%s", ip), nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("key", s.cfgKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	response := &GreyNoiseResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	if response.Noise == false {
		s.cache(ip)
	}

	//Noise bool returns if the ip is found.
	return response.Noise, nil
}
