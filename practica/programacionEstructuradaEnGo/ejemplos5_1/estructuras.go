package ejemplos5_1

import "fmt"

type Phone struct {
	nombre             string
	bateria            int
	capacidad_memoria  []int
	definicionDeCamara string
	precio             float32
}

func Celulares() {
	// Esta es la mejor forma de referenciar una estructura.
	my_phone := Phone{
		nombre:             "Iphone",
		bateria:            3500,
		capacidad_memoria:  []int{1, 2},
		definicionDeCamara: "18 mpx",
		precio:             1500.5,
	}
	fmt.Printf("Las caracteristicas de mi celular %s son: \n", my_phone.nombre)
	fmt.Println("Batería: ", my_phone.bateria)
	fmt.Println("capacidad de memoria: ", my_phone.capacidad_memoria)
	fmt.Println("camara: ", my_phone.definicionDeCamara)
	fmt.Println("precio: ", my_phone.precio)

	my_phone_2 := &my_phone
	my_phone_2.bateria, my_phone_2.nombre = 3000, "samsung"
	fmt.Printf("Las caracteristicas de mi celular %s son: \n", my_phone_2.nombre)
	fmt.Println("Batería: ", my_phone_2.bateria)
	fmt.Println("capacidad de memoria: ", my_phone_2.capacidad_memoria)
	fmt.Println("camara: ", my_phone_2.definicionDeCamara)
	fmt.Println("precio: ", my_phone_2.precio)
}

type Persona struct{
	Nombre    string
	Apellido  string
	Edad      int
	DNI       int
	Telefono  int
	CantHijos int
	Hijos     []*Persona
}

func Familia(){
	padre := Persona{"Luis","Antriolo",56,15464778,2332444888,2,[]*Persona{}}

	hijo1 := &Persona{
		Nombre:    "Lautaro",
		Apellido:  padre.Apellido,
		Edad:      28,
		DNI:       40123456,
		Telefono:  2325446677,
		CantHijos: 0,
	}

	hijo2 := &Persona{
		Nombre:    "Franco",
		Apellido:  padre.Apellido,
		Edad:      32,
		DNI:       41789012,
		Telefono:  2325448899,
		CantHijos: 0,
	}

	padre.Hijos = append(padre.Hijos, hijo1)
	padre.Hijos = append(padre.Hijos, hijo2)
	fmt.Print(padre)
}