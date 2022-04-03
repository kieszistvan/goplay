package strings

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func Sup(name string) string {
	return fmt.Sprintf("Sup' %s", name)
}

func MaxInt(i ...int) int {
	return funk.MaxInt(i)
}
