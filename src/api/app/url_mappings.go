package app

import (
	"github.com/zoharngo/src/api/controllers/polo"
	"github.com/zoharngo/oauth-api/src/api/controllers/oauth"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
