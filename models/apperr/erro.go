package apperr

import "fmt"

// Todo erro será um tipo de dado
type Erro struct {
	Mensagem string //Mensagem informando em que ação ocorreu o erro
	Causa    error  //Erro original retornado
}

func (e *Erro) Error() string {
	if e.Causa != nil {
		return fmt.Sprintf("%s: %v", e.Mensagem, e.Causa)
	}

	return e.Mensagem
}
