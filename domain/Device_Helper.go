package domain

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/johannes-kuhfuss/services_utils/logger"
)

type channelName struct {
	Name string `json:"name"`
}

func (d Device) FromRawDevice(rd RawDevice) (Device, error) {
	d.Ipv4 = rd.Ipv4
	d.MacAddress = rd.MacAddress
	d.ModelID = rd.ModelID
	d.Name = rd.Name
	d.SampleRate = rd.SampleRate
	d.ServerName = rd.ServerName
	d.Channels, _ = channelsFromRaw(rd.Channels)
	d.Services, _ = servicesFromRaw(rd.Services)
	d.Subscriptions = rd.Subscriptions
	return d, nil
}

func channelsFromRaw(rawData map[string]json.RawMessage) (Channels, error) {
	var (
		channel   Channel
		channels  []Channel
		chanData  map[string]json.RawMessage
		chanName  channelName
		convError bool = false
	)

	for key, data := range rawData {
		channel.Direction = key
		err := json.Unmarshal(data, &chanData)
		if err != nil {
			logger.Error("Could not decode channel data: ", err)
			convError = true
		}
		for id, nameData := range chanData {
			channel.Id, err = strconv.Atoi(id)
			if err != nil {
				logger.Error("Could not convert channel Id into int: ", err)
				convError = true
			}
			err = json.Unmarshal(nameData, &chanName)
			if err != nil {
				logger.Error("Could not decode channel Name: ", err)
				convError = true
			}
			channel.Name = chanName.Name
			channels = append(channels, channel)
		}
	}
	if convError {
		err := errors.New("Error during channel data convesion")
		return channels, err
	}
	return channels, nil
}

func servicesFromRaw(rawData map[string]json.RawMessage) (Services, error) {
	var (
		service   Service
		services  Services
		convError bool = false
	)

	for _, svcData := range rawData {
		err := json.Unmarshal(svcData, &service)
		if err != nil {
			logger.Error("Could not decode service data: ", err)
			convError = true
		}
		services = append(services, service)
	}
	if convError {
		err := errors.New("Error during services data convesion")
		return services, err
	}
	return services, nil
}
