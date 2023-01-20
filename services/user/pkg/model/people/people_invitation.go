package people

import (
	"time"

	"gorm.io/gorm"
)

const timeFormat = "2006-01-02 15:04:05"

type PeopleInvitationTable struct {
	Id        uint `gorm:"primary_key;auto_increment;not_null"`
	User      string
	Username  string
	Time      string
	Confirmed int `gorm:"default:0"`
}

type People struct {
	Username       string `json:"username"`
	Firstname      string `json:"firstname"`
	ProfilePicture string `json:"profilePicture"`
}

type AcceptInvite struct {
	Id    uint
	Value int
}

func (PeopleInvitationTable) TableName() string {
	return "people_invitations"
}

// Create new user invitatiom in DB
func CreatePeopleInvitation(db *gorm.DB, t *PeopleInvitationTable) (string, error) {
	var exists bool
	err := db.Table("users").Select("count(*) > 0").Where("username = ?", t.Username).Find(&exists).Error

	if exists {
		time := time.Now()
		t.Time = time.Format(timeFormat)
		return "User succesfully invited! 🎉", db.Create(t).Error
	}
	return "Sorry, this user does not exists 😔", err
}

// Get people from DB
func GetPeople(db *gorm.DB, t *People) ([]People, error) {
	query := `	SELECT
					*
				FROM
					users
				WHERE username IN (
							SELECT
								CASE T1.username
								WHEN '` + t.Username + `' THEN
									T1.user
								ELSE
									T2.username
								END AS username
							FROM
								people_invitations T1
								INNER JOIN people_invitations T2 ON T1.id = T2.id
							WHERE (T1.username = '` + t.Username + `'
								AND T1.confirmed = 1)
								OR(T2.user = '` + t.Username + `'
									AND T2.confirmed = 1))`

	people, err := GetPeopleFromQuery(db, query)
	if err != nil {
		return nil, err
	}

	return people, nil
}

func AcceptInvitation(db *gorm.DB, t *AcceptInvite) error {
	return db.Table("people_invitations").Where("id = ?", t.Id).Update("confirmed", t.Value).Error
}

func GetPeopleFromQuery(db *gorm.DB, query string) ([]People, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var people []People
	for rows.Next() {
		db.ScanRows(rows, &people)
	}

	return people, nil
}
