package seguranca

import "golang.org/x/crypto/bcrypt"

//Recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//Compara uma senha e um hash, retorna se são iguais
func VerificarSenha(senha string, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha), []byte(senhaComHash))
}
