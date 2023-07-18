// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package pop

type Response struct {
	CountAppend int64  `json:"count_append,omitempty"`
	NewToken    string `json:"new_token"`
}
