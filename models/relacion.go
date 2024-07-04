package models

type Relacion struct {
	UsuarioId         string `bson:"usuarioId" json:"usuarioId"`
	UsuarioRelacionId string `bson:"usuarioRelacionId" json:"usuarioRelacionId"`
}
