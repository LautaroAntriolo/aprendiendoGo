package datos
import (
	"fmt"
	"log"
	"os"
)

func Lecturas(ruta string) {
    datosComoBytes, err := os.ReadFile(ruta)
    if err != nil {
        log.Fatal(err)
	}
    datosComoString := string(datosComoBytes)
    
    fmt.Println(datosComoString)
}
