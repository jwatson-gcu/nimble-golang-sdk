/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.
// Export NsReplPairFields for advance operations like search filter etc.
var NsReplPairFields *NsReplPair

func init(){
	SrcNamefield:= "src_name"
	DstNamefield:= "dst_name"
		
	NsReplPairFields= &NsReplPair{
		SrcName: &SrcNamefield,
		DstName: &DstNamefield,
		
	}
}

type NsReplPair struct {
   
    // Name of the replicated obj on the source group.
    
 	SrcName *string `json:"src_name,omitempty"`
   
    // Name of the replicated obj on the destination group.
    
 	DstName *string `json:"dst_name,omitempty"`
}