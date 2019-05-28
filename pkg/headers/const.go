// MIT license

// WARNING: This file has automatically been generated on Tue, 28 May 2019 09:14:45 UTC.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package headers

/*
*/
import "C"

const (
	// RteEthdevQueueStatCntrs as defined in kaminek/natasha-cli:0
	RteEthdevQueueStatCntrs = 16
)

// NatashaCmdType as declared in kaminek/natasha-cli:0
type NatashaCmdType int32

// NatashaCmdType enumeration from kaminek/natasha-cli:0
const (
	NatashaCmdNone       = iota
	NatashaCmdStatus     = 1
	NatashaCmdExit       = 2
	NatashaCmdReload     = 3
	NatashaCmdResetStats = 4
	NatashaCmdDpdkStats  = 5
	NatashaCmdDpdkXstats = 6
	NatashaCmdAppStats   = 7
	NatashaCmdVersion    = 8
	NatashaCmdCpuUsage   = 9
)
