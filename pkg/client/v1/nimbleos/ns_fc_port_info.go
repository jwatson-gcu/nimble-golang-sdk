// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcPortInfoFields provides field names to use in filter parameters, for example.
var NsFcPortInfoFields *NsFcPortInfoFieldHandles

func init() {
	fieldName := "name"
	fieldBusLocation := "bus_location"
	fieldPort := "port"
	fieldSlot := "slot"

	NsFcPortInfoFields = &NsFcPortInfoFieldHandles{
		Name:        &fieldName,
		BusLocation: &fieldBusLocation,
		Port:        &fieldPort,
		Slot:        &fieldSlot,
	}
}

// NsFcPortInfo - Fibre Channel port information.
type NsFcPortInfo struct {
	// Name - Name of Fibre Channel port.
	Name *string `json:"name,omitempty"`
	// BusLocation - PCI bus location of the HBA for this Fibre Channel port.
	BusLocation *string `json:"bus_location,omitempty"`
	// Port - HBA port number for this Fibre Channel port.
	Port *int64 `json:"port,omitempty"`
	// Slot - HBA slot number for this Fibre Channel port.
	Slot *int64 `json:"slot,omitempty"`
}

// NsFcPortInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcPortInfoFieldHandles struct {
	Name        *string
	BusLocation *string
	Port        *string
	Slot        *string
}
