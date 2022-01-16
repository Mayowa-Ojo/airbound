package flights

import (
	"airbound/internal/ent"
)

type Datastore struct {
	c *ent.Client
}

func NewFlightRepository(client *ent.Client) FlightRepository {
	return &Datastore{c: client}
}

func (d *Datastore) FindListFlightInstances() {

}
