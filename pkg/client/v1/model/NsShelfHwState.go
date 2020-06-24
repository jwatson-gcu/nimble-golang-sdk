/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsShelfHwState.</p>
 */

type NsShelfHwState string 

const (
	NSSHELFHWSTATE_DISCOVERING NsShelfHwState = "discovering"
	NSSHELFHWSTATE_DISCONNECTED NsShelfHwState = "disconnected"
	NSSHELFHWSTATE_VOID NsShelfHwState = "void"
	NSSHELFHWSTATE_READY NsShelfHwState = "ready"
	NSSHELFHWSTATE_FAULTY NsShelfHwState = "faulty"

) 