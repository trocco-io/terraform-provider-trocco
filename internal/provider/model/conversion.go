package model

func int64PtrFromInt32Ptr(i32ptr *int32) *int64 {
	if i32ptr == nil {
		return nil
	}
	i64 := int64(*i32ptr)
	return &i64
}

func int32PtrFromInt64Ptr(i64ptr *int64) *int32 {
	if i64ptr == nil {
		return nil
	}
	i32 := int32(*i64ptr)
	return &i32
}