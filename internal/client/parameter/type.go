package parameter

import "encoding/json"

type NullableInt64 struct {
	Value int64
	Valid bool
}

func (n NullableInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type NullableInt32 struct {
	Value int32
	Valid bool
}

func (n NullableInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type NullableString struct {
	Value string
	Valid bool
}

func (n NullableString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type NullableBool struct {
	Value bool
	Valid bool
}

func (n NullableBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type NullableObject[T any] struct {
	Value *T
	Valid bool
}

// MarshalJSON nullの場合は空のオブジェクトをペイロードに乗せます.
func (n NullableObject[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("{}"), nil
	}
	return json.Marshal(n.Value)
}
