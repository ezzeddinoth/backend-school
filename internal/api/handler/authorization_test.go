package handler

import (
	"school/internal/api/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAuthorised(t *testing.T) {
	tests := []struct {
		name               string
		endpointPermission []EndpointPermission
		userPermissions    []string
		url                string
		method             string
		expected           bool
	}{
		{
			name:               "valid adminUser test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "GET", AllRequiredPermissions: []string{"permission1", "permission2"}}},
			userPermissions:    []string{ADMIN_PERMISSION},
			url:                "/api/resource",
			method:             "GET",
			expected:           true,
		},
		{
			name:               "invalid no permissions entry no endpoint url",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "GET", AllRequiredPermissions: []string{"read"}}, {URL: "/api/resource", Method: "POST", AllRequiredPermissions: []string{"write"}}},
			userPermissions:    []string{"read", "write"},
			url:                "/api/unknown",
			method:             "GET",
			expected:           false,
		},
		{
			name:               "invalid no permissions entry no method url",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "GET", AllRequiredPermissions: []string{"read"}}, {URL: "/api/resource", Method: "POST", AllRequiredPermissions: []string{"write"}}},
			userPermissions:    []string{"read", "write"},
			url:                "/api/resource",
			method:             "DELETE",
			expected:           false,
		},
		{
			name:               "valid allRequiredPermissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "POST", AllRequiredPermissions: []string{"read", "write"}}},
			userPermissions:    []string{"read", "write"},
			url:                "/api/resource",
			method:             "POST",
			expected:           true,
		},
		{
			name:               "invalid missing allRequiredPermissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AllRequiredPermissions: []string{"read", "write"}}},
			userPermissions:    []string{"read"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           false,
		},
		{
			name:               "valid atLeastOneRequiredPermissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AtLeastOneRequiredPermissions: []string{"write", "update"}}},
			userPermissions:    []string{"update"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           true,
		},
		{
			name:               "invalid missing atLeastOneRequiredPermissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AtLeastOneRequiredPermissions: []string{"write", "update"}}},
			userPermissions:    []string{"read"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           false,
		},
		{
			name:               "valid no permissions required test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AllRequiredPermissions: []string{}, AtLeastOneRequiredPermissions: []string{}}},
			userPermissions:    []string{"read"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           true,
		},
		{
			name:               "valid allRequired and atLeastOneRequired permissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AllRequiredPermissions: []string{"read", "write"}, AtLeastOneRequiredPermissions: []string{"delete", "update"}}},
			userPermissions:    []string{"read", "write", "update"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           true,
		},
		{
			name:               "invalid atLeastOneRequired permissions test ",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AllRequiredPermissions: []string{"read", "write"}, AtLeastOneRequiredPermissions: []string{"delete", "update"}}},
			userPermissions:    []string{"read", "write"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           false,
		},
		{
			name:               "invalid allRequired permissions test",
			endpointPermission: []EndpointPermission{{URL: "/api/resource", Method: "PATCH", AllRequiredPermissions: []string{"read", "write"}, AtLeastOneRequiredPermissions: []string{"delete", "update"}}},
			userPermissions:    []string{"read", "update"},
			url:                "/api/resource",
			method:             "PATCH",
			expected:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpointPermissions = tt.endpointPermission
			user := dto.TokenUser{Permissions: tt.userPermissions}
			assert.Equal(t, tt.expected, isAuthorised(user, tt.url, tt.method))
		})
	}
}
