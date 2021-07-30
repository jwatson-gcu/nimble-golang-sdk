// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeCollectionHandoverAttr - Arguments to handover a volume collection.

// Export NsVolumeCollectionHandoverAttrFields provides field names to use in filter parameters, for example.
var NsVolumeCollectionHandoverAttrFields *NsVolumeCollectionHandoverAttrStringFields

func init() {
	fieldID := "id"
	fieldReplicationPartnerId := "replication_partner_id"
	fieldNoReverse := "no_reverse"

	NsVolumeCollectionHandoverAttrFields = &NsVolumeCollectionHandoverAttrStringFields{
		ID:                   &fieldID,
		ReplicationPartnerId: &fieldReplicationPartnerId,
		NoReverse:            &fieldNoReverse,
	}
}

type NsVolumeCollectionHandoverAttr struct {
	// ID - ID of the volume collection to be handed over to the downstream replication partner.
	ID *string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner.
	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
	// NoReverse - Do not automatically reverse direction of replication. Using this argument will prevent the new owner from automatically replicating the volume collection to this node when the handover completes. The default behavior is to enable replication back to this node. Default: 'false'.
	NoReverse *bool `json:"no_reverse,omitempty"`
}

// Struct for NsVolumeCollectionHandoverAttrFields
type NsVolumeCollectionHandoverAttrStringFields struct {
	ID                   *string
	ReplicationPartnerId *string
	NoReverse            *string
}
