# Installation

```go 
go get github.com/lucasmdomingues/goclima
```

# Examples

### Search city data by ID.

```go
import (
	"fmt"

	"github.com/lucasmdomingues/goclima"
)

func main() {
	token := "TOKEN"

	locale, err := goclima.GetLocaleByID(token, 3477)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Search city data by Name and/or State.

```go
import (
	"fmt"
	"github.com/lucasmdomingues/goclima"
)

func main() {
	token := "TOKEN"

	locale, err := goclima.GetLocaleByNameState(token, "São Paulo","SP")
	if err != nil {
		log.Fatal(err)
	}
}

```

### Currently time by city ID.

```go
import (
	"fmt"
	"github.com/lucasmdomingues/goclima"
)

func main() {
	token := "TOKEN"

	weather, err := goclima.GetWeather(token, 3477)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Climatic rain by city ID.

```go
import (
	"fmt"
	"github.com/lucasmdomingues/goclima"
)

func main() {
	token := "TOKEN"

	climate, err := goclima.GetClimate(token, 3477)
	if err != nil {
		log.Fatal(err)
	}
}
```
### Clima Tempo
https://advisor.climatempo.com.br/

### TO DO

* GeoreferencedForecast
* History
* Forecast
* Index
