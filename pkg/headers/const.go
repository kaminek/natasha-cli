// MIT license

// WARNING: This file has automatically been generated on Mon, 06 May 2019 13:18:48 UTC.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package headers

/*
#include "/root/natv2/src/cli.h"
#include <stdlib.h>
#include "cgo_helpers.h"
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
)
