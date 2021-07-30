// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCertData - Detailed certificate information for an installed certificate.

// Export NsCertDataFields provides field names to use in filter parameters, for example.
var NsCertDataFields *NsCertDataStringFields

func init() {
	fieldDescription := "description"
	fieldSubject := "subject"
	fieldDnslist := "dnslist"
	fieldIplist := "iplist"
	fieldTrusted := "trusted"
	fieldPemText := "pem_text"

	NsCertDataFields = &NsCertDataStringFields{
		Description: &fieldDescription,
		Subject:     &fieldSubject,
		Dnslist:     &fieldDnslist,
		Iplist:      &fieldIplist,
		Trusted:     &fieldTrusted,
		PemText:     &fieldPemText,
	}
}

type NsCertData struct {
	// Description - Complete certificate description.
	Description *string `json:"description,omitempty"`
	// Subject - Certificate subject name.
	Subject *string `json:"subject,omitempty"`
	// Dnslist - Comma-separated list of DNS names from the Subject Alternate Name.
	Dnslist *string `json:"dnslist,omitempty"`
	// Iplist - Comma-separated list of IP addresses from the Subject Alternate Name.
	Iplist *string `json:"iplist,omitempty"`
	// Trusted - Certificate is an imported, trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// PemText - PEM text of the actual certificate.
	PemText *string `json:"pem_text,omitempty"`
}

// Struct for NsCertDataFields
type NsCertDataStringFields struct {
	Description *string
	Subject     *string
	Dnslist     *string
	Iplist      *string
	Trusted     *string
	PemText     *string
}
