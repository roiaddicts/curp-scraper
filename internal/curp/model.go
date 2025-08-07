package curp

type curpResponse struct {
	Registros []struct {
		Sexo               string `json:"sexo"`
		ClaveEntidad       string `json:"claveEntidad"`
		DatosDocProbatorio struct {
			EntidadRegistro        string `json:"entidadRegistro"`
			ClaveMunicipioRegistro string `json:"claveMunicipioRegistro"`
			MunicipioRegistro      string `json:"municipioRegistro"`
			AnioReg                string `json:"anioReg"`
			Foja                   string `json:"foja"`
			Tomo                   string `json:"tomo"`
			Libro                  string `json:"libro"`
			NumActa                string `json:"numActa"`
			ClaveEntidadRegistro   string `json:"claveEntidadRegistro"`
		} `json:"datosDocProbatorio"`
		Parametro       string `json:"parametro"`
		StatusCurp      string `json:"statusCurp"`
		Nombres         string `json:"nombres"`
		DocProbatorio   int    `json:"docProbatorio"`
		Nacionalidad    string `json:"nacionalidad"`
		PrimerApellido  string `json:"primerApellido"`
		Entidad         string `json:"entidad"`
		SegundoApellido string `json:"segundoApellido"`
		Curp            string `json:"curp"`
		FechaNacimiento string `json:"fechaNacimiento"`
	} `json:"registros"`
	Codigo  string `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CurpModel struct {
	Curp                   string `json:"curp"`
	Nombres                string `json:"nombres"`
	PrimerApellido         string `json:"primerApellido"`
	SegundoApellido        string `json:"segundoApellido"`
	ClaveGenero            string `json:"claveGenero"`
	Genero                 string `json:"genero"`
	FechaNacimiento        string `json:"fechaNacimiento"`
	DiaNacimiento          string `json:"diaNacimiento"`
	MesNacimiento          string `json:"mesNacimiento"`
	AnioNacimiento         string `json:"anioNacimiento"`
	ClaveEntidadNacimiento string `json:"claveEntidadNacimiento"`
	EntidadNacimiento      string `json:"entidadNacimiento"`
}
