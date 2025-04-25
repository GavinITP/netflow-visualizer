package models

type FileRecord struct {
	ID             uint   `gorm:"primaryKey;column:id"    json:"id"`
	FileName       string `gorm:"column:file_name"        json:"file_name"`
	Count          uint64 `gorm:"column:count"            json:"count"`
	ProtTCPCount   uint64 `gorm:"column:prot_tcp_count"   json:"prot_tcp_count"`
	ProtUDPCount   uint64 `gorm:"column:prot_udp_count"   json:"prot_udp_count"`
	ProtICMPCount  uint64 `gorm:"column:prot_icmp_count"  json:"prot_icmp_count"`
	ProtOtherCount uint64 `gorm:"column:prot_other_count" json:"prot_other_count"`
	SumDPkts       uint64 `gorm:"column:sum_dPkts"        json:"sum_dPkts"`
	SumDOctets     uint64 `gorm:"column:sum_dOctets"      json:"sum_dOctets"`
}

func (FileRecord) TableName() string {
	return "file_records"
}
