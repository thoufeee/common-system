package constants

// application status

const (
	StatusActive    = "ACTIVE"
	StatusInactive  = "INACTIVE"
	StatusPending   = "PENDING"
	StatusBlocked   = "BLOCKED"
	StatusCompleted = "COMPLETED"
)

// user status
type UserStatus string

const (
	UserActive  UserStatus = "ACTIVE"
	UserBlocked UserStatus = "BLOCKED"
)
