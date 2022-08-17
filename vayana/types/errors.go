package vayanaTypes

type Error struct {
	Message string    `json:"message"`
	Args    ErrorArgs `json:"args"`
}

type ErrorArgs struct {
	ParameterName string `json:"parameter-name"`
}
