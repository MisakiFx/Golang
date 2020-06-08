package types

import (
	"fmt"
	"io"
	"time"
)

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

type Datetime struct {
	t time.Time
}

const TimeLayout = "2006-01-02T15:04:05.000Z"

//gqlgen回调，对Datetime类型进行反序列化
func (d *Datetime) UnmarshaGQL(vi interface{}) (err error) {
	v, ok := vi.(string)
	if !ok {
		return fmt.Errorf("unknow type of Datetime: `%+v`", vi)
	}
	//反序列化
	if d.t, err = time.Parse(TimeLayout, v); err != nil {
		return err
	}
	return nil
}

//gqlgen序列化Datetime方法
func (d *Datetime) MarshalGQL(w io.Writer) {
	w.Write([]byte(d.t.Format(TimeLayout)))
}
