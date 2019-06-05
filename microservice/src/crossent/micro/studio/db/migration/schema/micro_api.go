package schema

import "crossent/micro/studio/db/migration"

func MicroApI(tx migration.LimitedTx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_api (
		        id serial PRIMARY KEY,
		        part varchar(10) DEFAULT '',
		        org_guid varchar(255) NOT NULL,
		        name varchar(255) NOT NULL,
		        host varchar(255) DEFAULT '',
		        path varchar(255) NOT NULL,
		        version varchar(255) DEFAULT '',
		        rest_api text DEFAULT '',
		        active varchar(1) DEFAULT 'Y' NOT NULL,
		        description varchar(255) DEFAULT '',
		        image text DEFAULT '200,200,200',
		        user_id varchar(255) NOT NULL,
			updated timestamp without time zone NOT NULL,
			UNIQUE (name)
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_api_rule (
		        id serial PRIMARY KEY,
		        rule text DEFAULT '' NOT NULL,
		        active varchar(1) DEFAULT 'Y' NOT NULL
		)
	`)
	if err != nil {
		return err
	}


	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_api_frontend (
		        id serial PRIMARY KEY,
		        api_id integer REFERENCES micro_api (id),
		        micro_id integer REFERENCES micro_app (id)
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS micro_app_api (
		        id serial PRIMARY KEY,
		        micro_id integer REFERENCES micro_app (id),
		        api_id integer REFERENCES micro_api (id),
		        username varchar(255) NOT NULL DEFAULT '',
		        active varchar(1) DEFAULT 'Y' NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
