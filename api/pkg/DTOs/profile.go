package dtos

// CreateProfileDTO representa os dados esperados ao criar um perfil.
type CreateProfileDTO struct {
	CompleteName   string  `json:"completeName"`
	CellPhone      string  `json:"cellphone"`
	InitialBalance float64 `json:"initialBalance"`
}

// UpdateProfileDTO representa os dados para o PATCH (atualização parcial).
type UpdateProfileDTO struct {
	CompleteName   *string  `json:"completeName"`
	CellPhone      *string  `json:"cellphone"`
	InitialBalance *float64 `json:"initialBalance"`
}

// ProfileResponseDTO representa os dados devolvidos ao cliente.
type ProfileResponseDTO struct {
	ID             uint    `json:"id"`
	UserId         uint    `json:"user_id"`
	CompleteName   string  `json:"completeName"`
	CellPhone      string  `json:"cellphone"`
	InitialBalance float64 `json:"initialBalance"`
}
