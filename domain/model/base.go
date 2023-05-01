package model

type PageMeta struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
	Page  int `json:"page"`
}

type GetAllModel struct {
	Meta PageMeta      `json:"meta"`
	Data []interface{} `json:"data"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type Binder interface {
	Bind(i interface{}) error
}

type Validator interface {
	Validate(data interface{}) error
}

type BindFunc func(binder Binder) (data interface{}, err error)

type ValideFunc func(validator Validator, data interface{}) error
