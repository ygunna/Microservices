package db

import (
	sq "github.com/Masterminds/squirrel"
	"crossent/micro/studio/db/lock"
	"crossent/micro/studio/domain"
)

type ComposeRepository interface {
	ID() int
	ListCompose() ([]domain.Compose, error)
	GetMicroservice(int) (domain.Compose, error)
	CreateMicroservice(domain.ComposeRequest) (int, error)
	UpdateMicroservice(domain.ComposeRequest) (bool, error)
	UpdateMicroserviceStatus(domain.ComposeRequest) (bool, error)
	CreateMicroserviceService(domain.MicroserviceService) (bool, error)
	CreateMicroserviceApp(domain.MicroserviceApp) (bool, error)
	UpdateMicroserviceApp(domain.MicroserviceApp) (bool, error)
	ListMicroserviceAppApp(int) ([]domain.MicroserviceApp, error)
	ListMicroserviceAppService(int) ([]domain.MicroserviceService, error)
	GetMicroserviceApp(int, string) (domain.MicroserviceApp, error)

	ListMicroserviceByName(string) (int, error)
	DeleteMicroserviceService(domain.MicroserviceService) error
	DeleteMicroserviceApp(domain.MicroserviceApp) error
}

type composeRepository struct {
	id          int
	conn        Conn
	lockFactory lock.LockFactory
	name        string

}

func (c *composeRepository) ID() int { return c.id }

func newComposeRepository(conn Conn, lockFactory lock.LockFactory) *composeRepository {
	return &composeRepository{
		conn:        conn,
		lockFactory: lockFactory,
	}
}

func (c *composeRepository) ListCompose() ([]domain.Compose, error) {
	return nil, nil
}

func (c *composeRepository) CreateMicroservice(r domain.ComposeRequest) (int, error) {
	var id int
	err := psql.Insert("micro_app").
		Columns("id", "name", "org_guid", "space_guid", "version", "description", "visible", "status, user_id").
		Values(sq.Expr("nextval('micro_app_id_seq')"), r.Name, r.OrgGuid, r.SpaceGuid, r.Version, r.Description, r.Visible, r.Status, r.UserId).
		Suffix("RETURNING id").
		RunWith(c.conn).
		QueryRow().
		Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *composeRepository) UpdateMicroservice(r domain.ComposeRequest) (bool, error) {
	if r.Url != "" {
		_, err := psql.Update("micro_app").
		//Set("name", r.Name).
			Set("version", r.Version).
			Set("visible", r.Visible).
			Set("status", r.Status).
			Set("url", r.Url).
			Where(sq.Eq{"id": r.ID}).
			RunWith(c.conn).
			Exec()
		if err != nil {
			return false, err
		}
	} else {
		_, err := psql.Update("micro_app").
		//Set("name", r.Name).
			Set("version", r.Version).
			Set("visible", r.Visible).
			Set("status", r.Status).
			Where(sq.Eq{"id": r.ID}).
			RunWith(c.conn).
			Exec()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (c *composeRepository) UpdateMicroserviceStatus(r domain.ComposeRequest) (bool, error) {
	_, err := psql.Update("micro_app").
		Set("status", r.Status).
		Where(sq.Eq{"id": r.ID}).
		RunWith(c.conn).
		Exec()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *composeRepository) CreateMicroserviceService(r domain.MicroserviceService) (bool, error) {
	result, err := psql.Insert("micro_app_service").
		Columns("id", "micro_id, service_guid").
		Values(sq.Expr("nextval('micro_app_service_id_seq')"), r.MicroID, r.ServiceGuid).
		RunWith(c.conn).
		Exec()
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows == 1, nil
}

func (c *composeRepository) CreateMicroserviceApp(r domain.MicroserviceApp) (bool, error) {
	result, err := psql.Insert("micro_app_app").
		Columns("id", "micro_id, app_guid, source_guid, essential").
		Values(sq.Expr("nextval('micro_app_app_id_seq')"), r.MicroID, r.AppGuid, r.SourceGuid, r.Essential).
		RunWith(c.conn).
		Exec()
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows == 1, nil
}

func (c *composeRepository) UpdateMicroserviceApp(r domain.MicroserviceApp) (bool, error) {
	_, err := psql.Update("micro_app_app").
		Set("app_guid", r.AppGuid).
		Where(sq.And{sq.Eq{"source_guid": r.SourceGuid}, sq.Eq{"micro_id": r.MicroID}}).
		RunWith(c.conn).
		Exec()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *composeRepository) GetMicroservice(id int) (domain.Compose, error) {
	compose := domain.Compose{}
	err := psql.Select("id, name").
		From("micro_app").
		Where(sq.Eq{"id": id}).
		RunWith(c.conn).
		QueryRow().
		Scan(&compose.ID, &compose.Name)
	if err != nil {
		return compose, err
	}

	return compose, nil
}

func (c *composeRepository) ListMicroserviceAppApp(id int) ([]domain.MicroserviceApp, error) {
	rows, err := psql.Select(`
			micro_id,
			app_guid,
			source_guid
		`).
		From("micro_app_app").
		Where(sq.Eq{"micro_id": id}).
		RunWith(c.conn).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apps := []domain.MicroserviceApp{}
	for rows.Next() {
		app := domain.MicroserviceApp{}
		err := rows.Scan(&app.MicroID, &app.AppGuid, &app.SourceGuid)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (c *composeRepository) ListMicroserviceAppService(id int) ([]domain.MicroserviceService, error) {
	rows, err := psql.Select(`
			micro_id,
			service_guid
		`).
		From("micro_app_service").
		Where(sq.Eq{"micro_id": id}).
		RunWith(c.conn).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := []domain.MicroserviceService{}
	for rows.Next() {
		service := domain.MicroserviceService{}
		err := rows.Scan(&service.MicroID, &service.ServiceGuid)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func (c *composeRepository) GetMicroserviceApp(id int, essential string) (domain.MicroserviceApp, error) {
	app := domain.MicroserviceApp{}
	err := psql.Select("id, app_guid").
		From("micro_app_app").
		Where(sq.And{sq.Eq{"micro_id": id}, sq.Eq{"essential": essential}}).
		RunWith(c.conn).
		QueryRow().
		Scan(&app.ID, &app.AppGuid)
	if err != nil {
		return app, err
	}

	return app, nil
}

func (c *composeRepository) ListMicroserviceByName(name string) (int, error) {
	var count int
	err := psql.Select("count(id)").
		From("micro_app").
		Where(sq.Eq{"name": name}).
		RunWith(c.conn).
		QueryRow().
		Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}


func (c *composeRepository) DeleteMicroserviceService(r domain.MicroserviceService) error {
	_, err := psql.Delete("micro_app_service").
		Where(sq.Eq{
		"service_guid": r.ServiceGuid,
	}).
		RunWith(c.conn).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func (c *composeRepository) DeleteMicroserviceApp(r domain.MicroserviceApp) error {
	_, err := psql.Delete("micro_app_app").
		Where(sq.Eq{
		"app_guid": r.AppGuid,
	}).
		RunWith(c.conn).
		Exec()

	if err != nil {
		return err
	}

	return nil
}