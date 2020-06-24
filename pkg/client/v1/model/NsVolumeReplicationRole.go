/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsVolumeReplicationRole.</p>
 */

type NsVolumeReplicationRole string 

const (
	NSVOLUMEREPLICATIONROLE_PERIODIC_SNAPSHOT_DOWNSTREAM NsVolumeReplicationRole = "periodic_snapshot_downstream"
	NSVOLUMEREPLICATIONROLE_SYNCHRONOUS_UPSTREAM NsVolumeReplicationRole = "synchronous_upstream"
	NSVOLUMEREPLICATIONROLE_SYNCHRONOUS_DOWNSTREAM NsVolumeReplicationRole = "synchronous_downstream"
	NSVOLUMEREPLICATIONROLE_NO_REPLICATION NsVolumeReplicationRole = "no_replication"
	NSVOLUMEREPLICATIONROLE_PERIODIC_SNAPSHOT_UPSTREAM NsVolumeReplicationRole = "periodic_snapshot_upstream"

) 
