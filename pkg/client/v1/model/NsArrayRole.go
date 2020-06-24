/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsArrayRole.</p>
 */

type NsArrayRole string 

const (
	NSARRAYROLE_LEADER NsArrayRole = "leader"
	NSARRAYROLE_NON_MEMBER NsArrayRole = "non_member"
	NSARRAYROLE_INVALID NsArrayRole = "invalid"
	NSARRAYROLE_BACKUP_LEADER NsArrayRole = "backup_leader"
	NSARRAYROLE_MEMBER NsArrayRole = "member"
	NSARRAYROLE_FAILED NsArrayRole = "failed"

) 
