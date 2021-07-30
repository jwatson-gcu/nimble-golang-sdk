// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsThrottle - A single throttle for the partner.

// Export NsThrottleFields provides field names to use in filter parameters, for example.
var NsThrottleFields *NsThrottleStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldDescription := "description"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldThrPartnerId := "thr_partner_id"
	fieldThrDayMask := "thr_day_mask"
	fieldDays := "days"
	fieldThrAtTime := "thr_at_time"
	fieldThrUntilTime := "thr_until_time"
	fieldThrBandwidth := "thr_bandwidth"
	fieldThrBandwidthKbps := "thr_bandwidth_kbps"
	fieldThrBandwidthLimitKbps := "thr_bandwidth_limit_kbps"

	NsThrottleFields = &NsThrottleStringFields{
		ID:                    &fieldID,
		Name:                  &fieldName,
		Description:           &fieldDescription,
		CreationTime:          &fieldCreationTime,
		LastModified:          &fieldLastModified,
		ThrPartnerId:          &fieldThrPartnerId,
		ThrDayMask:            &fieldThrDayMask,
		Days:                  &fieldDays,
		ThrAtTime:             &fieldThrAtTime,
		ThrUntilTime:          &fieldThrUntilTime,
		ThrBandwidth:          &fieldThrBandwidth,
		ThrBandwidthKbps:      &fieldThrBandwidthKbps,
		ThrBandwidthLimitKbps: &fieldThrBandwidthLimitKbps,
	}
}

type NsThrottle struct {
	// ID - Id of the throttle.
	ID *string `json:"id,omitempty"`
	// Name - Name of the throttle.
	Name *string `json:"name,omitempty"`
	// Description - Description of the throttle.
	Description *string `json:"description,omitempty"`
	// CreationTime - Creation time of the throttle.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Last modification time of the throttle.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ThrPartnerId - ID of the partner object.
	ThrPartnerId *string `json:"thr_partner_id,omitempty"`
	// ThrDayMask - Mask for days that the throttle operates.
	ThrDayMask *int64 `json:"thr_day_mask,omitempty"`
	// Days - List of days that the throttle operates.
	Days *string `json:"days,omitempty"`
	// ThrAtTime - Start time set for the throttle.
	ThrAtTime *int64 `json:"thr_at_time,omitempty"`
	// ThrUntilTime - End time set for the throttle.
	ThrUntilTime *int64 `json:"thr_until_time,omitempty"`
	// ThrBandwidth - Bandwidth set for the throttle in megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
	ThrBandwidth *int64 `json:"thr_bandwidth,omitempty"`
	// ThrBandwidthKbps - Bandwidth set for the throttle in kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
	ThrBandwidthKbps *int64 `json:"thr_bandwidth_kbps,omitempty"`
	// ThrBandwidthLimitKbps - Bandwidth set for the throttle in kilobits per second or -1 to indicate that there is no limit.
	ThrBandwidthLimitKbps *int64 `json:"thr_bandwidth_limit_kbps,omitempty"`
}

// Struct for NsThrottleFields
type NsThrottleStringFields struct {
	ID                    *string
	Name                  *string
	Description           *string
	CreationTime          *string
	LastModified          *string
	ThrPartnerId          *string
	ThrDayMask            *string
	Days                  *string
	ThrAtTime             *string
	ThrUntilTime          *string
	ThrBandwidth          *string
	ThrBandwidthKbps      *string
	ThrBandwidthLimitKbps *string
}
