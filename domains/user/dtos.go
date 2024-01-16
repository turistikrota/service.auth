package user

type ListDto struct{}

func (e *Entity) ToListDto() ListDto {
	return ListDto{}
}
