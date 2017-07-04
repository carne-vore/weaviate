/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	gouuid "github.com/satori/go.uuid"

	"math"
	"sort"

	"github.com/hashicorp/go-memdb"
	"github.com/weaviate/weaviate/connectors/utils"
)

// Datastore has some basic variables.
type Memory struct {
	client *memdb.MemDB
	kind   string
}

// SetConfig is used to fill in a struct with config variables
func (f *Memory) GetName() string {
	return "memory"
}

// GetName returns a unique connector name
func (f *Memory) SetConfig(interface{}) {
	// NOTHING
}

// Creates connection and tables if not already available (which is never because it is in memory)
func (f *Memory) Connect() error {

	// Create the weaviate DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			// create `weaviate` DB
			"weaviate": &memdb.TableSchema{
				Name: "weaviate",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
					"Deleted": &memdb.IndexSchema{
						Name:    "Deleted",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Deleted"},
					},
					"CreateTimeMs": &memdb.IndexSchema{
						Name:    "CreateTimeMs",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CreateTimeMs"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Owner": &memdb.IndexSchema{
						Name:    "Owner",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Owner"},
					},
					"RefType": &memdb.IndexSchema{
						Name:    "RefType",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RefType"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
			// create `weaviate` DB
			"weaviate_history": &memdb.TableSchema{
				Name: "weaviate_history",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},

					"Deleted": &memdb.IndexSchema{
						Name:    "Deleted",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Deleted"},
					},
					"CreateTimeMs": &memdb.IndexSchema{
						Name:    "CreateTimeMs",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CreateTimeMs"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Owner": &memdb.IndexSchema{
						Name:    "Owner",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Owner"},
					},
					"RefType": &memdb.IndexSchema{
						Name:    "RefType",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "RefType"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
			// create `weaviate` DB
			"weaviate_users": &memdb.TableSchema{
				Name: "weaviate_users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
					"KeyExpiresMs": &memdb.IndexSchema{
						Name:    "KeyExpiresMs",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "KeyExpiresMs"},
					},
					"KeyToken": &memdb.IndexSchema{
						Name:    "KeyToken",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "KeyToken"},
					},
					"Object": &memdb.IndexSchema{
						Name:    "Object",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Object"},
					},
					"Parent": &memdb.IndexSchema{
						Name:    "Parent",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Parent"},
					},
					"Uuid": &memdb.IndexSchema{
						Name:    "Uuid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Uuid"},
					},
				},
			},
		},
	}

	// Create a new data base
	client, err := memdb.NewMemDB(schema)

	// If error, return it. Otherwise set client.
	if err != nil {
		return err
	}

	f.client = client

	log.Println("INFO: In memory database is used for testing / development purposes only")

	return nil

}

// Creates a root key, normally this should be validaded, but because it is an inmemory DB it is created always
func (f *Memory) Init() error {
	dbObject := connector_utils.DatabaseUsersObject{}

	// Create key token
	dbObject.KeyToken = fmt.Sprintf("%v", gouuid.NewV4())

	// Uuid + name
	uuid := fmt.Sprintf("%v", gouuid.NewV4())

	// Auto set the parent ID to root *
	dbObject.Parent = "*"

	// Set Uuid
	dbObject.Uuid = uuid

	// Set chmod variables
	dbObjectObject := connector_utils.DatabaseUsersObjectsObject{}
	dbObjectObject.Read = true
	dbObjectObject.Write = true
	dbObjectObject.Delete = true

	// Get ips as v6
	var ips []string
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipv6 := ip.To16()
			ips = append(ips, ipv6.String())
		}
	}

	dbObjectObject.IpOrigin = ips

	// Marshall and add to object
	dbObjectObjectJSON, _ := json.Marshal(dbObjectObject)
	dbObject.Object = string(dbObjectObjectJSON)

	// Create a write transaction
	txn := f.client.Txn(true)

	// Saves the new entity.
	if err := txn.Insert("weaviate_users", dbObject); err != nil {
		return err
	}

	// commit transaction
	txn.Commit()

	// Print the key
	log.Println("INFO: A new root key is created. More info: https://github.com/weaviate/weaviate/blob/develop/README.md#authentication")
	log.Println("INFO: Auto set allowed IPs to: ", ips)
	log.Println("ROOTKEY=" + dbObject.KeyToken)

	return nil
}

