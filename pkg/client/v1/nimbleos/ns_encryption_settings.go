// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsEncryptionSettingsFields provides field names to use in filter parameters, for example.
var NsEncryptionSettingsFields *NsEncryptionSettingsFieldHandles

func init() {
	fieldMasterKeySet := "master_key_set"
	fieldMode := "mode"
	fieldScope := "scope"
	fieldCipher := "cipher"
	fieldEncryptionActive := "encryption_active"
	fieldEncryptionKeyManager := "encryption_key_manager"

	NsEncryptionSettingsFields = &NsEncryptionSettingsFieldHandles{
		MasterKeySet:         &fieldMasterKeySet,
		Mode:                 &fieldMode,
		Scope:                &fieldScope,
		Cipher:               &fieldCipher,
		EncryptionActive:     &fieldEncryptionActive,
		EncryptionKeyManager: &fieldEncryptionKeyManager,
	}
}

// NsEncryptionSettings - Group encryption settings.
type NsEncryptionSettings struct {
	// MasterKeySet - Is the master key set (output only).
	MasterKeySet *bool `json:"master_key_set,omitempty"`
	// Mode - Mode of encryption.
	Mode *NsEncryptionMode `json:"mode,omitempty"`
	// Scope - Encryption scope.
	Scope *NsEncryptionScope `json:"scope,omitempty"`
	// Cipher - Type of encryption cipher used.
	Cipher *NsEncryptionCipher `json:"cipher,omitempty"`
	// EncryptionActive - Is encryption active (output only).
	EncryptionActive *bool `json:"encryption_active,omitempty"`
	// EncryptionKeyManager - Is the master key on local or external key manager (output only).
	EncryptionKeyManager *NsEncryptionKeyManager `json:"encryption_key_manager,omitempty"`
}

// NsEncryptionSettingsFieldHandles provides a string representation for each AccessControlRecord field.
type NsEncryptionSettingsFieldHandles struct {
	MasterKeySet         *string
	Mode                 *string
	Scope                *string
	Cipher               *string
	EncryptionActive     *string
	EncryptionKeyManager *string
}
