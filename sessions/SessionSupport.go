package sessions

type SessionSupport struct {
	NewSession func() (Session, error) //session为事务操作
}
