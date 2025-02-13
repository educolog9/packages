package enums

// Role represents a user role in the system.
// @name Role
type Role string

const (
	// Admin represents the role of an administrator.
	// @description Administrator role
	Admin Role = "admin"

	// User represents the role of a regular user.
	// @description Regular user role
	User Role = "user"

	// Author represents the role of an author.
	// @description Author role
	Author Role = "author"

	// Editor represents the role of an editor.
	// @description Editor role
	Editor Role = "editor"

	// HR represents the role of a human resources manager.
	// @description Human resources manager role
	DirectorRRHH Role = "director_rrhh"

	// CoordinatorRRHH represents the role of a human resources coordinator.
	// @description Human resources coordinator role
	CoordinatorRRHH Role = "coordinator_rrhh"

	// OrgLeader represents the role of an organization leader user.
	// @description Organization leader role
	OrgLeader Role = "org_leader"

	// OrgUser represents the role of an organization user.
	// @description Organization user role
	OrgUser Role = "org_user"
)

// Admins is a list of admin roles for use in the application.
var ADMIN_GROUP = []Role{Admin}

// EDITORS_GROUP is a list of Editor and Admin roles for use in the application.
var EDITORS_GROUP = []Role{Admin, Editor}

// AUTHORS_GROUP is a list of Author and Admin roles for use in the application.
var AUTHORS_GROUP = []Role{Admin, Author}

// HR_ADMIN_GROUP is a list of DirectorRRHH and Admin roles for use in the application.
var HR_ADMIN_GROUP = []Role{Admin, DirectorRRHH, OrgLeader}

// HR_COORDINATOR_GROUP is a list of CoordinatorRRHH and Admin roles for use in the application.
var HR_COORDINATOR_GROUP = []Role{Admin, CoordinatorRRHH, DirectorRRHH, OrgLeader}

// ALL_ROLES is a list of all roles for use in the application.
var ALL_ROLES = []Role{Admin, User, Author, Editor, DirectorRRHH, CoordinatorRRHH, OrgLeader, OrgUser}

// ORGANIZATION_GROUP is a list of roles that can be assigned to an organization.
var ORGANIZATION_GROUP = []Role{Admin, DirectorRRHH, CoordinatorRRHH, OrgLeader, OrgUser}
