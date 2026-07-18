package middleware

import (
	"net/http"

	"github.com/Promise111/gin-quickstart.git/internal/utils"
	"github.com/gin-gonic/gin"
)

func SecurityHeaders(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Reject requests whose Host header does not match the expected host (helps block Host-header attacks).
		if c.Request.Host != utils.ExpectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid host header",
			})
			return
		}

		// X-Frame-Options: blocks the page from being embedded in an iframe (prevents clickjacking). DENY = nowhere.
		c.Header("X-Frame-Options", "DENY")

		// Content-Security-Policy: controls which sources can load scripts, styles, images, etc. (limits XSS and injection).
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")

		// X-XSS-Protection: legacy browser XSS filter; 1; mode=block means detect and stop rendering if an attack is found.
		c.Header("X-XSS-Protection", "1; mode=block")

		// Strict-Transport-Security (HSTS): forces browsers to use HTTPS only for this host (and subdomains) for max-age seconds.
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Referrer-Policy: controls how much of the URL is sent in the Referer header; strict-origin sends only the origin on HTTPS.
		c.Header("Referrer-Policy", "strict-origin")

		// X-Content-Type-Options: stops browsers from MIME-sniffing; they must respect the declared Content-Type (reduces some XSS).
		c.Header("X-Content-Type-Options", "nosniff")

		// Permissions-Policy: turns off or limits browser features (camera, mic, geolocation, etc.) for this page.
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	}
}
