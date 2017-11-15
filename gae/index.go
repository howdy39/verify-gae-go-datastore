package gae

import (
	"context"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
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

func GoonHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	g := goon.FromContext(ctx)

	employee := &Employee{
		Ymd:       "2014-12-24",
		FirstName: "Antonio Goon",
		LastName:  "Salieri Goon",
		HireDate:  time.Now(),
	}

	if _, err := g.Put(employee); err != nil {
		serveError(ctx, err)
	}

	e2 := &Employee{
		Ymd: employee.Ymd,
	}

	if err := g.Get(e2); err != nil {
		serveError(ctx, err)
	}

	log.Debugf(ctx, "%s,%s,%v", e2.FirstName, e2.LastName, e2.HireDate)
}

func GoongetHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	g := goon.FromContext(ctx)

	e2 := &Employee{
		Ymd: "2014-12-24",
	}

	if err := g.Get(e2); err != nil {
		serveError(ctx, err)
	}

	log.Debugf(ctx, "%s,%s,%v", e2.FirstName, e2.LastName, e2.HireDate)
}

func MemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	item := &memcache.Item{
		Key:   "myKey",
		Value: []byte("FirstMemcache"),
	}
	// [3]
	if err := memcache.Add(ctx, item); err == memcache.ErrNotStored {
		// 既に存在するKeyの場合の処理
		log.Debugf(ctx, "Key already exist. key is ", item.Key)
		return
	}
}

func serveError(ctx context.Context, err error) {
	log.Errorf(ctx, "%v", err)
}

type Employee struct {
	Ymd string `datastore:"-" goon:"id"`

	FirstName string    `datastore:firstName,noindex`
	LastName  string    `datastore:lastName,noindex`
	HireDate  time.Time `datastore:HireDate,noindex`
}
