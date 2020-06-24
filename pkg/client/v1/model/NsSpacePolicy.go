/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsSpacePolicy.</p>
 */

type NsSpacePolicy string 

const (
	NSSPACEPOLICY_OFFLINE NsSpacePolicy = "offline"
	NSSPACEPOLICY_LOGIN_ONLY NsSpacePolicy = "login_only"
	NSSPACEPOLICY_NON_WRITABLE NsSpacePolicy = "non_writable"
	NSSPACEPOLICY_READ_ONLY NsSpacePolicy = "read_only"
	NSSPACEPOLICY_INVALID NsSpacePolicy = "invalid"

) 
