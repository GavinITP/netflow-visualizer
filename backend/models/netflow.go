package models

import "time"

type BaseNetflow struct {
	SrcAddr string `json:"srcaddr"`
	DstAddr string `json:"dstaddr"`
	NextHop string `json:"nexthop"`
	DPkts   uint64 `json:"dPkts"`
	DOctets uint64 `json:"dOctets"`
	SrcPort int    `json:"srcport"`
	DstPort int    `json:"dstport"`
	Prot    string `json:"prot"`
	Tos     int    `json:"tos"`
}

type NormalNetflow struct {
	BaseNetflow
	First time.Time `json:"first"`
}

type AnomalyNetflow struct {
	BaseNetflow
	Input    int       `json:"input"`
	Output   int       `json:"output"`
	First    time.Time `json:"first"`
	Last     time.Time `json:"last"`
	TCPFlags string    `json:"tcp_flags"`
	SrcAS    int       `json:"src_as"`
	DstAS    int       `json:"dst_as"`
	SrcMask  int       `json:"src_mask"`
	DstMask  int       `json:"dst_mask"`
}
