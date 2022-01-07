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

// func (mysqlRepo *MysqlSessionRepository) Update(id string, data *session.Domain) (session.Domain, error) {

// }

func (mysqlRepo *MysqlSessionRepository) GetByID(id string) (session.Domain, error) {
	dataSession := Session{}

	err := mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).First(&dataSession, "id = ?", id).Error
	if err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

func (mysqlRepo *MysqlSessionRepository) GetByLatLong(lat, lng float64) ([]session.Result, error) {
	res := []Result{}
	qe := `SELECT
	*,
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
		INNER JOIN vaccines ON sessions.vaccine_id = vaccines.id
		INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		INNER JOIN current_addresses ON health_facilities.current_address_id = current_addresses.id
		)
	HAVING
		distance < 5
		&& date >= DATE(NOW())
	ORDER BY
		distance
	LIMIT
	0, 20;`
	err := mysqlRepo.Conn.Raw(qe).Find(&res).Error
	if err != nil {
		return nil, err
	}

	for i := range res {
		mysqlRepo.Conn.Preload("HealthFacilites.CurrentAddress").Preload(clause.Associations).Find(&res[i].Session)
	}

	return ToArrayOfDomainResult(res), nil
}

// func (mysqlRepo *MysqlSessionRepository) Delete(id string) (string, error) {

// }
