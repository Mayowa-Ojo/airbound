package rbac

import "github.com/fatih/structs"

var (
	Flights = struct{ GetFlights, UpdateReservation string }{
		"flights.get-flights",
		"flights.update-reservation",
	}

	AccessControl = struct {
		CreateRole,
		GetRole,
		GrantPermission,
		CreatePermission string
	}{
		"accesscontrol.create-role",
		"accesscontrol.get-role",
		"accesscontrol.create-permission",
		"accesscontrol.grant-permission",
	}
)

func ValidatePermission(permission string) bool {
	flights := structs.Map(Flights)
	for _, v := range flights {
		if v == permission {
			return true
		}
	}

	accesscontrol := structs.Map(AccessControl)
	for _, v := range accesscontrol {
		if v == permission {
			return true
		}
	}

	return false
}
