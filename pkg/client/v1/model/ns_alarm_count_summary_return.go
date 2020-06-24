/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAlarmCountSummaryReturn - List of alarm count for each category.
// Export NsAlarmCountSummaryReturnFields for advance operations like search filter etc.
var NsAlarmCountSummaryReturnFields *NsAlarmCountSummaryReturn

func init(){
		
	NsAlarmCountSummaryReturnFields= &NsAlarmCountSummaryReturn{
		
	}
}

type NsAlarmCountSummaryReturn struct {
   
    // List of alarm count for each category.
    
   	AlarmSummary []*NsAlarmCount `json:"alarm_summary,omitempty"`
}