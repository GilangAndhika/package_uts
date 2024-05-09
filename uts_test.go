package package_uts_test

import (
	"fmt"
	"testing"

	"github.com/gryzlegrizz/package_uts" 
)

func TestInsertMuseumCollection(t *testing.T) {
	name := "Museum Seni Modern"
	description := "Museum ini menyimpan karya-karya seni modern terkemuka dari berbagai seniman internasional."
	items := []package_uts.MuseumItem{
		{
			Name:        "Pemandangan Senja",
			Description: "Sebuah lukisan pemandangan senja yang indah.",
			Year:        1995,
		},
		{
			Name:        "Patung Abstrak",
			Description: "Patung abstrak yang menggambarkan perjuangan manusia.",
			Year:        2000,
		},
	}
	location := "Jakarta, Indonesia"
	openingHours := "Senin - Sabtu: 09.00 - 17.00"
	establishedYear := 1980
	director := "Jane Doe"
	website := "https://museumsenimodern.com"
	visitorCount := 10000

	insertedID := package_uts.InsertMuseumCollection(name, description, items, location, openingHours, establishedYear, director, website, visitorCount)

	if insertedID == nil {
		t.Error("Failed to insert museum collection.")
	}
	fmt.Println("Inserted Museum Collection ID:", insertedID)
}

func TestInsertMuseumItem(t *testing.T) {
	name := "Lukisan Mona Lisa"
	description := "Lukisan terkenal karya Leonardo da Vinci yang menggambarkan seorang wanita dengan senyuman misterius."
	year := 1503
	artist := "Leonardo da Vinci"
	medium := "Cat minyak di atas kayu"
	dimensions := "77 cm × 53 cm (30 in × 21 in)"
	origin := "Italia"
	acquisition := "Pembelian oleh Raja Francis I dari Prancis"
	condition := "Baik"

	insertedID := package_uts.InsertMuseumItem(name, description, year, artist, medium, dimensions, origin, acquisition, condition)

	if insertedID == nil {
		t.Error("Failed to insert museum item.")
	}
	fmt.Println("Inserted Museum Item ID:", insertedID)
}
