package main
import (
	"fmt"
)
//=============================================
//Usuario ...
type Usuario interface {
	Nombre() string
	Email() string
}

//Persona ...
type Persona struct {
	nombre string
	email  string
	edad   int
}
func (this Persona) Nombre() string {
	return this.nombre
}
func (this Persona) Email() string {
	return this.email
}

func Presentarse(this Usuario) {
	fmt.Println("Nombre: ", this.Nombre())
	fmt.Println("Email: ", this.Email())
}
//=============================================


//=============================================
//Moderador ...
type Moderador struct {
	Persona
	Foro string
}
//=============================================

//============================================
//Administrador ...
type Administrador struct {
	Persona
	Sección string
}
//============================================


func main() {
	alejandro := Persona{"Alejandro", "a@hmail.com", 29}
	Presentarse(alejandro)
	juan := Moderador{Persona{"Juan", "j@hmail.com", 46}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "p@hmail.com", 25}, "PC"}
	Presentarse(juan)
	Presentarse(pedro)

}
