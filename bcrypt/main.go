package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "minhaSenhaSuperSegura123"
	
	// Hashing a senha
	hash, err := hashPassword(password)
	if err != nil {
		fmt.Println("Erro ao gerar hash:", err)
		return
	}
	fmt.Println("Hash gerado:", hash)
	
	// Verificando a senha
	match := checkPasswordHash(password, hash)
	fmt.Println("Senha válida?", match)
}
