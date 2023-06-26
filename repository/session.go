package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	err := u.db.Where("token =?", token).Delete(&model.Session{})
	return err.Error // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	if fd := u.db.Model(
		model.Session{}).Where("email =?", session.Email).Updates(
		session); fd.Error != nil {
		return fd.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	var availemail model.Session
	err := u.db.Where("email = ?", email).First(&availemail).Error
	if err != nil {
		return model.Session{}, err
	}
	return availemail, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	var availtoken model.Session
	err := u.db.Where("token = ?", token).First(&availtoken).Error
	if err != nil {
		return model.Session{}, err
	}
	return availtoken, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
