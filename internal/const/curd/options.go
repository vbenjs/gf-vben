package curd

import "Gf-Vben/internal/model/entity"

type Options interface {
	entity.Permission | entity.Router
	GetOptionLabel() string
	GetOptionValue() string
	GetOptionKey() string
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Key   string `json:"key"`
}

func BuildOptions[T Options](a []T) (res []Option) {

	for _, t := range a {
		res = append(res, Option{
			Label: t.GetOptionLabel(),
			Value: t.GetOptionKey(),
			Key:   t.GetOptionKey(),
		})
	}
	return
}
