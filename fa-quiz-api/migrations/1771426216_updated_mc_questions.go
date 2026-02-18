package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2001502728")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2272205672",
			"hidden": false,
			"id": "relation4186027310",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "chapter",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2001502728")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation4186027310")

		return app.Save(collection)
	})
}
