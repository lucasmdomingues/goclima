# Instalação

```go 
go get github.com/lucasmdomingues/goclima
```

# Como usar

### Busca dados de cidade por ID.

```go
import (
	"fmt"
	"goclima"
)

func main() {

	token := "TOKEN"

	locale, err := goclima.GetLocaleByID(token, 3477)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(locale)
}
```

### Busca dados de cidades por Nome e/ou Estado.

```go
import (
	"fmt"
	"goclima"
)

func main() {

	token := "TOKEN"

	locale, err := goclima.GetLocaleByNameState(token, "São Paulo","SP")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(locale)
}

```

### Tempo no momento por ID da cidade.

```go
import (
	"fmt"
	"goclima"
)

func main() {

	token := "TOKEN"

	weather, err := goclima.GetWeather(token, 3477)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(weather)
}
```

### Chuva climática por ID da cidade.

```go
import (
	"fmt"
	"goclima"
)

func main() {

	token := "TOKEN"

	climate, err := goclima.GetClimate(token, 3477)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(climate)
}
```

