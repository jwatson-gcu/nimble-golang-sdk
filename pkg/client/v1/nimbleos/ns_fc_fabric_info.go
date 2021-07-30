// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcFabricInfo - Fibre Channel fabric information.

// Export NsFcFabricInfoFields provides field names to use in filter parameters, for example.
var NsFcFabricInfoFields *NsFcFabricInfoStringFields

func init() {
	fieldFabricSwitchName := "fabric_switch_name"
	fieldFabricSwitchPort := "fabric_switch_port"
	fieldFabricSwitchWwnn := "fabric_switch_wwnn"
	fieldFabricSwitchWwpn := "fabric_switch_wwpn"
	fieldFabricWwn := "fabric_wwn"
	fieldFcId := "fc_id"
	fieldLoggedIn := "logged_in"

	NsFcFabricInfoFields = &NsFcFabricInfoStringFields{
		FabricSwitchName: &fieldFabricSwitchName,
		FabricSwitchPort: &fieldFabricSwitchPort,
		FabricSwitchWwnn: &fieldFabricSwitchWwnn,
		FabricSwitchWwpn: &fieldFabricSwitchWwpn,
		FabricWwn:        &fieldFabricWwn,
		FcId:             &fieldFcId,
		LoggedIn:         &fieldLoggedIn,
	}
}

type NsFcFabricInfo struct {
	// FabricSwitchName - Name of the Fibre Channel switch.
	FabricSwitchName *string `json:"fabric_switch_name,omitempty"`
	// FabricSwitchPort - Port on the Fibre Channel switch to which connection is established.
	FabricSwitchPort *string `json:"fabric_switch_port,omitempty"`
	// FabricSwitchWwnn - WWNN (World Wide Node Name) for the connected port on the fabric switch.
	FabricSwitchWwnn *string `json:"fabric_switch_wwnn,omitempty"`
	// FabricSwitchWwpn - WWPN (World Wide Port Name) for the connected port on the fabric switch.
	FabricSwitchWwpn *string `json:"fabric_switch_wwpn,omitempty"`
	// FabricWwn - WWNN(World Wide Node Name) for the Fabric Switch.
	FabricWwn *string `json:"fabric_wwn,omitempty"`
	// FcId - FCID assigned to the Fabric Channel fabric port.
	FcId *string `json:"fc_id,omitempty"`
	// LoggedIn - Login information for interface. True if interface has logged in to the Fibre Channel fabric, else false.
	LoggedIn *bool `json:"logged_in,omitempty"`
}

// Struct for NsFcFabricInfoFields
type NsFcFabricInfoStringFields struct {
	FabricSwitchName *string
	FabricSwitchPort *string
	FabricSwitchWwnn *string
	FabricSwitchWwpn *string
	FabricWwn        *string
	FcId             *string
	LoggedIn         *string
}
