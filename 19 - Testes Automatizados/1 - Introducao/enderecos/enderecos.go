package enderecos

import "strings"

// TipoDeEndereco verifica se o tipo é valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Rodovia"}

	primeiraPalavra := strings.Split(endereco, " ")[0]

	for _, tipo := range tiposValidos {
		if strings.ToLower(primeiraPalavra) == strings.ToLower(tipo) {
			return strings.Title(strings.ToLower(primeiraPalavra))
		}
	}
	return "Não identificado"
}
