package util

import "time"

// Http Methods
// const (
// 	Create   = "Create"
// 	Find     = "Find"
// 	FindByID = "FindByID"
// 	FindOne  = "FindOne"
// 	Count    = "Count"
// 	Exists   = "Exists"
// )

// GenericResponse general model
type GenericResponse struct {
	Data       interface{} `json:"data"`
	ResultCode int         `json:"resultCode"`
	MetaData   *MetaData   `json:"meta,omitempty"`
}

type MetaData struct {
	TotalSize int64 `json:"totalSize,omitempty"`
}

// ModelCUDBy defines CreatedBy, UpdatedBy and DeleteBy for gorm
type ModelCUDBy struct {
	CreatedBy string `json:"-"`
	UpdatedBy string `json:"-"`
	DeletedBy string `json:"-"`
}

// ModelCUDAt defines CreatedAt, UpdatedAt and DeleteAt for gorm
type ModelCUDAt struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// ModelCUD defines both  Created, Updated and Delete for [AT] and [BY] for gorm
type ModelCUD struct {
	CreatedBy string `json:"-"`
	UpdatedBy string `json:"-"`
	DeletedBy string `json:"-"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// TransientRequest is a request for transient
type TransientRequest struct {
	Type string `json:"-"`
}

// AuthorizationRequest ...
type AuthorizationRequest struct {
	UserID, Token, Platform, Domain string
}

// AuthorizationResponse ...
type AuthorizationResponse struct {
	Authorized bool
	Username   string
	// xutils.Err
}
