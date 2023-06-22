package domain

import (
	"encoding/json"
)

type Device struct {
	Channels      Channels      `json:"channels"`
	Ipv4          string        `json:"ipv4"`
	MacAddress    string        `json:"mac_address"`
	ModelID       string        `json:"model_id"`
	Name          string        `json:"name"`
	SampleRate    int           `json:"sample_rate"`
	ServerName    string        `json:"server_name"`
	Services      Services      `json:"services"`
	Subscriptions Subscriptions `json:"subscriptions"`
}

type Devices []Device

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

type Service struct {
	Ipv4       string            `json:"ipv4"`
	Name       string            `json:"name"`
	Port       int               `json:"port"`
	Properties map[string]string `json:"properties"`
	ServerName string            `json:"server_name"`
	Type       string            `json:"type"`
}

type Services []Service

type Channel struct {
	Direction string `json:"direction"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
}

type Channels []Channel
