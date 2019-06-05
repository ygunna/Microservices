package schema

import "crossent/micro/studio/db/migration"

func MicroAppAdd(tx migration.LimitedTx) error {
	_, err := tx.Exec(`
		ALTER TABLE micro_app ADD column url varchar(255)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		ALTER TABLE micro_app ADD column swagger text
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		ALTER TABLE micro_app ADD column user_id varchar(255) NOT NULL
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		ALTER TABLE micro_app ADD column active varchar(1) DEFAULT 'Y' NOT NULL
	`)
	if err != nil {
		return err
	}


	return nil
}
