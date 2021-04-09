package main

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateDuplicate(t *testing.T) {
	uuid := uuid.NewString()

	prod := NewProduct(db)
	err = prod.create(uuid)
	if err != nil {
		t.Fatalf("Initial create new item failed")
	}

	prod1 := NewProduct(db)
	err = prod1.create(uuid)
	if err == nil {
		t.Fatalf("Duplicate item created")
	}

	/*
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello("Gladys")
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	*/
}
func TestCreate(t *testing.T) {
	uuid := uuid.NewString()

	prod := NewProduct(db)
	err = prod.create(uuid)
	if err != nil {
		t.Fatalf("Create new item failed")
	}
}
