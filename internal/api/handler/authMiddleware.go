package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"school/internal/api/dto"
	"school/internal/config"
	"slices"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	dto.TokenUser
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	jwtKeyfunc, err := newAuthMiddleware()
	if err != nil {
		log.Fatalf("error keycloak: %v", err)
	}

	return func(c *gin.Context) {
		claims, err := validateToken(c.GetHeader("Authorization"), jwtKeyfunc)
		if err != nil {
			log.Printf("Not Authorized: %v", err)
			c.JSON(http.StatusUnauthorized, dto.ResponseError{ErrorMessage: "Not Authorized"})
			c.Abort()
			return
		}

		user := claims.TokenUser

		if !isAuthorised(user, c.FullPath(), c.Request.Method) {
			c.JSON(http.StatusForbidden, dto.ResponseError{ErrorMessage: "Forbidden: insufficient permissions."})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func validateToken(authHeader string, jwtKeyfunc keyfunc.Keyfunc) (*Claims, error) {
	accessToken := bearerAuthToken(authHeader)
	if accessToken == "" {
		return nil, fmt.Errorf("no access token")
	}

	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, jwtKeyfunc.Keyfunc, jwt.WithValidMethods([]string{"RS256"}), jwt.WithIssuer(os.Getenv("KEYCLOAK_URL")))
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}

	claims, _ := token.Claims.(*Claims)
	return claims, nil
}

func bearerAuthToken(authHeader string) string {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}

func newAuthMiddleware() (jwtKeyfunc keyfunc.Keyfunc, err error) {
	raw, err := getJWKS()
	if err != nil {
		return nil, fmt.Errorf("could not get json web key sets: %v", err)
	}

	keyfunc, err := keyfunc.NewJWKSetJSON(raw)
	if err != nil {
		return nil, fmt.Errorf("could not create keyfunc from raw json: %v", err)
	}

	return keyfunc, nil
}

func getJWKS() (raw json.RawMessage, err error) {
	request, err := http.NewRequest(http.MethodGet, os.Getenv("KEYCLOAK_JWKS_URL"), nil)
	if err != nil {
		return raw, fmt.Errorf("could not create request: %v", err)
	}

	client := &http.Client{Transport: config.NewKeycloakTLSTransport()}

	response, err := client.Do(request)
	if err != nil {
		return raw, fmt.Errorf("could not make http request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return raw, fmt.Errorf("could not get certs: statusCode %v", response.StatusCode)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return raw, fmt.Errorf("could not read response: %v", err)
	}

	return json.RawMessage(responseBytes), nil
}

const ADMIN_PERMISSION = "ADMIN_FULL_ACCESS"

type EndpointPermission struct {
	URL                           string
	Method                        string
	AllRequiredPermissions        []string
	AtLeastOneRequiredPermissions []string
}

var endpointPermissions = []EndpointPermission{
	{
		URL:    "/api/student",
		Method: "GET",
		AllRequiredPermissions: []string{
			"READ",
		},
	},
	{
		URL:    "/api/student",
		Method: "POST",
		AllRequiredPermissions: []string{
			"CREATE",
		},
	},
	{
		URL:    "/api/student",
		Method: "PUT",
		AllRequiredPermissions: []string{
			"UPDATE",
		},
	},
	{
		URL:    "/api/student",
		Method: "DELETE",
		AllRequiredPermissions: []string{
			"DELETE",
		},
	},
	{
		URL:    "/api/menu",
		Method: "GET",
		AllRequiredPermissions: []string{
			"READ",
		},
	},
	{
		URL:    "/api/menu",
		Method: "POST",
		AllRequiredPermissions: []string{
			"CREATE",
		},
	},
	{
		URL:    "/api/menu",
		Method: "PUT",
		AllRequiredPermissions: []string{
			"UPDATE",
		},
	},
	{
		URL:    "/api/menu",
		Method: "DELETE",
		AllRequiredPermissions: []string{
			"DELETE",
		},
	},
	{
		URL:    "/api/subscribe",
		Method: "POST",
		AllRequiredPermissions: []string{
			"CREATE",
		},
	},
}

func isAuthorised(user dto.TokenUser, url, method string) bool {
	if slices.Contains(user.Permissions, ADMIN_PERMISSION) {
		return true
	}

	allPermissions, atLeastOnePermissions := getRequiredPermissionSets(url, method)
	if allPermissions == nil && atLeastOnePermissions == nil {
		return false
	}

	return hasAllRequiredPermissions(user.Permissions, allPermissions) && hasAtLeastOneRequiredPermission(user.Permissions, atLeastOnePermissions)
}

// Check if user has all permissions in a set
func hasAllRequiredPermissions(userPermissions []string, required []string) bool {
	for _, perm := range required {
		if !slices.Contains(userPermissions, perm) {
			return false
		}
	}
	return true
}

func getRequiredPermissionSets(url, method string) (all []string, atLeastOne []string) {
	for _, ep := range endpointPermissions {
		if ep.URL == url && ep.Method == method {
			return ep.AllRequiredPermissions, ep.AtLeastOneRequiredPermissions
		}
	}
	return nil, nil
}

// Check if user has one or more permissions in a set
func hasAtLeastOneRequiredPermission(userPermissions []string, required []string) bool {
	if len(required) == 0 {
		return true
	}
	for _, perm := range required {
		if slices.Contains(userPermissions, perm) {
			return true
		}
	}
	return false
}

func GetEndpointPermissions() []EndpointPermission {
	return endpointPermissions
}
