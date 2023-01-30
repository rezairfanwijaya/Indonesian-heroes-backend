package hero

type Hero struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	BirthYear   int32
	DeathYear   int32
	Description string
	Age         string
}
