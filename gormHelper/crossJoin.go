package gormhelper

import arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"

// use for join by different db
// encourage to sort according given ids in getDerivedStructs, it will improve processing speed
func CrossJoin[T any, J any](
	mainStructs []T,
	getDerivedId func(mainStruct T) interface{},
	getDerivedStructs func(ids []interface{}) (derivedStructs []J, err error),
	isFulFill func(mainStruct T, derivedStruct J) bool,
	setDerived func(mainStruct *T, derivedStruct J),
) ([]T, error) {
	derivedIds := arrayhelper.Unique(
		arrayhelper.Map(mainStructs, getDerivedId),
	)
	derivedStructs, err := getDerivedStructs(derivedIds)
	if err != nil {
		return mainStructs, err
	}
	if len(derivedStructs) == 0 {
		return mainStructs, nil
	}
	lastDerivedId := 0
	for i := range mainStructs {
		derivedIds := arrayhelper.RotateByIndex(arrayhelper.Init(len(derivedStructs), func(index int) int {
			return index
		}),
			lastDerivedId,
		)
		for _, j := range derivedIds {
			// only need to check first one because it has been sort
			if isFulFill(mainStructs[i], derivedStructs[j]) {
				lastDerivedId = j
				setDerived(&mainStructs[i], derivedStructs[j])
				break
			}
		}
	}

	return mainStructs, err
}
