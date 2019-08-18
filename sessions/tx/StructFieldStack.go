package tx

import "reflect"

type StructField struct {
	i    int
	data []reflect.StructField //方法队列
}

func (t StructField) New() StructField {
	return StructField{
		data: []reflect.StructField{},
		i:    0,
	}
}

func (t *StructField) Push(k reflect.StructField) {
	t.data = append(t.data, k)
	t.i++
}

func (t *StructField) Pop() (ret reflect.StructField) {
	t.i--
	ret = t.data[t.i]
	t.data = t.data[0:t.i]
	return
}

func (t *StructField) Len() int {
	return t.i
}
