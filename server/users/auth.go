package users

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	errorsP "github.com/pkg/errors"

	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/facebook"
	"github.com/richardlt/the-collector/server/types"
)

// HandleLogin .
func HandleLogin(c echo.Context) error {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(5 * time.Minute).Unix(),
	})

	state, err := t.SignedString([]byte(config.jwtSecret))
	if err != nil {
		return errorsP.WithStack(err)
	}

	uri := facebook.GenerateDialogURI(config.redirectURI, state)

	return c.Redirect(http.StatusMovedPermanently, uri)
}

// HandleCallback .
func HandleCallback(c echo.Context) error {
	state := c.QueryParam("state")

	_, valid, err := api.ParseJWT(config.jwtSecret, state)
	if err != nil || !valid {
		return errors.NewAction("invalid state value")
	}

	code := c.QueryParam("code")

	start := time.Now()
	at, err := facebook.GetAccessToken(config.redirectURI, code)
	if err != nil {
		return err
	}
	expires := start.Add(time.Duration(at.ExpiresIn) * time.Second)

	me, err := facebook.GetMe(at.AccessToken)
	if err != nil {
		return err
	}

	u, err := get(context.Background(), newCriteria().FacebookID(me.ID))
	if err != nil {
		return err
	}
	if u == nil {
		u = &types.User{
			Name:       me.Name,
			FacebookID: me.ID,
		}
		if err := create(context.Background(), u); err != nil {
			return err
		}
	} else {
		if _, err := deleteAllAuth(context.Background(), newCriteriaAuth().
			Type(types.Facebook).UserID(u.ID)); err != nil {
			return err
		}
	}

	encrypted, err := api.Encrypt([]byte(at.AccessToken), config.secret)
	if err != nil {
		return err
	}

	if err := createAuth(context.Background(), &types.Auth{
		UserID:      u.ID,
		Type:        types.Facebook,
		AccessToken: string(encrypted),
		Expires:     expires,
	}); err != nil {
		return err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": u.UUID,
		"exp":  expires.Unix(),
	})
	token, err := t.SignedString([]byte(config.jwtSecret))
	if err != nil {
		return errorsP.WithStack(err)
	}

	if err := createAuth(context.Background(), &types.Auth{
		UserID:      u.ID,
		Type:        types.Local,
		AccessToken: token,
		Expires:     expires,
	}); err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:    "_token",
		Value:   token,
		Expires: expires,
		Path:    "/",
	})

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// MiddlewareAuth .
func MiddlewareAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey:  []byte(config.jwtSecret),
			TokenLookup: fmt.Sprintf("header:%s", echo.HeaderAuthorization),
			AuthScheme:  "Bearer",
		}),
		middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			a, err := getAuth(context.Background(), newCriteriaAuth().
				Type(types.Local).AccessToken(key))
			if err != nil {
				return false, err
			}
			if a == nil {
				return false, nil
			}

			c.Set("auth", a)
			return true, nil
		}),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return echo.HandlerFunc(func(c echo.Context) (err error) {
				a := c.Get("auth").(*types.Auth)

				me, err := get(context.Background(), newCriteria().ID(a.UserID))
				if err != nil {
					return err
				}

				c.Set("me", me)
				return next(c)
			})
		},
	}
}
