package main // paketimizin ismi, her go uygulaması çalışmaya main paketinden başlar

import ( // kullanacağımız kütüphaneleri ekliyoruz (import)
	"fmt"
	"math/rand"
)

func main() { // her uygulama çalışmaya main fonksiyonundan başlar

	fmt.Println("Aklımda tuttuğum sayı ", rand.Intn(10))

}
