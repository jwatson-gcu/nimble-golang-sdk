/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFcInterfaceUpdateInfo - Fibre Channel interface information to update.
// Export NsFcInterfaceUpdateInfoFields for advance operations like search filter etc.
var NsFcInterfaceUpdateInfoFields *NsFcInterfaceUpdateInfo

func init(){
	IDfield:= "id"
		
	NsFcInterfaceUpdateInfoFields= &NsFcInterfaceUpdateInfo{
		ID: &IDfield,
		
	}
}

type NsFcInterfaceUpdateInfo struct {
   
    // ID of Fibre Channel interface.
    
 	ID *string `json:"id,omitempty"`
   
    // Identify whether the Fibre Channel interface is online.
    
 	Online *bool `json:"online,omitempty"`
}