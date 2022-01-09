// Code generated by entc, DO NOT EDIT.

package ent

import (
	"airbound/internal/ent/account"
	"airbound/internal/ent/address"
	"airbound/internal/ent/admin"
	"airbound/internal/ent/aircraft"
	"airbound/internal/ent/airline"
	"airbound/internal/ent/airport"
	"airbound/internal/ent/crew"
	"airbound/internal/ent/customer"
	"airbound/internal/ent/flight"
	"airbound/internal/ent/flightinstance"
	"airbound/internal/ent/flightreservation"
	"airbound/internal/ent/flightschedule"
	"airbound/internal/ent/flightseat"
	"airbound/internal/ent/frontdesk"
	"airbound/internal/ent/itenerary"
	"airbound/internal/ent/passenger"
	"airbound/internal/ent/permission"
	"airbound/internal/ent/pilot"
	"airbound/internal/ent/role"
	"airbound/internal/ent/schema"
	"airbound/internal/ent/seat"
	"airbound/internal/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[4].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[5].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescStreet is the schema descriptor for street field.
	addressDescStreet := addressFields[1].Descriptor()
	// address.StreetValidator is a validator for the "street" field. It is called by the builders before save.
	address.StreetValidator = addressDescStreet.Validators[0].(func(string) error)
	// addressDescCity is the schema descriptor for city field.
	addressDescCity := addressFields[2].Descriptor()
	// address.CityValidator is a validator for the "city" field. It is called by the builders before save.
	address.CityValidator = addressDescCity.Validators[0].(func(string) error)
	// addressDescState is the schema descriptor for state field.
	addressDescState := addressFields[3].Descriptor()
	// address.StateValidator is a validator for the "state" field. It is called by the builders before save.
	address.StateValidator = addressDescState.Validators[0].(func(string) error)
	// addressDescZipcode is the schema descriptor for zipcode field.
	addressDescZipcode := addressFields[4].Descriptor()
	// address.ZipcodeValidator is a validator for the "zipcode" field. It is called by the builders before save.
	address.ZipcodeValidator = addressDescZipcode.Validators[0].(func(string) error)
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressFields[5].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressFields[6].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// addressDescID is the schema descriptor for id field.
	addressDescID := addressFields[0].Descriptor()
	// address.DefaultID holds the default value on creation for the id field.
	address.DefaultID = addressDescID.Default.(func() uuid.UUID)
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescTwoFaSecret is the schema descriptor for two_fa_secret field.
	adminDescTwoFaSecret := adminFields[1].Descriptor()
	// admin.TwoFaSecretValidator is a validator for the "two_fa_secret" field. It is called by the builders before save.
	admin.TwoFaSecretValidator = adminDescTwoFaSecret.Validators[0].(func(string) error)
	// adminDescTwoFaCompleted is the schema descriptor for two_fa_completed field.
	adminDescTwoFaCompleted := adminFields[2].Descriptor()
	// admin.DefaultTwoFaCompleted holds the default value on creation for the two_fa_completed field.
	admin.DefaultTwoFaCompleted = adminDescTwoFaCompleted.Default.(bool)
	// adminDescToken is the schema descriptor for token field.
	adminDescToken := adminFields[3].Descriptor()
	// admin.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	admin.TokenValidator = adminDescToken.Validators[0].(func(string) error)
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminFields[4].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminFields[5].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminFields[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() uuid.UUID)
	aircraftFields := schema.Aircraft{}.Fields()
	_ = aircraftFields
	// aircraftDescTailNumber is the schema descriptor for tail_number field.
	aircraftDescTailNumber := aircraftFields[1].Descriptor()
	// aircraft.TailNumberValidator is a validator for the "tail_number" field. It is called by the builders before save.
	aircraft.TailNumberValidator = aircraftDescTailNumber.Validators[0].(func(string) error)
	// aircraftDescManufacturer is the schema descriptor for manufacturer field.
	aircraftDescManufacturer := aircraftFields[2].Descriptor()
	// aircraft.ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	aircraft.ManufacturerValidator = aircraftDescManufacturer.Validators[0].(func(string) error)
	// aircraftDescModel is the schema descriptor for model field.
	aircraftDescModel := aircraftFields[3].Descriptor()
	// aircraft.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	aircraft.ModelValidator = aircraftDescModel.Validators[0].(func(string) error)
	// aircraftDescCapacity is the schema descriptor for capacity field.
	aircraftDescCapacity := aircraftFields[4].Descriptor()
	// aircraft.CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	aircraft.CapacityValidator = aircraftDescCapacity.Validators[0].(func(int) error)
	// aircraftDescRange is the schema descriptor for range field.
	aircraftDescRange := aircraftFields[5].Descriptor()
	// aircraft.RangeValidator is a validator for the "range" field. It is called by the builders before save.
	aircraft.RangeValidator = aircraftDescRange.Validators[0].(func(int) error)
	// aircraftDescCreatedAt is the schema descriptor for created_at field.
	aircraftDescCreatedAt := aircraftFields[6].Descriptor()
	// aircraft.DefaultCreatedAt holds the default value on creation for the created_at field.
	aircraft.DefaultCreatedAt = aircraftDescCreatedAt.Default.(func() time.Time)
	// aircraftDescUpdatedAt is the schema descriptor for updated_at field.
	aircraftDescUpdatedAt := aircraftFields[7].Descriptor()
	// aircraft.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	aircraft.DefaultUpdatedAt = aircraftDescUpdatedAt.Default.(func() time.Time)
	// aircraft.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	aircraft.UpdateDefaultUpdatedAt = aircraftDescUpdatedAt.UpdateDefault.(func() time.Time)
	// aircraftDescID is the schema descriptor for id field.
	aircraftDescID := aircraftFields[0].Descriptor()
	// aircraft.DefaultID holds the default value on creation for the id field.
	aircraft.DefaultID = aircraftDescID.Default.(func() uuid.UUID)
	airlineFields := schema.Airline{}.Fields()
	_ = airlineFields
	// airlineDescName is the schema descriptor for name field.
	airlineDescName := airlineFields[1].Descriptor()
	// airline.NameValidator is a validator for the "name" field. It is called by the builders before save.
	airline.NameValidator = airlineDescName.Validators[0].(func(string) error)
	// airlineDescIataCode is the schema descriptor for iata_code field.
	airlineDescIataCode := airlineFields[2].Descriptor()
	// airline.IataCodeValidator is a validator for the "iata_code" field. It is called by the builders before save.
	airline.IataCodeValidator = airlineDescIataCode.Validators[0].(func(string) error)
	// airlineDescCreatedAt is the schema descriptor for created_at field.
	airlineDescCreatedAt := airlineFields[3].Descriptor()
	// airline.DefaultCreatedAt holds the default value on creation for the created_at field.
	airline.DefaultCreatedAt = airlineDescCreatedAt.Default.(func() time.Time)
	// airlineDescUpdatedAt is the schema descriptor for updated_at field.
	airlineDescUpdatedAt := airlineFields[4].Descriptor()
	// airline.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	airline.DefaultUpdatedAt = airlineDescUpdatedAt.Default.(func() time.Time)
	// airline.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	airline.UpdateDefaultUpdatedAt = airlineDescUpdatedAt.UpdateDefault.(func() time.Time)
	// airlineDescID is the schema descriptor for id field.
	airlineDescID := airlineFields[0].Descriptor()
	// airline.DefaultID holds the default value on creation for the id field.
	airline.DefaultID = airlineDescID.Default.(func() uuid.UUID)
	airportFields := schema.Airport{}.Fields()
	_ = airportFields
	// airportDescName is the schema descriptor for name field.
	airportDescName := airportFields[1].Descriptor()
	// airport.NameValidator is a validator for the "name" field. It is called by the builders before save.
	airport.NameValidator = airportDescName.Validators[0].(func(string) error)
	// airportDescIataCode is the schema descriptor for iata_code field.
	airportDescIataCode := airportFields[2].Descriptor()
	// airport.IataCodeValidator is a validator for the "iata_code" field. It is called by the builders before save.
	airport.IataCodeValidator = airportDescIataCode.Validators[0].(func(string) error)
	// airportDescIcaoCode is the schema descriptor for icao_code field.
	airportDescIcaoCode := airportFields[3].Descriptor()
	// airport.IcaoCodeValidator is a validator for the "icao_code" field. It is called by the builders before save.
	airport.IcaoCodeValidator = airportDescIcaoCode.Validators[0].(func(string) error)
	// airportDescCreatedAt is the schema descriptor for created_at field.
	airportDescCreatedAt := airportFields[4].Descriptor()
	// airport.DefaultCreatedAt holds the default value on creation for the created_at field.
	airport.DefaultCreatedAt = airportDescCreatedAt.Default.(func() time.Time)
	// airportDescUpdatedAt is the schema descriptor for updated_at field.
	airportDescUpdatedAt := airportFields[5].Descriptor()
	// airport.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	airport.DefaultUpdatedAt = airportDescUpdatedAt.Default.(func() time.Time)
	// airport.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	airport.UpdateDefaultUpdatedAt = airportDescUpdatedAt.UpdateDefault.(func() time.Time)
	// airportDescID is the schema descriptor for id field.
	airportDescID := airportFields[0].Descriptor()
	// airport.DefaultID holds the default value on creation for the id field.
	airport.DefaultID = airportDescID.Default.(func() uuid.UUID)
	crewFields := schema.Crew{}.Fields()
	_ = crewFields
	// crewDescEmployeeID is the schema descriptor for employee_id field.
	crewDescEmployeeID := crewFields[1].Descriptor()
	// crew.EmployeeIDValidator is a validator for the "employee_id" field. It is called by the builders before save.
	crew.EmployeeIDValidator = crewDescEmployeeID.Validators[0].(func(string) error)
	// crewDescCreatedAt is the schema descriptor for created_at field.
	crewDescCreatedAt := crewFields[2].Descriptor()
	// crew.DefaultCreatedAt holds the default value on creation for the created_at field.
	crew.DefaultCreatedAt = crewDescCreatedAt.Default.(func() time.Time)
	// crewDescUpdatedAt is the schema descriptor for updated_at field.
	crewDescUpdatedAt := crewFields[3].Descriptor()
	// crew.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	crew.DefaultUpdatedAt = crewDescUpdatedAt.Default.(func() time.Time)
	// crew.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	crew.UpdateDefaultUpdatedAt = crewDescUpdatedAt.UpdateDefault.(func() time.Time)
	// crewDescID is the schema descriptor for id field.
	crewDescID := crewFields[0].Descriptor()
	// crew.DefaultID holds the default value on creation for the id field.
	crew.DefaultID = crewDescID.Default.(func() uuid.UUID)
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescFrequentFlyerNumber is the schema descriptor for frequent_flyer_number field.
	customerDescFrequentFlyerNumber := customerFields[1].Descriptor()
	// customer.FrequentFlyerNumberValidator is a validator for the "frequent_flyer_number" field. It is called by the builders before save.
	customer.FrequentFlyerNumberValidator = customerDescFrequentFlyerNumber.Validators[0].(func(string) error)
	// customerDescCreatedAt is the schema descriptor for created_at field.
	customerDescCreatedAt := customerFields[2].Descriptor()
	// customer.DefaultCreatedAt holds the default value on creation for the created_at field.
	customer.DefaultCreatedAt = customerDescCreatedAt.Default.(func() time.Time)
	// customerDescUpdatedAt is the schema descriptor for updated_at field.
	customerDescUpdatedAt := customerFields[3].Descriptor()
	// customer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customer.DefaultUpdatedAt = customerDescUpdatedAt.Default.(func() time.Time)
	// customer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customer.UpdateDefaultUpdatedAt = customerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerDescID is the schema descriptor for id field.
	customerDescID := customerFields[0].Descriptor()
	// customer.DefaultID holds the default value on creation for the id field.
	customer.DefaultID = customerDescID.Default.(func() uuid.UUID)
	flightFields := schema.Flight{}.Fields()
	_ = flightFields
	// flightDescFlightNumber is the schema descriptor for flight_number field.
	flightDescFlightNumber := flightFields[1].Descriptor()
	// flight.FlightNumberValidator is a validator for the "flight_number" field. It is called by the builders before save.
	flight.FlightNumberValidator = flightDescFlightNumber.Validators[0].(func(string) error)
	// flightDescDuration is the schema descriptor for duration field.
	flightDescDuration := flightFields[2].Descriptor()
	// flight.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	flight.DurationValidator = flightDescDuration.Validators[0].(func(int) error)
	// flightDescDistance is the schema descriptor for distance field.
	flightDescDistance := flightFields[3].Descriptor()
	// flight.DistanceValidator is a validator for the "distance" field. It is called by the builders before save.
	flight.DistanceValidator = flightDescDistance.Validators[0].(func(int) error)
	// flightDescCreatedAt is the schema descriptor for created_at field.
	flightDescCreatedAt := flightFields[5].Descriptor()
	// flight.DefaultCreatedAt holds the default value on creation for the created_at field.
	flight.DefaultCreatedAt = flightDescCreatedAt.Default.(func() time.Time)
	// flightDescUpdatedAt is the schema descriptor for updated_at field.
	flightDescUpdatedAt := flightFields[6].Descriptor()
	// flight.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flight.DefaultUpdatedAt = flightDescUpdatedAt.Default.(func() time.Time)
	// flight.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	flight.UpdateDefaultUpdatedAt = flightDescUpdatedAt.UpdateDefault.(func() time.Time)
	// flightDescID is the schema descriptor for id field.
	flightDescID := flightFields[0].Descriptor()
	// flight.DefaultID holds the default value on creation for the id field.
	flight.DefaultID = flightDescID.Default.(func() uuid.UUID)
	flightinstanceFields := schema.FlightInstance{}.Fields()
	_ = flightinstanceFields
	// flightinstanceDescDepartureGate is the schema descriptor for departure_gate field.
	flightinstanceDescDepartureGate := flightinstanceFields[1].Descriptor()
	// flightinstance.DepartureGateValidator is a validator for the "departure_gate" field. It is called by the builders before save.
	flightinstance.DepartureGateValidator = flightinstanceDescDepartureGate.Validators[0].(func(int) error)
	// flightinstanceDescArrivalGate is the schema descriptor for arrival_gate field.
	flightinstanceDescArrivalGate := flightinstanceFields[2].Descriptor()
	// flightinstance.ArrivalGateValidator is a validator for the "arrival_gate" field. It is called by the builders before save.
	flightinstance.ArrivalGateValidator = flightinstanceDescArrivalGate.Validators[0].(func(int) error)
	// flightinstanceDescCreatedAt is the schema descriptor for created_at field.
	flightinstanceDescCreatedAt := flightinstanceFields[4].Descriptor()
	// flightinstance.DefaultCreatedAt holds the default value on creation for the created_at field.
	flightinstance.DefaultCreatedAt = flightinstanceDescCreatedAt.Default.(func() time.Time)
	// flightinstanceDescUpdatedAt is the schema descriptor for updated_at field.
	flightinstanceDescUpdatedAt := flightinstanceFields[5].Descriptor()
	// flightinstance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flightinstance.DefaultUpdatedAt = flightinstanceDescUpdatedAt.Default.(func() time.Time)
	// flightinstance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	flightinstance.UpdateDefaultUpdatedAt = flightinstanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// flightinstanceDescID is the schema descriptor for id field.
	flightinstanceDescID := flightinstanceFields[0].Descriptor()
	// flightinstance.DefaultID holds the default value on creation for the id field.
	flightinstance.DefaultID = flightinstanceDescID.Default.(func() uuid.UUID)
	flightreservationFields := schema.FlightReservation{}.Fields()
	_ = flightreservationFields
	// flightreservationDescReservationNumber is the schema descriptor for reservation_number field.
	flightreservationDescReservationNumber := flightreservationFields[1].Descriptor()
	// flightreservation.ReservationNumberValidator is a validator for the "reservation_number" field. It is called by the builders before save.
	flightreservation.ReservationNumberValidator = flightreservationDescReservationNumber.Validators[0].(func(string) error)
	// flightreservationDescCreatedAt is the schema descriptor for created_at field.
	flightreservationDescCreatedAt := flightreservationFields[3].Descriptor()
	// flightreservation.DefaultCreatedAt holds the default value on creation for the created_at field.
	flightreservation.DefaultCreatedAt = flightreservationDescCreatedAt.Default.(func() time.Time)
	// flightreservationDescUpdatedAt is the schema descriptor for updated_at field.
	flightreservationDescUpdatedAt := flightreservationFields[4].Descriptor()
	// flightreservation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flightreservation.DefaultUpdatedAt = flightreservationDescUpdatedAt.Default.(func() time.Time)
	// flightreservation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	flightreservation.UpdateDefaultUpdatedAt = flightreservationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// flightreservationDescID is the schema descriptor for id field.
	flightreservationDescID := flightreservationFields[0].Descriptor()
	// flightreservation.DefaultID holds the default value on creation for the id field.
	flightreservation.DefaultID = flightreservationDescID.Default.(func() uuid.UUID)
	flightscheduleFields := schema.FlightSchedule{}.Fields()
	_ = flightscheduleFields
	// flightscheduleDescCreatedAt is the schema descriptor for created_at field.
	flightscheduleDescCreatedAt := flightscheduleFields[6].Descriptor()
	// flightschedule.DefaultCreatedAt holds the default value on creation for the created_at field.
	flightschedule.DefaultCreatedAt = flightscheduleDescCreatedAt.Default.(func() time.Time)
	// flightscheduleDescUpdatedAt is the schema descriptor for updated_at field.
	flightscheduleDescUpdatedAt := flightscheduleFields[7].Descriptor()
	// flightschedule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flightschedule.DefaultUpdatedAt = flightscheduleDescUpdatedAt.Default.(func() time.Time)
	// flightschedule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	flightschedule.UpdateDefaultUpdatedAt = flightscheduleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// flightscheduleDescID is the schema descriptor for id field.
	flightscheduleDescID := flightscheduleFields[0].Descriptor()
	// flightschedule.DefaultID holds the default value on creation for the id field.
	flightschedule.DefaultID = flightscheduleDescID.Default.(func() uuid.UUID)
	flightseatFields := schema.FlightSeat{}.Fields()
	_ = flightseatFields
	// flightseatDescFare is the schema descriptor for fare field.
	flightseatDescFare := flightseatFields[1].Descriptor()
	// flightseat.FareValidator is a validator for the "fare" field. It is called by the builders before save.
	flightseat.FareValidator = flightseatDescFare.Validators[0].(func(float64) error)
	// flightseatDescCreatedAt is the schema descriptor for created_at field.
	flightseatDescCreatedAt := flightseatFields[2].Descriptor()
	// flightseat.DefaultCreatedAt holds the default value on creation for the created_at field.
	flightseat.DefaultCreatedAt = flightseatDescCreatedAt.Default.(func() time.Time)
	// flightseatDescUpdatedAt is the schema descriptor for updated_at field.
	flightseatDescUpdatedAt := flightseatFields[3].Descriptor()
	// flightseat.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	flightseat.DefaultUpdatedAt = flightseatDescUpdatedAt.Default.(func() time.Time)
	// flightseat.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	flightseat.UpdateDefaultUpdatedAt = flightseatDescUpdatedAt.UpdateDefault.(func() time.Time)
	// flightseatDescID is the schema descriptor for id field.
	flightseatDescID := flightseatFields[0].Descriptor()
	// flightseat.DefaultID holds the default value on creation for the id field.
	flightseat.DefaultID = flightseatDescID.Default.(func() uuid.UUID)
	frontdeskFields := schema.FrontDesk{}.Fields()
	_ = frontdeskFields
	// frontdeskDescEmployeeID is the schema descriptor for employee_id field.
	frontdeskDescEmployeeID := frontdeskFields[1].Descriptor()
	// frontdesk.EmployeeIDValidator is a validator for the "employee_id" field. It is called by the builders before save.
	frontdesk.EmployeeIDValidator = frontdeskDescEmployeeID.Validators[0].(func(string) error)
	// frontdeskDescCreatedAt is the schema descriptor for created_at field.
	frontdeskDescCreatedAt := frontdeskFields[2].Descriptor()
	// frontdesk.DefaultCreatedAt holds the default value on creation for the created_at field.
	frontdesk.DefaultCreatedAt = frontdeskDescCreatedAt.Default.(func() time.Time)
	// frontdeskDescUpdatedAt is the schema descriptor for updated_at field.
	frontdeskDescUpdatedAt := frontdeskFields[3].Descriptor()
	// frontdesk.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	frontdesk.DefaultUpdatedAt = frontdeskDescUpdatedAt.Default.(func() time.Time)
	// frontdesk.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	frontdesk.UpdateDefaultUpdatedAt = frontdeskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// frontdeskDescID is the schema descriptor for id field.
	frontdeskDescID := frontdeskFields[0].Descriptor()
	// frontdesk.DefaultID holds the default value on creation for the id field.
	frontdesk.DefaultID = frontdeskDescID.Default.(func() uuid.UUID)
	iteneraryFields := schema.Itenerary{}.Fields()
	_ = iteneraryFields
	// iteneraryDescCreatedAt is the schema descriptor for created_at field.
	iteneraryDescCreatedAt := iteneraryFields[1].Descriptor()
	// itenerary.DefaultCreatedAt holds the default value on creation for the created_at field.
	itenerary.DefaultCreatedAt = iteneraryDescCreatedAt.Default.(func() time.Time)
	// iteneraryDescUpdatedAt is the schema descriptor for updated_at field.
	iteneraryDescUpdatedAt := iteneraryFields[2].Descriptor()
	// itenerary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	itenerary.DefaultUpdatedAt = iteneraryDescUpdatedAt.Default.(func() time.Time)
	// itenerary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	itenerary.UpdateDefaultUpdatedAt = iteneraryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// iteneraryDescID is the schema descriptor for id field.
	iteneraryDescID := iteneraryFields[0].Descriptor()
	// itenerary.DefaultID holds the default value on creation for the id field.
	itenerary.DefaultID = iteneraryDescID.Default.(func() uuid.UUID)
	passengerFields := schema.Passenger{}.Fields()
	_ = passengerFields
	// passengerDescFirstname is the schema descriptor for firstname field.
	passengerDescFirstname := passengerFields[1].Descriptor()
	// passenger.FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	passenger.FirstnameValidator = passengerDescFirstname.Validators[0].(func(string) error)
	// passengerDescLastname is the schema descriptor for lastname field.
	passengerDescLastname := passengerFields[2].Descriptor()
	// passenger.LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	passenger.LastnameValidator = passengerDescLastname.Validators[0].(func(string) error)
	// passengerDescAge is the schema descriptor for age field.
	passengerDescAge := passengerFields[3].Descriptor()
	// passenger.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	passenger.AgeValidator = passengerDescAge.Validators[0].(func(int) error)
	// passengerDescPassportNumber is the schema descriptor for passport_number field.
	passengerDescPassportNumber := passengerFields[4].Descriptor()
	// passenger.PassportNumberValidator is a validator for the "passport_number" field. It is called by the builders before save.
	passenger.PassportNumberValidator = passengerDescPassportNumber.Validators[0].(func(string) error)
	// passengerDescCreatedAt is the schema descriptor for created_at field.
	passengerDescCreatedAt := passengerFields[5].Descriptor()
	// passenger.DefaultCreatedAt holds the default value on creation for the created_at field.
	passenger.DefaultCreatedAt = passengerDescCreatedAt.Default.(func() time.Time)
	// passengerDescUpdatedAt is the schema descriptor for updated_at field.
	passengerDescUpdatedAt := passengerFields[6].Descriptor()
	// passenger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	passenger.DefaultUpdatedAt = passengerDescUpdatedAt.Default.(func() time.Time)
	// passenger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	passenger.UpdateDefaultUpdatedAt = passengerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// passengerDescID is the schema descriptor for id field.
	passengerDescID := passengerFields[0].Descriptor()
	// passenger.DefaultID holds the default value on creation for the id field.
	passenger.DefaultID = passengerDescID.Default.(func() uuid.UUID)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescPermission is the schema descriptor for permission field.
	permissionDescPermission := permissionFields[1].Descriptor()
	// permission.PermissionValidator is a validator for the "permission" field. It is called by the builders before save.
	permission.PermissionValidator = permissionDescPermission.Validators[0].(func(string) error)
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionFields[2].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionFields[3].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionFields[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() uuid.UUID)
	pilotFields := schema.Pilot{}.Fields()
	_ = pilotFields
	// pilotDescEmployeeID is the schema descriptor for employee_id field.
	pilotDescEmployeeID := pilotFields[1].Descriptor()
	// pilot.EmployeeIDValidator is a validator for the "employee_id" field. It is called by the builders before save.
	pilot.EmployeeIDValidator = pilotDescEmployeeID.Validators[0].(func(string) error)
	// pilotDescLicenceNumber is the schema descriptor for licence_number field.
	pilotDescLicenceNumber := pilotFields[2].Descriptor()
	// pilot.LicenceNumberValidator is a validator for the "licence_number" field. It is called by the builders before save.
	pilot.LicenceNumberValidator = pilotDescLicenceNumber.Validators[0].(func(string) error)
	// pilotDescFlightHours is the schema descriptor for flight_hours field.
	pilotDescFlightHours := pilotFields[3].Descriptor()
	// pilot.DefaultFlightHours holds the default value on creation for the flight_hours field.
	pilot.DefaultFlightHours = pilotDescFlightHours.Default.(int)
	// pilot.FlightHoursValidator is a validator for the "flight_hours" field. It is called by the builders before save.
	pilot.FlightHoursValidator = pilotDescFlightHours.Validators[0].(func(int) error)
	// pilotDescCreatedAt is the schema descriptor for created_at field.
	pilotDescCreatedAt := pilotFields[4].Descriptor()
	// pilot.DefaultCreatedAt holds the default value on creation for the created_at field.
	pilot.DefaultCreatedAt = pilotDescCreatedAt.Default.(func() time.Time)
	// pilotDescUpdatedAt is the schema descriptor for updated_at field.
	pilotDescUpdatedAt := pilotFields[5].Descriptor()
	// pilot.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pilot.DefaultUpdatedAt = pilotDescUpdatedAt.Default.(func() time.Time)
	// pilot.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pilot.UpdateDefaultUpdatedAt = pilotDescUpdatedAt.UpdateDefault.(func() time.Time)
	// pilotDescID is the schema descriptor for id field.
	pilotDescID := pilotFields[0].Descriptor()
	// pilot.DefaultID holds the default value on creation for the id field.
	pilot.DefaultID = pilotDescID.Default.(func() uuid.UUID)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[2].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[3].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() uuid.UUID)
	seatFields := schema.Seat{}.Fields()
	_ = seatFields
	// seatDescSeatNumber is the schema descriptor for seat_number field.
	seatDescSeatNumber := seatFields[1].Descriptor()
	// seat.SeatNumberValidator is a validator for the "seat_number" field. It is called by the builders before save.
	seat.SeatNumberValidator = seatDescSeatNumber.Validators[0].(func(int) error)
	// seatDescSeatRow is the schema descriptor for seat_row field.
	seatDescSeatRow := seatFields[2].Descriptor()
	// seat.SeatRowValidator is a validator for the "seat_row" field. It is called by the builders before save.
	seat.SeatRowValidator = seatDescSeatRow.Validators[0].(func(string) error)
	// seatDescCreatedAt is the schema descriptor for created_at field.
	seatDescCreatedAt := seatFields[5].Descriptor()
	// seat.DefaultCreatedAt holds the default value on creation for the created_at field.
	seat.DefaultCreatedAt = seatDescCreatedAt.Default.(func() time.Time)
	// seatDescUpdatedAt is the schema descriptor for updated_at field.
	seatDescUpdatedAt := seatFields[6].Descriptor()
	// seat.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	seat.DefaultUpdatedAt = seatDescUpdatedAt.Default.(func() time.Time)
	// seat.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	seat.UpdateDefaultUpdatedAt = seatDescUpdatedAt.UpdateDefault.(func() time.Time)
	// seatDescID is the schema descriptor for id field.
	seatDescID := seatFields[0].Descriptor()
	// seat.DefaultID holds the default value on creation for the id field.
	seat.DefaultID = seatDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstname is the schema descriptor for firstname field.
	userDescFirstname := userFields[1].Descriptor()
	// user.FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	user.FirstnameValidator = userDescFirstname.Validators[0].(func(string) error)
	// userDescLastname is the schema descriptor for lastname field.
	userDescLastname := userFields[2].Descriptor()
	// user.LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	user.LastnameValidator = userDescLastname.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[4].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
