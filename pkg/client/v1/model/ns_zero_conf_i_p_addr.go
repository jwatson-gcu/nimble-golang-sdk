/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsZeroConfIPAddr - Zero Conf of array.
// Export NsZeroConfIPAddrFields for advance operations like search filter etc.
var NsZeroConfIPAddrFields *NsZeroConfIPAddr

func init(){
	Nicfield:= "nic"
	LocalIpaddrfield:= "local_ipaddr"
	RemoteIpaddrfield:= "remote_ipaddr"
		
	NsZeroConfIPAddrFields= &NsZeroConfIPAddr{
		Nic: &Nicfield,
		LocalIpaddr: &LocalIpaddrfield,
		RemoteIpaddr: &RemoteIpaddrfield,
		
	}
}

type NsZeroConfIPAddr struct {
   
    // Nic of array.
    
 	Nic *string `json:"nic,omitempty"`
   
    // Local IP address of array.
    
 	LocalIpaddr *string `json:"local_ipaddr,omitempty"`
   
    // Remote IP address of array.
    
 	RemoteIpaddr *string `json:"remote_ipaddr,omitempty"`
}