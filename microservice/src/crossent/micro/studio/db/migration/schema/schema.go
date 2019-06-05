package schema

import "crossent/micro/studio/db/migration"

func New() []migration.Migrator {
	return []migration.Migrator{
		InitialSchema,
		//MicroAppAdd,
		MicroApI,
	}
}
