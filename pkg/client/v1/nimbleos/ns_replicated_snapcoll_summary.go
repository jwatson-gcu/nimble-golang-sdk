// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.

// Export NsReplicatedSnapcollSummaryFields provides field names to use in filter parameters, for example.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummaryStringFields

func init() {
	fieldSnapcollId := "snapcoll_id"
	fieldSnapcollName := "snapcoll_name"
	fieldSnapcollCreationTime := "snapcoll_creation_time"
	fieldDownstreamPartnerName := "downstream_partner_name"
	fieldVolIds := "vol_ids"

	NsReplicatedSnapcollSummaryFields = &NsReplicatedSnapcollSummaryStringFields{
		SnapcollId:            &fieldSnapcollId,
		SnapcollName:          &fieldSnapcollName,
		SnapcollCreationTime:  &fieldSnapcollCreationTime,
		DownstreamPartnerName: &fieldDownstreamPartnerName,
		VolIds:                &fieldVolIds,
	}
}

type NsReplicatedSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
	// DownstreamPartnerName - Name of partner the snapshot collection is replicated to.
	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
	// VolIds - Id of volumes that have snapshots which are part of the snapcoll.
	VolIds []*string `json:"vol_ids,omitempty"`
}

// Struct for NsReplicatedSnapcollSummaryFields
type NsReplicatedSnapcollSummaryStringFields struct {
	SnapcollId            *string
	SnapcollName          *string
	SnapcollCreationTime  *string
	DownstreamPartnerName *string
	VolIds                *string
}
