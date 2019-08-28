package tx

import (
	"database/sql"
	"sync"
)

type TxStack struct {
	i            int
	data         []*sql.Tx      //队列
	propagations []*Propagation //队列
	l            sync.Mutex     // 队列锁
}

func (t TxStack) New() TxStack {
	return TxStack{
		data:         []*sql.Tx{},
		propagations: []*Propagation{},
		i:            0,
	}
}

func (t *TxStack) Push(k *sql.Tx, p *Propagation) {
	t.l.Lock()
	t.data = append(t.data, k)
	t.propagations = append(t.propagations, p)
	t.i++
	t.l.Unlock()
}

func (t *TxStack) Pop() (*sql.Tx, *Propagation) {
	if t.i == 0 {
		return nil, nil
	}
	t.l.Lock()
	t.i--
	var ret = t.data[t.i]
	t.data = t.data[0:t.i]

	var p = t.propagations[t.i]
	t.propagations = t.propagations[0:t.i]
	t.l.Unlock()
	return ret, p
}
func (t *TxStack) First() (*sql.Tx, *Propagation) {
	if t.i == 0 {
		return nil, nil
	}
	var ret = t.data[0]
	var p = t.propagations[0]
	return ret, p
}
func (t *TxStack) Last() (*sql.Tx, *Propagation) {
	if t.i == 0 {
		return nil, nil
	}
	var ret = t.data[t.i-1]
	var p = t.propagations[t.i-1]
	return ret, p
}

func (t *TxStack) Len() int {
	return t.i
}

func (t *TxStack) HaveTx() bool {
	return t.Len() > 0
}
