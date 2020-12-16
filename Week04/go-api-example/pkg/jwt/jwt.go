package jwt

import (
	"github/akalitt/go-errors-example/pkg/errno"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	USERID   int
	USERNAME string
	ROLE     string
)

const JWTSECRET = "4Rtg8BPKwixXy2ktDPxoMMAhRzmo9mmuZjvKONGPZZQSaJWNLijxR42qRgq0iBb5"

func Token(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	switch path[len(path)-1] {
	case "login", "logout", "favicon.ico", "token":
		c.Next()
		return
	}

	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		errno.Abort(errno.ErrMissingAuthorization, nil, c)
		return
	}

	// Parse the header to get the token part.
	t := strings.Replace(header, "Bearer ", "", 1)
	parseToken(t, JWTSECRET, c)
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func parseToken(tokenString string, secret string, c *gin.Context) {
	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		errno.Abort(errno.ErrTokenParse, err, c)
		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		USERID = int(claims["id"].(float64))
		USERNAME = claims["username"].(string)
		ROLE = claims["role"].(string)
		c.Set("userid", USERID)
		c.Set("username", USERNAME)
		c.Set("token", tokenString)
		c.Set("role", ROLE)
		// Token is invalid.
	} else {
		errno.Abort(errno.ErrTokenInvalid, nil, c)
	}
}
