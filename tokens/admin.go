package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/loganmac/pocketbase/core"
	"github.com/loganmac/pocketbase/models"
	"github.com/loganmac/pocketbase/tools/security"
)

// NewAdminAuthToken generates and returns a new admin authentication token.
func NewAdminAuthToken(app core.App, admin *models.Admin) (string, error) {
	return security.NewToken(
		jwt.MapClaims{"id": admin.Id, "type": TypeAdmin},
		(admin.TokenKey + app.Settings().AdminAuthToken.Secret),
		app.Settings().AdminAuthToken.Duration,
	)
}

// NewAdminResetPasswordToken generates and returns a new admin password reset request token.
func NewAdminResetPasswordToken(app core.App, admin *models.Admin) (string, error) {
	return security.NewToken(
		jwt.MapClaims{"id": admin.Id, "type": TypeAdmin, "email": admin.Email},
		(admin.TokenKey + app.Settings().AdminPasswordResetToken.Secret),
		app.Settings().AdminPasswordResetToken.Duration,
	)
}

// NewAdminFileToken generates and returns a new admin private file access token.
func NewAdminFileToken(app core.App, admin *models.Admin) (string, error) {
	return security.NewToken(
		jwt.MapClaims{"id": admin.Id, "type": TypeAdmin},
		(admin.TokenKey + app.Settings().AdminFileToken.Secret),
		app.Settings().AdminFileToken.Duration,
	)
}
