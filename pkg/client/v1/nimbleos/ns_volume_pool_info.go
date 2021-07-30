// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVolumePoolInfoFields provides field names to use in filter parameters, for example.
var NsVolumePoolInfoFields *NsVolumePoolInfoFieldHandles

func init() {
	fieldVolId := "vol_id"
	fieldVolName := "vol_name"
	fieldPoolId := "pool_id"
	fieldPoolName := "pool_name"

	NsVolumePoolInfoFields = &NsVolumePoolInfoFieldHandles{
		VolId:    &fieldVolId,
		VolName:  &fieldVolName,
		PoolId:   &fieldPoolId,
		PoolName: &fieldPoolName,
	}
}

// NsVolumePoolInfo - Volume information along with the pool to which it belongs to.
type NsVolumePoolInfo struct {
	// VolId - ID of the volume.
	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of the volume.
	VolName *string `json:"vol_name,omitempty"`
	// PoolId - ID of the pool to which the volume belongs to.
	PoolId *string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool to which volume belongs to.
	PoolName *string `json:"pool_name,omitempty"`
}

// NsVolumePoolInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumePoolInfoFieldHandles struct {
	VolId    *string
	VolName  *string
	PoolId   *string
	PoolName *string
}
