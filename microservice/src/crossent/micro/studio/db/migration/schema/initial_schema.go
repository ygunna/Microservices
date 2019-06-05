package schema

import "crossent/micro/studio/db/migration"

func InitialSchema(tx migration.LimitedTx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_app (
		        id serial PRIMARY KEY,
			name text NOT NULL,
			org_guid varchar(255) NOT NULL,
			space_guid varchar(255) NOT NULL,
			version varchar(255) NOT NULL,
			description text,
			visible varchar(255) NOT NULL DEFAULT 'private',
			status varchar(255),
			url varchar(255),
			swagger text,
			user_id varchar(255) NOT NULL,
			active varchar(1) DEFAULT 'Y' NOT NULL,
			UNIQUE (name, space_guid),
			UNIQUE (name, space_guid, version)
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_app_app (
		        id serial PRIMARY KEY,
		        micro_id integer REFERENCES micro_app (id),
		        app_guid varchar(255) NOT NULL,
		        source_guid varchar(255) NOT NULL,
		        essential varchar(10)
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_app_service (
		        id serial PRIMARY KEY,
		        micro_id integer REFERENCES micro_app (id),
		        service_guid varchar(255) NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
