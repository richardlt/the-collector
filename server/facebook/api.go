package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	errorsP "github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/richardlt/the-collector/server/api/errors"
)

const (
	webURL   = "https://www.facebook.com/v3.1"
	graphURL = "https://graph.facebook.com/v3.1"
)

// GenerateDialogURI .
func GenerateDialogURI(redirectURI, state string) string {
	return fmt.Sprintf(
		"%s/dialog/oauth?client_id=%s&redirect_uri=%s&state=%s&display=popup",
		webURL, config.appID, redirectURI, state,
	)
}

// GetMe .
func GetMe(token string) (*User, error) {
	res, err := http.Get(fmt.Sprintf("%s/me?access_token=%s&fields=name,email",
		graphURL, token))
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	var u User
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, errorsP.WithStack(err)
	}
	if u.Error != nil {
		logrus.Errorf("%+v", errorsP.WithStack(fmt.Errorf("%v", *u.Error)))
		return nil, errors.NewRemote()
	}

	return &u, nil
}

// GetAccessToken .
func GetAccessToken(redirectURI, code string) (*AccessToken, error) {
	res, err := http.Get(fmt.Sprintf(
		"%s/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s",
		graphURL, config.appID, redirectURI, config.appSecret, code,
	))
	if err != nil {
		logrus.Errorf("%+v", errorsP.WithStack(err))
		return nil, errors.NewRemote()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errorsP.WithStack(err)
	}

	var at AccessToken
	if err := json.Unmarshal(body, &at); err != nil {
		return nil, errorsP.WithStack(err)
	}

	if at.Error != nil {
		logrus.Errorf("%+v", errorsP.WithStack(fmt.Errorf("%v", *at.Error)))
		return nil, errors.NewRemote()
	}

	return &at, nil
}
