/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsLunReturn - Return LU number.
// Export NsLunReturnFields for advance operations like search filter etc.
var NsLunReturnFields *NsLunReturn

func init(){
		
	NsLunReturnFields= &NsLunReturn{
		
	}
}

type NsLunReturn struct {
   
    // LU number in hexadecimal.
    
   	Lun *int64 `json:"lun,omitempty"`
   
    // LU number in decimal.
    
   	LuNumber *int64 `json:"lu_number,omitempty"`
}