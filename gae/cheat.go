package gae

import (
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
)

type Parent struct {
	ParentID  string    `datastore:"-" goon:"id"`
	Name      string    `datastore:"Name,noindex"`
	Age       string    `datastore:"Age,noindex"`
	CreatedAt time.Time `datastore:"CreatedAt,noindex"`
}
type Child struct {
	Parent    *datastore.Key `datastore:"-" goon:"parent"`
	ChildID   string         `datastore:"-" goon:"id"`
	Name      string         `datastore:"Name,noindex"`
	CreatedAt time.Time      `datastore:"CreatedAt,noindex"`
}

func CheatSaveHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := goon.FromContext(c)

	parent := &Parent{
		ParentID:  "20171127",
		Name:      "ParentName1",
		Age:       "20",
		CreatedAt: time.Now(),
	}

	pkey, err := g.Put(parent)
	if err != nil {
		log.Errorf(c, "g.Put(parent):%v", err)
		return
	}

	child := &Child{
		Parent:    pkey,
		ChildID:   parent.ParentID + "-ABC",
		Name:      "ChildName1",
		CreatedAt: time.Now(),
	}

	_, err = g.Put(child)
	if err != nil {
		log.Errorf(c, "g.Put(child):%v", err)
		return
	}

}
