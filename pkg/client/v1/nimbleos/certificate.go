// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Certificate - Manage certificates used by SSL/TLS.

// Export CertificateFields provides field names to use in filter parameters, for example.
var CertificateFields *CertificateStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldSubject := "subject"
	fieldDnslist := "dnslist"
	fieldIplist := "iplist"
	fieldCertList := "cert_list"
	fieldCsr := "csr"
	fieldValidDays := "valid_days"
	fieldInput := "input"
	fieldPassword := "password"
	fieldHttps := "https"
	fieldApis := "apis"
	fieldForce := "force"
	fieldCheck := "check"
	fieldPkcs12 := "pkcs12"
	fieldTrusted := "trusted"
	fieldReloadHttps := "reload_https"
	fieldReloadApis := "reload_apis"

	CertificateFields = &CertificateStringFields{
		ID:          &fieldID,
		Name:        &fieldName,
		Subject:     &fieldSubject,
		Dnslist:     &fieldDnslist,
		Iplist:      &fieldIplist,
		CertList:    &fieldCertList,
		Csr:         &fieldCsr,
		ValidDays:   &fieldValidDays,
		Input:       &fieldInput,
		Password:    &fieldPassword,
		Https:       &fieldHttps,
		Apis:        &fieldApis,
		Force:       &fieldForce,
		Check:       &fieldCheck,
		Pkcs12:      &fieldPkcs12,
		Trusted:     &fieldTrusted,
		ReloadHttps: &fieldReloadHttps,
		ReloadApis:  &fieldReloadApis,
	}
}

type Certificate struct {
	// ID - Dummy identifier (copy of name).
	ID *string `json:"id,omitempty"`
	// Name - Name of certficiate.
	Name *string `json:"name,omitempty"`
	// Subject - Subject for the certificate.
	Subject *string `json:"subject,omitempty"`
	// Dnslist - DNS list for subject alternate name.
	Dnslist *string `json:"dnslist,omitempty"`
	// Iplist - List of IP addresses for subject alternate name.
	Iplist *string `json:"iplist,omitempty"`
	// CertList - Certificate information for certificates in file.
	CertList []*NsCertData `json:"cert_list,omitempty"`
	// Csr - Certificate Signing Request information.
	Csr *NsCertData `json:"csr,omitempty"`
	// ValidDays - Number of days certificate is valid.
	ValidDays *int64 `json:"valid_days,omitempty"`
	// Input - Certificate input for import operation. The format of the input is a base-64 encoded string for either PKCS-12 input or PEM encoded certificate(s).  PEM certificate(s) may also be specified as-is, with newlines replaced by "\n".
	Input *string `json:"input,omitempty"`
	// Password - Password used to protect PKCS-12 certificate bundle.
	Password *string `json:"password,omitempty"`
	// Https - For use command, Use specified type for HTTPS access.
	Https *bool `json:"https,omitempty"`
	// Apis - For use command, use specified type for REST API access.
	Apis *bool `json:"apis,omitempty"`
	// Force - Force regeneration or import when certificate already exists.
	Force *bool `json:"force,omitempty"`
	// Check - Check existing group certificate parameters against inputs.
	Check *bool `json:"check,omitempty"`
	// Pkcs12 - Input is a PKCS-12 bundle if true.
	Pkcs12 *bool `json:"pkcs12,omitempty"`
	// Trusted - Certificate is an imported trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// ReloadHttps - HTTPS certificate changed, so needs to be reloaded.
	ReloadHttps *bool `json:"reload_https,omitempty"`
	// ReloadApis - API certificate changed, so needs to be reloaded.
	ReloadApis *bool `json:"reload_apis,omitempty"`
}

// Struct for CertificateFields
type CertificateStringFields struct {
	ID          *string
	Name        *string
	Subject     *string
	Dnslist     *string
	Iplist      *string
	CertList    *string
	Csr         *string
	ValidDays   *string
	Input       *string
	Password    *string
	Https       *string
	Apis        *string
	Force       *string
	Check       *string
	Pkcs12      *string
	Trusted     *string
	ReloadHttps *string
	ReloadApis  *string
}
