package data_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stuartforrest-infinity/websocket-lambda/data"
)

func TestConnections_PutConnectionID(t *testing.T) {
	sess, table, drop, err := CreateTestTable()
	assert.NoError(t, err, "create table")
	defer drop()
	datastore := data.NewStore(table, sess)

	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// Create a connection identifier.
	ci := data.ConnectionItemData{
		ID: "test-id",
	}
	// No connections should be returned
	r1, err := datastore.GetAllConnectionIDs(now, now.Add(-(time.Minute * 2)))
	assert.NoError(t, err, "get all connections - should be no error")
	assert.Empty(t, r1)
	err = datastore.PutConnectionID(ci, now.Add(-(time.Minute * 1)))
	assert.NoError(t, err, "insert connection")
	// Get connections that have been inserted
	r2, err := datastore.GetAllConnectionIDs(now, now.Add(-(time.Minute * 2)))
	assert.NoError(t, err, "get all connections - should be no error")
	assert.Equal(t, r2, []data.ConnectionItemData{
		{
			ID: "test-id",
		},
	}, "should return connection with ID: test-id")
	// Only get connections inserted within the last 2 minutes
	err = datastore.PutConnectionID(ci, now.Add(-(time.Minute * 3)))
	// Get connections that have been inserted
	r3, err := datastore.GetAllConnectionIDs(now, now.Add(-(time.Minute * 2)))
	assert.NoError(t, err, "get all connections - should be no error")
	assert.Equal(t, r3, []data.ConnectionItemData{
		{
			ID: "test-id",
		},
	}, "should return connection with ID: test-id")
}
