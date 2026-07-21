package utils

import (
	"github.com/gin-contrib/sessions/cookie"
)

var SessionStore = cookie.NewStore([]byte(CookieSecret))

const (
	MaxUploadSizeSingle   = 1 << 20 // 1 MB
	MaxUploadSizeMultiple = 4 << 20 // 4 MB
	APIPrefix             = "/api"
	V1Prefix              = "/v1"
	V2Prefix              = "v2"
	V3Prefix              = "v3"
	V4Prefix              = "v4"
	ExpectedHost          = "localhost:8080"
	CookieSecret          = "IhunnaPro199"
)
