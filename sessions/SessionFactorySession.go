package sessions

import (
	"context"
	"github.com/agui2200/GoMybatis/sessions/tx"
	"github.com/agui2200/GoMybatis/utils"
)

type SessionFactorySession struct {
	Session Session
	Factory *SessionFactory
}

func (it *SessionFactorySession) Id() string {
	if it.Session == nil {
		return ""
	}
	return it.Session.Id()
}
func (it *SessionFactorySession) Query(sqlorArgs string) ([]map[string][]byte, error) {
	if it.Session == nil {
		return nil, utils.NewError("SessionFactorySession", " can not run Id(),it.Session == nil")
	}
	return it.Session.Query(sqlorArgs)
}
func (it *SessionFactorySession) Exec(sqlorArgs string) (*Result, error) {
	if it.Session == nil {
		return nil, utils.NewError("SessionFactorySession", " can not run Exec(),it.Session == nil")
	}
	return it.Session.Exec(sqlorArgs)
}
func (it *SessionFactorySession) Rollback() error {
	if it.Session == nil {
		return utils.NewError("SessionFactorySession", " can not run Rollback(),it.Session == nil")
	}
	return it.Session.Rollback()
}
func (it *SessionFactorySession) Commit() error {
	if it.Session == nil {
		return utils.NewError("SessionFactorySession", " can not run Commit(),it.Session == nil")
	}
	return it.Session.Commit()
}

func (it *SessionFactorySession) Begin() error {
	if it.Session == nil {
		return utils.NewError("SessionFactorySession", " can not run Begin(),it.Session == nil")
	}
	return it.Session.Begin()
}

func (it *SessionFactorySession) BeginTrans(p tx.Propagation) (err error) {
	if it.Session == nil {
		return utils.NewError("SessionFactorySession", " can not run Begin(),it.Session == nil")
	}
	return it.Session.BeginTrans(p)
}

func (it *SessionFactorySession) Close() {
	var id = it.Id()
	var s, _ = it.Factory.SessionMap.Load(id)
	if s != nil {
		if it.Session != nil {
			it.Session.Close()
		}
		it.Factory.SessionMap.Delete(id)
	}
}

func (it *SessionFactorySession) LastPROPAGATION() *tx.Propagation {
	return it.Session.LastPROPAGATION()
}

func (it *SessionFactorySession) WithContext(ctx context.Context) {
	it.Session.WithContext(ctx)
}
