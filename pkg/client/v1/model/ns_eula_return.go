/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsEulaReturn - Return end-user license information.
// Export NsEulaReturnFields for advance operations like search filter etc.
var NsEulaReturnFields *NsEulaReturn

func init(){
	Eulafield:= "eula"
		
	NsEulaReturnFields= &NsEulaReturn{
		Eula: &Eulafield,
		
	}
}

type NsEulaReturn struct {
   
    // License information.
    
 	Eula *string `json:"eula,omitempty"`
}