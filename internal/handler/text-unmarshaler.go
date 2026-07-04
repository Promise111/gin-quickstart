package handler

// TextUnmarshaler — stdlib custom parsing for query, form, AND JSON.
//
// WHEN:  Same type bound from JSON body + query/form. Requires parser tag on each field.
// WHY:   Transform incoming strings before they hit your handler logic.
// HOW:   Implement UnmarshalText([]byte) error + encoding.TextUnmarshaler check.
//
// curl "localhost:8080/api/v1/test1?birthday=2000-01-01&birthdays=2000-01-01,2000-01-02"
//
// ── Example patterns (copy & adapt) ──────────────────────────────────────────
//
// 1. Validated date string (works in JSON too)
//    type Date string
//    func (d *Date) UnmarshalText(text []byte) error {
//        if _, err := time.Parse("2006-01-02", string(text)); err != nil {
//            return errors.New("use YYYY-MM-DD")
//        }
//        *d = Date(text)
//        return nil
//    }
//    // JSON:  `json:"birthday,parser=encoding.TextUnmarshaler"`
//    // Query: `form:"birthday,parser=encoding.TextUnmarshaler"`
//
// 2. Normalize email
//    type Email string
//    func (e *Email) UnmarshalText(text []byte) error {
//        *e = Email(strings.ToLower(strings.TrimSpace(string(text))))
//        return nil
//    }
//
// 3. Currency string → cents (avoid float rounding)
//    type Money int64
//    func (m *Money) UnmarshalText(text []byte) error {
//        f, err := strconv.ParseFloat(string(text), 64)
//        if err != nil { return err }
//        *m = Money(int64(f * 100))
//        return nil
//    }
//    // JSON body: {"price": "19.99"} → Money(1999)
//
// 4. Slugify
//    type Slug string
//    func (s *Slug) UnmarshalText(text []byte) error {
//        *s = Slug(strings.ToLower(strings.ReplaceAll(string(text), " ", "-")))
//        return nil
//    }
//
// Tag usage — parser tag is REQUIRED (unlike BindUnmarshaler):
//    Birthday Birthday `form:"birthday,parser=encoding.TextUnmarshaler" json:"birthday,parser=encoding.TextUnmarshaler"`
//
// BindUnmarshaler vs TextUnmarshaler:
//    Query/form only     → BindUnmarshaler (simpler, no parser tag)
//    JSON + query/form   → TextUnmarshaler (this file)

import (
	"encoding"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)

func TextUmarshal(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestText struct {
			Birthday     Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
			Birthdays    []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
			BirthdaysDef []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
		}

		if bindErr := c.ShouldBindQuery(&requestText); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": bindErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Ok",
			"data":    requestText,
		})
	}
}
