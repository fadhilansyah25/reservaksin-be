package session

import (
	"ca-reservaksin/businesses/session"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlSessionRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) session.Repository {
	return &MysqlSessionRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlSessionRepository) Create(data *session.Domain) (session.Domain, error) {
	dataSession := FromDomain(data)

	err := mysqlRepo.Conn.Create(&dataSession).Error
	if err != nil {
		return session.Domain{}, err
	}

	if err := mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).Find(&dataSession).Error; err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

func (mysqlRepo *MysqlSessionRepository) FetchAll() ([]session.Domain, error) {
	dataSession := []Session{}

	if err := mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).Find(&dataSession).Error; err != nil {
		return []session.Domain{}, err
	}

	return ToArrayOfDomain(dataSession), nil
}

func (mysqlRepo *MysqlSessionRepository) GetByID(id string) (session.Domain, error) {
	dataSession := Session{}

	err := mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).First(&dataSession, "id = ?", id).Error
	if err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

func (mysqlRepo *MysqlSessionRepository) GetByLatLong(lat, lng float64) ([]session.SessionDistance, error) {
	res := []SessionDistance{}
	qe := `SELECT 
		sessions.*,
		current_addresses.lat,
		current_addresses.lng,
		(
		6371.04 * acos(
			cos(radians(` + strconv.FormatFloat(lat, 'f', -1, 64) + `)) * cos(radians(lat)) * cos(radians(lng) 
			- radians(` + strconv.FormatFloat(lng, 'f', -1, 64) + `)) 
			+ sin(radians(` + strconv.FormatFloat(lat, 'f', -1, 64) + `)) * sin(radians(lat))
		)
		) AS distance
		FROM
			(
			sessions
			INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
			INNER JOIN current_addresses ON health_facilities.current_address_id = current_addresses.id
			)
		HAVING
			distance < 10
			&& date >= DATE(NOW())
		ORDER BY
			distance
		LIMIT
		0, 20;`
	err := mysqlRepo.Conn.Raw(qe).Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return ToArrayOfDomainResult(res), nil
}

func (mysqlRepo *MysqlSessionRepository) Update(id string, data *session.Domain) (session.Domain, error) {
	dataSession := FromDomain(data)
	err := mysqlRepo.Conn.Where("id = ?", id).Updates(&dataSession).Error
	if err != nil {
		return session.Domain{}, err
	}

	if err := mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).Find(&dataSession).Error; err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

func (mysqlRepo *MysqlSessionRepository) Delete(id string) (string, error) {
	recSession := Session{}
	err := mysqlRepo.Conn.Delete(&recSession, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}

func (mysqlRepo *MysqlSessionRepository) FetchByHistory(adminID, history string) ([]session.Domain, error) {
	dataSession := []Session{}

	switch history {
	case "upcoming":
		qe := `SELECT sessions.* FROM sessions INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		WHERE date > DATE(NOW()) && admin_id = ?`
		err := mysqlRepo.Conn.Raw(qe, adminID).Find(&dataSession).Error
		if err != nil {
			return nil, err
		}
	case "history":
		qe := `SELECT sessions.* FROM sessions INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		WHERE date < DATE(NOW()) && admin_id = ?`
		err := mysqlRepo.Conn.Raw(qe, adminID).Find(&dataSession).Error
		if err != nil {
			return nil, err
		}
	default:
		qe := `SELECT sessions.* FROM sessions INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		WHERE date = DATE(NOW()) && admin_id = ?`
		err := mysqlRepo.Conn.Raw(qe, adminID).Find(&dataSession).Error
		if err != nil {
			return nil, err
		}
	}

	mysqlRepo.Conn.Preload(clause.Associations).Find(&dataSession)

	return ToArrayOfDomain(dataSession), nil
}

func (mysqlRepo *MysqlSessionRepository) FetchAllByAdminID(adminID string) ([]session.Domain, error) {
	dataSession := []Session{}
	qe := `SELECT sessions.* FROM sessions INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		WHERE admin_id = ?`
	err := mysqlRepo.Conn.Raw(qe, adminID).Find(&dataSession).Error
	if err != nil {
		return nil, err
	}

	mysqlRepo.Conn.Preload(clause.Associations).Find(&dataSession)

	return ToArrayOfDomain(dataSession), nil
}
