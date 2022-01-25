package enums

import "database/sql/driver"

type BoardingPolicy string

const (
	GroupBased BoardingPolicy = "GROUP_BASED"
	ZoneBased  BoardingPolicy = "ZONE_BASED"
)

func (BoardingPolicy) Values() (kinds []string) {
	for _, s := range []BoardingPolicy{GroupBased, ZoneBased} {
		kinds = append(kinds, string(s))
	}
	return
}

type FlightStatus string

const (
	FlightActive    FlightStatus = "ACTIVE"
	FlightScheduled FlightStatus = "SCHEDULED"
	FlightDelayed   FlightStatus = "DELAYED"
	FlightDeparted  FlightStatus = "DEPARTED"
	FlightArrived   FlightStatus = "ARRIVED"
	FlightCanceled  FlightStatus = "CANCELED"
)

func (FlightStatus) Values() (kinds []string) {
	for _, s := range []FlightStatus{
		FlightActive, FlightScheduled, FlightDelayed, FlightDeparted, FlightArrived, FlightCanceled,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

type AccountStatus string

const (
	AccountStatusNone        AccountStatus = "NONE"
	AccountStatusActive      AccountStatus = "ACTIVE"
	AccountStatusClosed      AccountStatus = "CLOSED"
	AccountStatusBlacklisted AccountStatus = "BLACKLISTED"
	AccountStatusBlocked     AccountStatus = "BLOCKED"
)

func (AccountStatus) Values() (kinds []string) {
	for _, s := range []AccountStatus{
		AccountStatusActive, AccountStatusClosed, AccountStatusBlacklisted, AccountStatusBlocked, AccountStatusNone,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

type SeatType string

const (
	SeatTypeRegular       SeatType = "REGULAR"
	SeatTypeEmergencyExit SeatType = "EMERGENCY_EXIT"
	SeatTypeAccessible    SeatType = "ACCESSIBLE"
)

func (SeatType) Values() (kinds []string) {
	for _, s := range []SeatType{SeatTypeRegular, SeatTypeEmergencyExit, SeatTypeAccessible} {
		kinds = append(kinds, string(s))
	}
	return
}

type SeatClass string

const (
	SeatClassEconomy     SeatClass = "ECONOMY"
	SeatClassEconomyPlus SeatClass = "ECONOMY_PLUS"
	SeatClassBusiness    SeatClass = "BUSINESS"
	SeatClassFirst       SeatClass = "FIRST"
)

func (SeatClass) Values() (kinds []string) {
	for _, s := range []SeatClass{
		SeatClassEconomy, SeatClassEconomyPlus, SeatClassBusiness, SeatClassFirst,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

type ReservationStatus string

const (
	ReservationStatusRequested ReservationStatus = "REQUESTED"
	ReservationStatusPending   ReservationStatus = "PENDING"
	ReservationStatusConfirmed ReservationStatus = "CONFIRMED"
	ReservationStatusCanceled  ReservationStatus = "CANCELED"
	ReservationStatusCheckedIn ReservationStatus = "CHECKED_IN"
	ReservationStatusAbandoned ReservationStatus = "ABANDONED"
)

func (ReservationStatus) Values() (kinds []string) {
	for _, s := range []ReservationStatus{
		ReservationStatusRequested, ReservationStatusPending, ReservationStatusConfirmed,
		ReservationStatusCanceled, ReservationStatusCheckedIn, ReservationStatusAbandoned,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

type FlightScheduleType string

const (
	FlightScheduleWeekly FlightScheduleType = "WEEKLY"
	FlightScheduleCustom FlightScheduleType = "CUSTOM"
)

func (FlightScheduleType) Values() (kinds []string) {
	for _, s := range []FlightScheduleType{FlightScheduleWeekly, FlightScheduleCustom} {
		kinds = append(kinds, string(s))
	}
	return
}

type TripType string

const (
	OneWay    TripType = "ONE_WAY"
	RoundTrip TripType = "ROUND_TRIP"
)

func (TripType) Values() (kinds []string) {
	for _, s := range []TripType{OneWay, RoundTrip} {
		kinds = append(kinds, string(s))
	}
	return
}

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RolePilot     Role = "PILOT"
	RoleCrew      Role = "CREW"
	RoleFrontDesk Role = "FRONT_DESK"
	RoleCustomer  Role = "CUSTOMER"
)

func (Role) Values() (kinds []string) {
	for _, s := range []Role{RoleAdmin, RolePilot, RoleCrew, RoleFrontDesk, RoleCustomer} {
		kinds = append(kinds, string(s))
	}
	return
}

type AdminLevel int

const (
	AdminLevelI AdminLevel = iota + 1
	AdminLevelII
	AdminLevelIII
)

type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w WeekDay) String() string {
	switch w {
	case Monday:
		return "MONDAY"
	case Tuesday:
		return "TUESDAY"
	case Wednesday:
		return "WEDNESDAY"
	case Thursday:
		return "THURSDAY"
	case Friday:
		return "FRIDAY"
	case Saturday:
		return "SATURDAY"
	case Sunday:
		return "SUNDAY"
	default:
		return ""
	}
}

func (WeekDay) Values() []string {
	return []string{
		Monday.String(), Tuesday.String(),
		Wednesday.String(), Thursday.String(),
		Friday.String(), Saturday.String(), Sunday.String(),
	}
}

func (w WeekDay) Value() (driver.Value, error) {
	return w.String(), nil
}

func (w *WeekDay) Scan(val interface{}) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}

	switch s {
	case "MONDAY":
		*w = Monday
	case "TUESDAY":
		*w = Tuesday
	case "WEDNESDAY":
		*w = Wednesday
	case "THURSDAY":
		*w = Thursday
	case "FRIDAY":
		*w = Friday
	case "SATURDAY":
		*w = Saturday
	case "SUNDAY":
		*w = Sunday
	}

	return nil
}
