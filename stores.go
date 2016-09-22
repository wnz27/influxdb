package mrfusion

import (
	"time"

	"golang.org/x/net/context"
)

// Permission is a specific allowance for `User` or `Role`.
type Permission string
type Permissions []Permission

// Represents an authenticated user.
type User struct {
	ID          int
	Name        string
	Permissions Permissions
	Roles       []Role
}

// Role is a set of permissions that may be associated with `User`s
type Role struct {
	ID          int
	Name        string
	Permissions Permissions
	Users       []User
}

// AuthStore is the Storage and retrieval of authentication information
type AuthStore struct {
	Permissions interface {
		// Returns a list of all possible permissions support by the AuthStore.
		All(context.Context) (Permissions, error)
	}
	// User management for the AuthStore
	Users interface {
		// Create a new User in the AuthStore
		Add(context.Context, User) error
		// Delete the User from the AuthStore
		Delete(context.Context, User) error
		// Retrieve a user if `ID` exists.
		Get(ctx context.Context, ID int) error
		// Update the user's permissions or roles
		Update(context.Context, User) error
	}

	// Roles are sets of permissions.
	Roles interface {
		// Create a new role to encapsulate a set of permissions.
		Add(context.Context, Role) error
		// Delete the role
		Delete(context.Context, Role) error
		// Retrieve the role and the associated users if `ID` exists.
		Get(ctx context.Context, ID int) error
		// Update the role to change permissions or users.
		Update(context.Context, Role) error
	}
}

// UserID is a unique ID for a source user.
type UserID int

// ExplorationID is a unique ID for an exploration.
type ExplorationID int

// Exploration is a serialization of front-end Data Explorer.
type Exploration struct {
	ID        ExplorationID
	Name      string    // User provided name of the exploration.
	UserID    UserID    // UserID is the owner of this exploration.
	Data      string    // Opaque blob of JSON data
	CreatedAt time.Time // Time the exploration was first created
	UpdatedAt time.Time // Latest time the exploration was updated.
}

// ExplorationStore stores front-end serializations of data explorer sessions.
type ExplorationStore interface {
	// Search the ExplorationStore for all explorations owned by userID.
	Query(ctx context.Context, userID UserID) ([]Exploration, error)
	// Create a new Exploration in the ExplorationStore
	Add(context.Context, Exploration) error
	// Delete the exploration from the ExplorationStore
	Delete(context.Context, Exploration) error
	// Retrieve an exploration if `ID` exists.
	Get(ctx context.Context, ID ExplorationID) (Exploration, error)
	// Update the exploration; will update `UpdatedAt`.
	Update(context.Context, Exploration) error
}

// Cell is a rectangle and multiple time series queries to visualize.
type Cell struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Queries []Query
}

// Dashboard is a collection of Cells for visualization
type Dashboard struct {
	ID    int
	Cells []Cell
}

// DashboardStore stores dashboards and associated Cells
type DashboardStore interface {
	// Add creates a new dashboard in the DashboardStore
	Add(context.Context, Dashboard) error
	// Delete the dashboard from the store
	Delete(context.Context, Dashboard) error
	// Get retrieves Dashboard if `ID` exists
	Get(ctx context.Context, ID int) error
	// Update the dashboard in the store.
	Update(context.Context, Dashboard) error
}
