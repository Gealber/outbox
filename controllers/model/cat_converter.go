package model

import "github.com/Gealber/outbox/repositories/model"

func (request CreateCatRequest) ToRepoModel() model.Cat {
	return model.Cat{
		Name:         request.Name,
		Color:        request.Color,
		Weight:       request.Weight,
		Intelligence: request.Intelligence,
		Laziness:     request.Laziness,
		Curiosity:    request.Curiosity,
		Sociability:  request.Sociability,
		Egoism:       request.Egoism,
		MiauPower:    request.MiauPower,
		Attack:       request.Attack,
	}
}

func (request UpdateCatRequest) ToRepoModel() model.Cat {
	return model.Cat{
		Name:         request.Name,
		Color:        request.Color,
		Weight:       request.Weight,
		Intelligence: request.Intelligence,
		Laziness:     request.Laziness,
		Curiosity:    request.Curiosity,
		Sociability:  request.Sociability,
		Egoism:       request.Egoism,
		MiauPower:    request.MiauPower,
		Attack:       request.Attack,
	}
}

func (CatResponse) FromRepoModel(cat *model.Cat) CatResponse {
	data := Cat{
		ID:           cat.ID,
		Name:         cat.Name,
		Color:        cat.Color,
		Weight:       cat.Weight,
		Intelligence: cat.Intelligence,
		Laziness:     cat.Laziness,
		Curiosity:    cat.Curiosity,
		Sociability:  cat.Sociability,
		Egoism:       cat.Egoism,
		MiauPower:    cat.MiauPower,
		Attack:       cat.Attack,
	}

	return CatResponse{Cat: data}
}
