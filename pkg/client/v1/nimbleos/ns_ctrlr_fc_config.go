// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsCtrlrFcConfigFields provides field names to use in filter parameters, for example.
var NsCtrlrFcConfigFields *NsCtrlrFcConfigFieldHandles

func init() {
	fieldFcPortList := "fc_port_list"
	fieldFcInterfaceList := "fc_interface_list"

	NsCtrlrFcConfigFields = &NsCtrlrFcConfigFieldHandles{
		FcPortList:      &fieldFcPortList,
		FcInterfaceList: &fieldFcInterfaceList,
	}
}

// NsCtrlrFcConfig - Controller Fibre Channel configuration.
type NsCtrlrFcConfig struct {
	// FcPortList - List of Fibre Channel ports.
	FcPortList []*NsFcPortInfo `json:"fc_port_list,omitempty"`
	// FcInterfaceList - List of Fibre Channel interfaces.
	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}

// NsCtrlrFcConfigFieldHandles provides a string representation for each AccessControlRecord field.
type NsCtrlrFcConfigFieldHandles struct {
	FcPortList      *string
	FcInterfaceList *string
}
