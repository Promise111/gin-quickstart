package handler

// BindUnmarshaler — Gin-specific custom parsing for query & form data.
//
// WHEN:  Query/form only. Gin calls UnmarshalParam automatically — no parser tag needed.
// WHY:   Transform incoming strings before they hit your handler logic.
// HOW:   Implement UnmarshalParam(string) error on your type + binding.BindUnmarshaler check.
//
// curl "localhost:8080/api/v1/test2?birthday=2000-01-01&birthdays=2000-01-01,2000-01-02"
//
// ── Example patterns (copy & adapt) ──────────────────────────────────────────
//
// 1. Normalize email
//    type Email string
//    func (e *Email) UnmarshalParam(param string) error {
//        *e = Email(strings.ToLower(strings.TrimSpace(param)))
//        return nil
//    }
//
// 2. Strip non-digits from phone
//    type Phone string
//    func (p *Phone) UnmarshalParam(param string) error {
//        *p = Phone(strings.Map(func(r rune) rune {
//            if r >= '0' && r <= '9' { return r }
//            return -1
//        }, param))
//        return nil
//    }
//
// 3. Enum — only allow certain values
//    type Role string
//    func (r *Role) UnmarshalParam(param string) error {
//        switch param {
//        case "admin", "user", "guest":
//            *r = Role(param)
//            return nil
//        }
//        return errors.New("role must be admin, user, or guest")
//    }
//
// 4. Validated date string
//    type Date string
//    func (d *Date) UnmarshalParam(param string) error {
//        if _, err := time.Parse("2006-01-02", param); err != nil {
//            return errors.New("use YYYY-MM-DD")
//        }
//        *d = Date(param)
//        return nil
//    }
//
// Tag usage:
//    Birthday Birthday `form:"birthday"`                              // single value
//    IDs      []ID     `form:"ids" collection_format:"csv"`          // comma-separated
//    Dates    []Date   `form:"dates,default=2020-01-01" collection_format:"csv"`

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

var _ binding.BindUnmarshaler = (*Birthday)(nil)

func UnmarshalParam(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBind struct {
			Birthday         Birthday   `form:"birthday"`
			Birthdays        []Birthday `form:"birthdays" collection_format:"csv"`
			BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
		}

		if err := c.ShouldBindQuery(&requestBind); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": requestBind,
		})
	}
}
