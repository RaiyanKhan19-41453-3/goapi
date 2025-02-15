package middleware

import(
	"errors"
	"net/http"

	"github.com/RaiyanKhan19-41453-3/goapi/api"
	"github.com/RaiyanKhan19-41453-3/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")
