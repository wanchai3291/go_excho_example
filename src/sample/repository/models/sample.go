package SampleModel

type Sample struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}
