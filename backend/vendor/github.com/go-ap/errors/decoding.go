package errors

import (
	"github.com/valyala/fastjson"
)

func UnmarshalJSON(data []byte) ([]Error, error) {
	if len(data) == 0 {
		return nil, nil
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	v := val.Get("errors")
	if v == nil {
		return nil, wrap(nil, "invalid errors array")
	}
	items := make([]Error, 0)
	switch v.Type() {
	case fastjson.TypeArray:
		for _, v := range v.GetArray() {
			status := v.GetInt("status")
			localErr := errorFromStatus(status)
			if err := localErr.UnmarshalJSON([]byte(v.String())); err == nil {
				items = append(items, localErr)
			}
		}
		return items, err
	case fastjson.TypeObject:
		status := v.GetInt("status")
		localErr := errorFromStatus(status)
		if err := localErr.UnmarshalJSON(data); err == nil {
			items = append(items, localErr)
		}
	case fastjson.TypeString:
		it := new(Err)
		it.m = string(data)
		items = append(items, it)
	}
	return items, nil
}

func (e *Err) UnmarshalJSON(data []byte) error {
	if m := fastjson.GetString(data, "message"); len(m) > 0 {
		e.m = m
	}
	return nil
}

func (n *notFound) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (m *methodNotAllowed) UnmarshalJSON(data []byte) error {
	return m.Err.UnmarshalJSON(data)
}
func (n *notValid) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (f *forbidden) UnmarshalJSON(data []byte) error {
	return f.Err.UnmarshalJSON(data)
}
func (n *notImplemented) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (b *badRequest) UnmarshalJSON(data []byte) error {
	return b.Err.UnmarshalJSON(data)
}
func (u *unauthorized) UnmarshalJSON(data []byte) error {
	return u.Err.UnmarshalJSON(data)
}
func (n *notSupported) UnmarshalJSON(data []byte) error {
	return n.Err.UnmarshalJSON(data)
}
func (t *timeout) UnmarshalJSON(data []byte) error {
	return t.Err.UnmarshalJSON(data)
}
func (b *badGateway) UnmarshalJSON(data []byte) error {
	return b.Err.UnmarshalJSON(data)
}
