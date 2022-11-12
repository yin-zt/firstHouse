package model

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Fview             string `gorm:"type:varchar(128);not null;default:'';comment:'view的视图'" json:"view"`
	Fzone             string `gorm:"type:varchar(128);not null;default:'';comment:'zone的空间'" json:"zone"`
	Fdomain           string `gorm:"type:varchar(128);not null;default:'';comment:'域名信息'" json:"domain"`
	Ftype             string `gorm:"type:varchar(256);not null;default:'';comment:'解析类型'" json:"type"`
	Ftarget           string `gorm:"type:varchar(256);not null;default:'';comment:'解析目标(IP或域名)'" json:"target"`
	Ftarget_slave     string `gorm:"type:varchar(128);default:'';comment:'备用解析'" json:"target_slave"`
	Ftarget_flag      string `gorm:"type:varchar(128);not null;default:'';comment:'备用解析标志(master||slave)'" json:"target_flag"`
	Fonline_status    string `gorm:"type:varchar(128);not null;default:'';comment:'有效状态(0:失效；1:启用；3:剔除；4:备用)'" json:"statue"`
	Fsystem_name      string `gorm:"type:varchar(128);default:'';comment:'系统名'" json:"system"`
	Fbusiness_line    string `gorm:"type:varchar(128);not null;default:'';comment:'业务线'" json:"line"`
	Fcharge_person    string `gorm:"type:varchar(256);not null;default:'';comment:'负责人'" json:"charge"`
	Fcandidate_person string `gorm:"type:varchar(128);not null;default:'';comment:'申请人'" json:"candidate"`
	Fdeveloper        string `gorm:"type:varchar(200);default:'';comment:'研发负责人'" json:"developer"`
	Fbackup_person    string `gorm:"type:varchar(250);default:'';comment:'备份负责人'" json:"backupPerson"`
	Fprotocol_port    string `gorm:"type:varchar(128);not null;default:'';comment:'协议和端口'" json:"port"`
	Fdetection_time   string `gorm:"type:varchar(128);default:'';comment:'自定义异常检测时间'" json:"detect"`
	Ftest_type        string `gorm:"type:varchar(32);not null;default:'';comment:'验证阶段(0:通过；1：验证中；2：待上线；3：待下线)'" json:"testtype"`
	Flevel            string `gorm:"type:varchar(20);not null;default: A0;comment:'告警级别(A0,A1,A2)'" json:"level"`
	Fmaintain         string `gorm:"type:int(10);default: '0';comment:'维护状态(0：在线；1:维护；2:待下线)'" json:"maintain"`
	Fversion          string `gorm:"type:varchar(32);default: '0';comment:'版本信息'" json:"version"`
	//Fcreate_time      time.Time `gorm:"type:datetime(3);comment:'创建时间'" json:"createTime"`
	//Fmodify_time      time.Time `gorm:"type:datetime(3);comment:'修改时间'" json:"modifyTime"`
}
