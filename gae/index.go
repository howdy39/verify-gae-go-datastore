package gae

import (
	"context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	employee := &Employee{
		FirstName: "Antonio",
		LastName:  "Salieri",
		HireDate:  time.Now(),
	}
	employee.AttendedHRTraining = true

	key := datastore.NewIncompleteKey(ctx, "Employee", nil)
	employeeKey, err := datastore.Put(ctx, key, employee)
	if err != nil {
		serveError(ctx, err)
	}

	var e2 Employee
	if err := datastore.Get(ctx, employeeKey, &e2); err != nil {
		serveError(ctx, err)
	}

	log.Debugf(ctx, "%s,%s,%v", e2.FirstName, e2.LastName, e2.HireDate)
}

func serveError(ctx context.Context, err error) {
	log.Errorf(ctx, "%v", err)
}

type Employee struct {
	FirstName          string
	LastName           string
	HireDate           time.Time
	AttendedHRTraining bool
}
