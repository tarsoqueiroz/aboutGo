package inspect

import "cmdlnCobraStringer/pkg/geral"

func Inspect(input string, digits bool) (count int, kind string) {

	if !digits {
		return len(input), "char"
	}

	return geral.InspectNumbers(input), "digit"

}
