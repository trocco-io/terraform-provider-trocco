package utils

// MatchByKey reorders the api slice to match the order of the ref slice using
// keys produced by primaryKey. For each ref whose primaryKey has no match in
// api, it falls back to fallbackKey for that ref. Elements in api whose key
// is not consumed by ref are appended at the end in their original relative
// order.
//
// fallbackKey may be nil; in that case unmatched refs are simply skipped and
// the corresponding api elements end up at the tail of the result.
//
// This is used to preserve the user's configured ordering for collections
// represented as types.List, when the API response order is not guaranteed to
// match the plan/state order.
//
// The two-key design supports the case where ref carries an unknown id at
// plan time (e.g. on CREATE, where API-assigned ids are not yet populated).
// In that case primaryKey — which typically includes the id — cannot match,
// and fallbackKey — which typically uses only the user-visible identity
// (type, destination_type, …) — provides positional matching among the
// remaining elements.
func MatchByKey[T any](api, ref []T, primaryKey, fallbackKey func(T) string) []T {
	if len(ref) == 0 {
		return api
	}
	result := make([]T, 0, len(api))
	used := make([]bool, len(api))
	for _, r := range ref {
		rk := primaryKey(r)
		matched := false
		for i, a := range api {
			if !used[i] && primaryKey(a) == rk {
				result = append(result, a)
				used[i] = true
				matched = true
				break
			}
		}
		if matched || fallbackKey == nil {
			continue
		}
		rk = fallbackKey(r)
		for i, a := range api {
			if !used[i] && fallbackKey(a) == rk {
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
