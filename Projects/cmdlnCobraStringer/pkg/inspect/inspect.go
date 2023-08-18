package inspect

import cmdlnCobraStringer/pkg/geral

func Inspect(input string, digits bool) (count int, kind string) {

	if !digits {
		return len(input), "char"
	}

	return InspectNumbers(input), "digit"
	// return InspectNumbers(input), "digit"

}
