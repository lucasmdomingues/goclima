# Instalação


``` 
go get github.com/lucasmdomingues/goclima
```

# Como usar

# Busca dados de cidade por ID.

´´´
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
´´´

# Busca dados de cidades por Nome e/ou Estado.

```

```

# Chuva climática por ID da cidade ou latitude e longitude.

```

```

# Tempo no momento por ID da cidade.

```

```

