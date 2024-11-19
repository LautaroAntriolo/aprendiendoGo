package fechas

import (
	"fmt"
	"time"
	"os/exec"
)

var localizacion *time.Location = time.Now().Location()
var ahora time.Time = time.Now()

func TiempoANavidad() string {
	anioActual := time.Now().Year()
	navidad := time.Date(anioActual, time.December, 25, 0, 0, 0, 0, localizacion)	
	diferencia := navidad.Sub(ahora);
	
	if diferencia == 0 {
		cmd := exec.Command("notify-send", "Recordatorio", "¡Feliz Navidad!")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error al enviar la notificación:", err)
		}
	} else {
		// Calculamos las horas, minutos y segundos a partir de la duración
		dias := int(diferencia.Hours())/24;
		horas := int(diferencia.Hours()) % 24;
		minutos := int(diferencia.Minutes()) % 60;
		segundos := int(diferencia.Seconds()) % 60;
		mensaje := fmt.Sprintf("Faltan %d días, %d horas, %d minutos y %d segundos", dias, horas, minutos, segundos)
		fmt.Println(mensaje);
		return mensaje
	}
	return ""
}

func Edad(anio int, mes time.Month, dia int){
	miCumple := time.Date(anio,mes,dia,0,0,0,0,localizacion);
	edad := ahora.Year() - miCumple.Year()
	if ahora.Month() < miCumple.Month() || (ahora.Month() == miCumple.Month() && ahora.Day() < miCumple.Day()) {
		edad--
	}
	fmt.Println("Edad: ", edad)
}

func Recordatorio(anio int, mes time.Month, dia int, hora int, minutos int, mensaje string){
	fecha := time.Date(anio,mes,dia,hora,minutos,0,0,localizacion)
	// fmt.Println("El mensaje es el siguiente: ", mensaje)
	if ahora.Year() == fecha.Year() && ahora.Month() == fecha.Month() && ahora.Hour() == fecha.Hour() && ahora.Minute() == fecha.Minute(){
		fmt.Println()
		cmd := exec.Command("notify-send", "Recordatorio", mensaje)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error al enviar la notificación:", err)
		}
	}
}