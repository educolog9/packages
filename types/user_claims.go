package types

import (
	"github.com/educolog7/packages/enums"
	"github.com/golang-jwt/jwt"
)

// UserClaims represents the claims of a user in the system.
type UserClaims struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	LastName       string       `json:"lastName"`
	ProfilePicture string       `json:"profilePicture"`
	DepartmentID   string       `json:"departmentId"`
	OrganizationID string       `json:"organizationId"`
	Email          string       `json:"email"`
	Roles          []enums.Role `json:"roles"`
	IsConfirmed    bool         `json:"isConfirmed"`
	IsBlocked      bool         `json:"isBlocked"`
	jwt.StandardClaims
}

// IsAdmin checks if the user has the admin role.
func (uc *UserClaims) IsAdmin() bool {
	for _, role := range uc.Roles {
		if role == enums.Admin {
			return true
		}
	}
	return false
}

// IsUser checks if the user has the user role.
func (uc *UserClaims) IsUser() bool {
	for _, role := range uc.Roles {
		if role == enums.User || role == enums.Admin {
			return true
		}
	}
	return false
}

// IsAuthor checks if the user has the author role.
func (uc *UserClaims) IsAuthor() bool {
	for _, role := range uc.Roles {
		if role == enums.Author || role == enums.Admin {
			return true
		}
	}
	return false
}

// IsEditor checks if the user has the editor role.
func (uc *UserClaims) IsEditor() bool {
	for _, role := range uc.Roles {
		if role == enums.Editor || role == enums.Admin {
			return true
		}
	}
	return false
}

// IsDirectorRRHH checks if the user has the director_rrhh role.
func (uc *UserClaims) IsDirectorRRHH() bool {
	for _, role := range uc.Roles {
		if role == enums.DirectorRRHH || role == enums.Admin {
			return true
		}
	}
	return false
}

// IsCoordinatorRRHH checks if the user has the coordinator_rrhh role.
func (uc *UserClaims) IsCoordinatorRRHH() bool {
	for _, role := range uc.Roles {
		if role == enums.CoordinatorRRHH || role == enums.DirectorRRHH || role == enums.Admin {
			return true
		}
	}
	return false
}
