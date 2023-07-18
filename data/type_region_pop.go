// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

type RegionPop struct {
	RegionCode string `json:"region_code" gorm:"primaryKey"`
	Count      int64  `json:"count" gorm:"not null"`
}
