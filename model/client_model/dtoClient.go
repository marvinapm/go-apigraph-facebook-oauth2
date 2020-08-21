package client_model

type InformacionResponse struct {
	Parametros  Parametros    `json:"parametros"`
	Informacion []Informacion `json:"informacion"`
}

type Informacion struct {
	ID        string `json:"id"`
	Tiutlo    string `json:"tiutlo"`
	URL       string `json:"url"`
	Usuaior   string `json:"usuaior"`
	Enlace    string `json:"enlace"`
	Fechahora string `json:"fechahora"`
}

type Parametros struct {
	Codigo string `json:"codigo"`
	Glosa  string `json:"glosa"`
}
