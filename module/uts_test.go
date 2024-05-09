package package_uts_test

import (
	"fmt"
	"testing"

	"github.com/gryzlegrizz/package_uts/module"
)

func TestInsertMuseumCollection(t *testing.T) {
	items := []interface{}{
		map[string]interface{}{
			"name":        "Painting 1",
			"description": "A beautiful painting",
			"year":        2000,
			"artist":      "Artist 1",
			"medium":      "Oil on canvas",
			"dimensions":  "50x70 cm",
			"origin":      "Country 1",
			"acquisition": "Purchase",
			"condition":   "Good",
		},
		map[string]interface{}{
			"name":        "Sculpture 1",
			"description": "An amazing sculpture",
			"year":        1995,
			"artist":      "Artist 2",
			"medium":      "Marble",
			"dimensions":  "30x30x50 cm",
			"origin":      "Country 2",
			"acquisition": "Donation",
			"condition":   "Excellent",
		},
	}

	insertedID := package_uts.InsertMuseumCollection(
		"Example Museum",
		"A wonderful museum",
		items,
		"City X",
		"9:00 AM - 5:00 PM",
		1990,
		"Director X",
		"http://example.com",
		1000,
	)
	fmt.Println("Inserted Museum Collection ID:", insertedID)
}

func TestGetAllMuseumCollections(t *testing.T) {
	collections := package_uts.GetAllMuseumCollections()
	fmt.Println("All Museum Collections:", collections)
}

func TestGetAllMuseumItems(t *testing.T) {
	items := package_uts.GetAllMuseumItems()
	fmt.Println("All Museum Items:", items)
}
