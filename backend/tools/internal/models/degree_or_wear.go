package models

type DegreeOfWear struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (DegreeOfWear) TableName() string {
	return "degree_of_wears"
}

type EDegreeOfWear int

const (
	KDegreeOfWearUnknown EDegreeOfWear = iota
	KDegreeOfWearNew
	KDegreeOfWearBroken
	KDegreeOfWearWornOut
	KDegreeOfWearUtilized
)

func (e EDegreeOfWear) Name() string {
	switch e {
	case KDegreeOfWearNew:
		return "Новый"
	case KDegreeOfWearBroken:
		return "Сломанный"
	case KDegreeOfWearWornOut:
		return "Надежный"
	case KDegreeOfWearUtilized:
		return "Изношенный"
	default:
		return "Неизвестный"
	}
}
