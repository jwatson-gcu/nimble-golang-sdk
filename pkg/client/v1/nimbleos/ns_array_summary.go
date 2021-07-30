// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArraySummaryFields provides field names to use in filter parameters, for example.
var NsArraySummaryFields *NsArraySummaryFieldHandles

func init() {
	fieldID := "id"
	fieldArrayId := "array_id"
	fieldName := "name"
	fieldArrayName := "array_name"

	NsArraySummaryFields = &NsArraySummaryFieldHandles{
		ID:        &fieldID,
		ArrayId:   &fieldArrayId,
		Name:      &fieldName,
		ArrayName: &fieldArrayName,
	}
}

// NsArraySummary - Array summary information.
type NsArraySummary struct {
	// ID - Array API ID.
	ID *string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId *string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName *string `json:"array_name,omitempty"`
}

// NsArraySummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsArraySummaryFieldHandles struct {
	ID        *string
	ArrayId   *string
	Name      *string
	ArrayName *string
}
