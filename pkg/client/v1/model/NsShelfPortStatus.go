/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsShelfPortStatus.</p>
 */

type NsShelfPortStatus string 

const (
	NSSHELFPORTSTATUS_CONNECTED NsShelfPortStatus = "connected"
	NSSHELFPORTSTATUS_DISCONNECTED NsShelfPortStatus = "disconnected"
	NSSHELFPORTSTATUS_DISABLED NsShelfPortStatus = "disabled"
	NSSHELFPORTSTATUS_UNKNOWN NsShelfPortStatus = "unknown"

) 
