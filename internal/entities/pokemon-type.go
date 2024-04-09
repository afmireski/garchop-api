package entities

type PokemonType struct {	
	Id string `json:"id"`
	ReferenceId uint `json:"reference_id"`
	Name string `json:"name"`
}

func NewPokemonType(id string, referenceId uint, name string) *PokemonType {
	return &PokemonType{
		Id: id,
		ReferenceId: referenceId,
		Name: name,
	}
}