package domain

import (
	"encoding/json"
	"strconv"

	"github.com/johannes-kuhfuss/services_utils/logger"
)

type RawDevice struct {
	Channels      map[string]json.RawMessage `json:"channels"`
	Ipv4          string                     `json:"ipv4"`
	MacAddress    string                     `json:"mac_address"`
	ModelID       string                     `json:"model_id"`
	Name          string                     `json:"name"`
	SampleRate    int                        `json:"sample_rate"`
	ServerName    string                     `json:"server_name"`
	Services      map[string]json.RawMessage `json:"services"`
	Subscriptions Subscriptions              `json:"subscriptions"`
}

type Subscription struct {
	RxChannel  string `json:"rx_channel"`
	RxDevice   string `json:"rx_device"`
	TxChannel  string `json:"tx_channel"`
	TxDevice   string `json:"tx_device"`
	StatusText string `json:"status_text"`
}

type Subscriptions []Subscription

type Channel struct {
	Direction string `json:"direction"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
}

type Device struct {
	Channels      []Channel     `json:"channels"`
	Ipv4          string        `json:"ipv4"`
	MacAddress    string        `json:"mac_address"`
	ModelID       string        `json:"model_id"`
	Name          string        `json:"name"`
	SampleRate    int           `json:"sample_rate"`
	ServerName    string        `json:"server_name"`
	Services      []Service     `json:"services"`
	Subscriptions Subscriptions `json:"subscriptions"`
}

func (d Device) FromRawDevice(rd RawDevice) (Device, error) {
	d.Ipv4 = rd.Ipv4
	d.MacAddress = rd.MacAddress
	d.ModelID = rd.ModelID
	d.Name = rd.Name
	d.SampleRate = rd.SampleRate
	d.ServerName = rd.ServerName
	d.Channels = channelsFromRaw(rd.Channels)
	d.Services = servicesFromRaw(rd.Services)
	d.Subscriptions = rd.Subscriptions
	return d, nil
}

type channelName struct {
	Name string `json:"name"`
}

func channelsFromRaw(rawData map[string]json.RawMessage) []Channel {
	var channels []Channel
	var c Channel
	var d map[string]json.RawMessage
	var cn channelName
	for k, v := range rawData {
		c.Direction = k
		err := json.Unmarshal(v, &d)
		if err != nil {
			logger.Error("Oops: ", err)
		}
		for x, y := range d {
			c.Id, err = strconv.Atoi(x)
			if err != nil {
				logger.Error("Oops: ", err)
			}
			err = json.Unmarshal(y, &cn)
			if err != nil {
				logger.Error("Oops: ", err)
			}
			c.Name = cn.Name
			channels = append(channels, c)
		}
	}
	return channels
}

func servicesFromRaw(rawData map[string]json.RawMessage) []Service {
	var services []Service
	var s Service
	for _, v := range rawData {
		err := json.Unmarshal(v, &s)
		if err != nil {
			logger.Error("Oops: ", err)
		}
		services = append(services, s)
	}
	return services
}

type Service struct {
	Ipv4       string            `json:"ipv4"`
	Name       string            `json:"name"`
	Port       int               `json:"port"`
	Properties map[string]string `json:"properties"`
	ServerName string            `json:"server_name"`
	Type       string            `json:"type"`
}
