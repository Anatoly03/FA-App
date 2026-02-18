package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2272205672")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1345586850",
			"hidden": false,
			"id": "relation1061296345",
			"maxSelect": 999,
			"minSelect": 0,
			"name": "lessons",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2272205672")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1061296345")

		return app.Save(collection)
	})
}