func (f *Memory) Add(dbObject connector_utils.DatabaseObject) (string, error) {

	// Create a write transaction
	txn := f.client.Txn(true)

	// Saves the new entity.
	if err := txn.Insert("weaviate", dbObject); err != nil {
		return "Error", err
	}

	// commit transaction
	txn.Commit()

	// Return the ID that is used to create.
	return dbObject.Uuid, nil

}

func (f *Memory) Get(Uuid string) (connector_utils.DatabaseObject, error) {

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.First("weaviate", "Uuid", Uuid)
	if err != nil {
		return connector_utils.DatabaseObject{}, err
	}

	// Return 'not found'
	if result == nil {
		notFoundErr := errors.New("no object with such UUID found")
		return connector_utils.DatabaseObject{}, notFoundErr
	}

	// Return found object
	return result.(connector_utils.DatabaseObject), nil

}

// return a list
func (f *Memory) List(refType string, limit int, page int, referenceFilter *connector_utils.ObjectReferences) (connector_utils.DatabaseObjects, int64, error) {
	dataObjs := connector_utils.DatabaseObjects{}

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.Get("weaviate", "id")

	// return the error
	if err != nil {
		return dataObjs, 0, err
	}

	if result != nil {

		// loop through the results
		loopResults := true
		for loopResults == true {
			singleResult := result.Next()
			if singleResult == nil {
				// no results left, stop the loop
				loopResults = false
			} else {
				// only store if refType is correct
				if singleResult.(connector_utils.DatabaseObject).RefType == refType &&
					!singleResult.(connector_utils.DatabaseObject).Deleted {

					if referenceFilter != nil {
						// check for extra filters
						if referenceFilter.ThingID != "" &&
							singleResult.(connector_utils.DatabaseObject).RelatedObjects.ThingID == referenceFilter.ThingID {
							dataObjs = append(dataObjs, singleResult.(connector_utils.DatabaseObject))
						}
					} else {
						dataObjs = append(dataObjs, singleResult.(connector_utils.DatabaseObject))
					}
				}
			}
		}

		// Sorting on CreateTimeMs
		sort.Sort(dataObjs)

		// count total
		totalResults := len(dataObjs)

		// calculate the amount to chop off totalResults-limit
		offset := (limit * (page - 1))
		end := int(math.Min(float64(limit*(page)), float64(totalResults)))
		dataObjs := dataObjs[offset:end]

		// return found set
		return dataObjs, int64(totalResults), err
	}

	// nothing found
	return dataObjs, 0, nil
}

// Validate if a user has access, returns permissions object
func (f *Memory) ValidateKey(token string) ([]connector_utils.DatabaseUsersObject, error) {

	dbUsersObjects := []connector_utils.DatabaseUsersObject{}

	// Create read-only transaction
	txn := f.client.Txn(false)
	defer txn.Abort()

	// Lookup by Uuid
	result, err := txn.First("weaviate_users", "KeyToken", token)
	if err != nil || result == nil {
		return []connector_utils.DatabaseUsersObject{}, err
	}

	// add to results
	dbUsersObjects = append(dbUsersObjects, result.(connector_utils.DatabaseUsersObject))

	// keys are found, return true
	return dbUsersObjects, nil
}

// AddUser to DB
func (f *Memory) AddKey(parentUuid string, dbObject connector_utils.DatabaseUsersObject) (connector_utils.DatabaseUsersObject, error) {

	// Create a write transaction
	txn := f.client.Txn(true)

	// Create key token
	dbObject.KeyToken = fmt.Sprintf("%v", gouuid.NewV4())

	// Auto set the parent ID
	dbObject.Parent = parentUuid

	// Saves the new entity.
	if err := txn.Insert("weaviate_users", dbObject); err != nil {
		return dbObject, err
	}

	// commit transaction
	txn.Commit()

	// Return the ID that is used to create.
	return dbObject, nil

}