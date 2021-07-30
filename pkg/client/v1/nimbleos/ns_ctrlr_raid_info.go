// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsCtrlrRaidInfoFields provides field names to use in filter parameters, for example.
var NsCtrlrRaidInfoFields *NsCtrlrRaidInfoFieldHandles

func init() {
	fieldRaidId := "raid_id"
	fieldRaidType := "raid_type"
	fieldMaxCopies := "max_copies"
	fieldCurCopies := "cur_copies"
	fieldIsResyncing := "is_resyncing"

	NsCtrlrRaidInfoFields = &NsCtrlrRaidInfoFieldHandles{
		RaidId:      &fieldRaidId,
		RaidType:    &fieldRaidType,
		MaxCopies:   &fieldMaxCopies,
		CurCopies:   &fieldCurCopies,
		IsResyncing: &fieldIsResyncing,
	}
}

// NsCtrlrRaidInfo - Information about a controller's raid configuration.
type NsCtrlrRaidInfo struct {
	// RaidId - Raid ID for this raid array.
	RaidId *int64 `json:"raid_id,omitempty"`
	// RaidType - Type of raid for this array.
	RaidType *string `json:"raid_type,omitempty"`
	// MaxCopies - Maximum number of copies.
	MaxCopies *int64 `json:"max_copies,omitempty"`
	// CurCopies - Current number of copies.
	CurCopies *int64 `json:"cur_copies,omitempty"`
	// IsResyncing - Is this raid array resynchronizing.
	IsResyncing *bool `json:"is_resyncing,omitempty"`
}

// NsCtrlrRaidInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsCtrlrRaidInfoFieldHandles struct {
	RaidId      *string
	RaidType    *string
	MaxCopies   *string
	CurCopies   *string
	IsResyncing *string
}
