package utils

// MatchByKey reorders the api slice to match the order of the ref slice using
// keys produced by keyFn. Elements in api whose key is not present in ref are
// appended at the end in their original relative order.
//
// This is used to preserve the user's configured ordering for collections
// represented as types.List, when the API response order is not guaranteed to
// match the plan/state order. The keyFn must return a stable identifier that
// uniquely identifies an element across api and ref.
func MatchByKey[T any](api, ref []T, keyFn func(T) string) []T {
	if len(ref) == 0 {
		return api
	}
	result := make([]T, 0, len(api))
	used := make([]bool, len(api))
	for _, r := range ref {
		rk := keyFn(r)
		for i, a := range api {
			if !used[i] && keyFn(a) == rk {
				result = append(result, a)
				used[i] = true
				break
			}
		}
	}
	for i, a := range api {
		if !used[i] {
			result = append(result, a)
		}
	}
	return result
}
