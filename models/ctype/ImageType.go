package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	Cloud ImageType = 2
)

func (img ImageType) MarshJSON() ([]byte, error) {
	return json.Marshal(img.String())
}

func (img ImageType) String() string {
	var str string
	switch img {
	case Local:
		str = "本地"
	case Cloud:
		str = "云端"
	default:
		str = "其他"
	}
	return str
}
