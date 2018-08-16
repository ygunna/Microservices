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


	return nil
}
