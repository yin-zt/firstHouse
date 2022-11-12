package common

import (
	"errors"
	"firstHouse/config"
	"firstHouse/model"
	"gorm.io/gorm"
)

// 初始化mysql数据
func InitData() {
	// 是否初始化数据
	if !config.Conf.System.InitData {
		return
	}

	// 1.写入解析记录数据
	newRecords := make([]*model.Record, 0)
	records := []*model.Record{
		{
			Model:             gorm.Model{ID: 1},
			Fview:             "gz_online",
			Fzone:             "guangzhou.com",
			Fdomain:           "www.guangzhou.com",
			Ftype:             "A",
			Ftarget:           "172.10.1.1",
			Ftarget_flag:      "master",
			Fonline_status:    "1",
			Fsystem_name:      "推荐系统",
			Fbusiness_line:    "莆田业务线",
			Fcharge_person:    "liyanhong",
			Fcandidate_person: "wangfeng",
			Fdeveloper:        "chenyixun",
			Fprotocol_port:    "tcp:80",
			Fmaintain:         "0",
			Ftest_type:        "0",
			Fversion:          "0.01",
			Flevel:            "A0",
		},
		{
			Model:             gorm.Model{ID: 2},
			Fview:             "sz_online",
			Fzone:             "guangzhou.com",
			Fdomain:           "www.guangzhou.com",
			Ftype:             "A",
			Ftarget:           "172.10.1.2",
			Ftarget_flag:      "master",
			Fonline_status:    "1",
			Fsystem_name:      "推荐系统",
			Fbusiness_line:    "莆田业务线",
			Fcharge_person:    "liyanhong",
			Fcandidate_person: "wangfeng",
			Fdeveloper:        "chenyixun",
			Fprotocol_port:    "tcp:80",
			Fmaintain:         "0",
			Ftest_type:        "0",
			Fversion:          "0.01",
			Flevel:            "A0",
		},
	}

	for _, record := range records {
		err := DB.First(&record, record.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRecords = append(newRecords, record)
		}
	}

	if len(newRecords) > 0 {
		err := DB.Create(&newRecords).Error
		if err != nil {
			Log.Errorf("写入系统角色数据失败：%v", err)
		}
	}

}
