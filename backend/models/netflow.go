package models

import "time"

type Netflow struct {
	ID        uint      `json:"id"`        // CSV field: id
	Timestamp time.Time `json:"timestamp"` // CSV field: timestamp
	SrcAddr   string    `json:"srcaddr"`   // CSV field: srcaddr
	DstAddr   string    `json:"dstaddr"`   // CSV field: dstaddr
	NextHop   string    `json:"nexthop"`   // CSV field: nexthop
	DPkts     uint64    `json:"dPkts"`     // CSV field: dPkts
	DOctets   uint64    `json:"dOctets"`   // CSV field: dOctets
	SrcPort   int       `json:"srcport"`   // CSV field: srcport
	DstPort   int       `json:"dstport"`   // CSV field: dstport
	Prot      string    `json:"prot"`      // CSV field: prot
	Tos       int       `json:"tos"`       // CSV field: tos
}
