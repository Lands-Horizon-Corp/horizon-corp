package events

import (
	"encoding/json"
	"horizon-core/internal/websocket"
)

// Define your event types
const (
	AdminCreated EventType = "admin.created"
	AdminUpdated EventType = "admin.updated"
	AdminDeleted EventType = "admin.deleted"

	BranchCreated EventType = "branch.created"
	BranchUpdated EventType = "branch.updated"
	BranchDeleted EventType = "branch.deleted"

	CompanyCreated EventType = "company.created"
	CompanyUpdated EventType = "company.updated"
	CompanyDeleted EventType = "company.deleted"

	EmployeeCreated EventType = "employee.created"
	EmployeeUpdated EventType = "employee.updated"
	EmployeeDeleted EventType = "employee.deleted"

	MediaCreated EventType = "branch.created"
	MediaUpdated EventType = "branch.updated"
	MediaDeleted EventType = "branch.deleted"

	MemberCreated EventType = "member.created"
	MemberUpdated EventType = "member.updated"
	MemberDeleted EventType = "member.deleted"

	OwnerCreated EventType = "owner.created"
	OwnerUpdated EventType = "owner.updated"
	OwnerDeleted EventType = "owner.deleted"

	PermissionCreated EventType = "permission.created"
	PermissionUpdated EventType = "permission.updated"
	PermissionDeleted EventType = "permission.deleted"

	RoleCreated EventType = "role.created"
	RoleUpdated EventType = "role.updated"
	RoleDeleted EventType = "role.deleted"
)

// Create a slice containing all event types
var AllEventTypes = []EventType{
	AdminCreated,
	AdminUpdated,
	AdminDeleted,
	BranchCreated,
	BranchUpdated,
	BranchDeleted,
	CompanyCreated,
	CompanyUpdated,
	CompanyDeleted,
	EmployeeCreated,
	EmployeeUpdated,
	EmployeeDeleted,
	MediaCreated,
	MediaUpdated,
	MediaDeleted,
	MemberCreated,
	MemberUpdated,
	MemberDeleted,
	OwnerCreated,
	OwnerUpdated,
	OwnerDeleted,
	PermissionCreated,
	PermissionUpdated,
	PermissionDeleted,
	RoleCreated,
	RoleUpdated,
	RoleDeleted,
}

// RegisterEventListeners registers event listeners dynamically
func RegisterEventListeners(hub *websocket.Hub) {
	dispatcher := GetDispatcher()

	// Loop over all event types and register the same listener
	for _, eventType := range AllEventTypes {
		// Important: Capture the current value of eventType
		currentEventType := eventType

		dispatcher.Register(currentEventType, func(event Event) {
			broadcastEvent(hub, event)
		})
	}
}

func broadcastEvent(hub *websocket.Hub, event Event) {
	message, err := json.Marshal(event)
	if err != nil {
		return
	}
	hub.Broadcast <- message
}
