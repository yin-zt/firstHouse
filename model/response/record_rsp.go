package response

import "firstHouse/model"

type UserListRsp struct {
	Total int            `json:"total"`
	Users []model.Record `json:"records"`
}
