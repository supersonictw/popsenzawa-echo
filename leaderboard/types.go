// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package leaderboard

type Response struct {
	Global  int64            `json:"global"`
	Regions map[string]int64 `json:"regions"`
}
