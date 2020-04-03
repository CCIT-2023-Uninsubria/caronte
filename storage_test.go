package main

import (
	"testing"
)

const testCollection = "characters"

func testInsert(t *testing.T) {
	// insert a document in an invalid connection
	insertedId, err := storage.InsertOne(testContext, "invalid_collection",
		OrderedDocument{{"key", "invalid"}})
	if insertedId != nil || err == nil {
		t.Fatal("inserting documents in invalid collections must fail")
	}

	// insert ordered document
	beatriceId, err := storage.InsertOne(testContext, testCollection,
		OrderedDocument{{"name", "Beatrice"}, {"description", "blablabla"}})
	if err != nil {
		t.Fatal(err)
	}
	if beatriceId == nil {
		t.Fatal("failed to insert an ordered document")
	}

	// insert unordered document
	virgilioId, err := storage.InsertOne(testContext, testCollection,
		UnorderedDocument{"name": "Virgilio", "description": "blablabla"})
	if err != nil {
		t.Fatal(err)
	}
	if virgilioId == nil {
		t.Fatal("failed to insert an unordered document")
	}

	// insert document with custom id
	danteId := "000000"
	insertedId, err = storage.InsertOne(testContext, testCollection,
		UnorderedDocument{"_id": danteId, "name": "Dante Alighieri", "description": "blablabla"})
	if err != nil {
		t.Fatal(err)
	}
	if insertedId != danteId {
		t.Fatal("returned id doesn't match")
	}

	// insert duplicate document
	insertedId, err = storage.InsertOne(testContext, testCollection,
		UnorderedDocument{"_id": danteId, "name": "Dante Alighieri", "description": "blablabla"})
	if insertedId != nil || err == nil {
		t.Fatal("inserting duplicate id must fail")
	}
}

func testFindOne(t *testing.T) {
	// find a document in an invalid connection
	result, err := storage.FindOne(testContext, "invalid_collection",
		OrderedDocument{{"key", "invalid"}})
	if result != nil || err == nil {
		t.Fatal("find a document in an invalid collections must fail")
	}

	// find an existing document
	result, err = storage.FindOne(testContext, testCollection, OrderedDocument{{"_id", "000000"}})
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("FindOne cannot find the valid document")
	}
	name, ok := result["name"]
	if !ok || name != "Dante Alighieri" {
		t.Fatal("document retrieved with FindOne is invalid")
	}

	// find an existing document
	result, err = storage.FindOne(testContext, testCollection, OrderedDocument{{"_id", "invalid_id"}})
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Fatal("FindOne cannot find an invalid document")
	}
}

func TestBasicOperations(t *testing.T) {
	t.Run("testInsert", testInsert)
	t.Run("testFindOne", testFindOne)
}