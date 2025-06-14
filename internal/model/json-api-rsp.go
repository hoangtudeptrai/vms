package model

// Response for one DTO
type JsonDTORsp[M any] struct {
	Code    int64  `json:"code"`
	Data    M      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// Response for list of DTO with paging
type JsonDTOListRsp[M any] struct {
	Code    int64  `json:"code"`
	Count   int64  `json:"count"`
	Data    []M    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Page    int64  `json:"page"`
	Size    int64  `json:"size"`
}

func NewJsonDTORsp[M any]() *JsonDTORsp[M] {
	var dto M
	return &JsonDTORsp[M]{
		Code:    0,
		Data:    dto,
		Message: "Success",
	}
}

func NewJsonDTOListRsp[M any]() *JsonDTOListRsp[M] {
	dtoList := make([]M, 0)
	return &JsonDTOListRsp[M]{
		Code:    0,
		Count:   0,
		Data:    dtoList,
		Message: "Success",
		Size:    0,
		Page:    1,
	}
}
