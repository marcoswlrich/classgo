package users

import (
	"fmt"
	"time"

	"github.com/marcoswlrich/classgo/modelos"
)

func NovoUsuario() {
	u := new(modelos.User)
	u.AddUser(12, "Marcos", time.Now(), true)
	fmt.Println(u)
}
