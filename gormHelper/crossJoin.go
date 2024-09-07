package gormhelper

import arrayhelper "github.com/jbterrylin/go-helper/arrayHelper"

// use for join by different db
// getDerivedStructs shouldn't return paginated result
// encourage to sort according given ids in getDerivedStructs, it will improve processing speed
func CrossJoin[T any, J any, U comparable](
	mainStructs []T,
	getDerivedId func(mainStruct T) U,
	getDerivedStructs func(ids []U) (derivedStructs []J, err error),
	setDerived func(mainStruct *T, derivedStruct J) (nextMain bool),
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
	for i := range mainStructs {
		for j := range derivedStructs {
			nextMain := setDerived(&mainStructs[i], derivedStructs[j])
			if nextMain {
				break
			}
		}
	}

	return mainStructs, err
}

// use for join by different db
// getDerivedStructs shouldn't return paginated result
// encourage to sort according given ids in getDerivedStructs, it will improve processing speed
// because inner join
// so getMainStructs will run 2 times and second time will return derivedIds to filter again
// filterUnusedDerivedStruct is used to relief stress of nested loop
// filterUnusedDerivedStruct can just simply return true or function below
// func(derivedStruct J, ids []U) { arrayhelper.Includes(ids, derivedStruct.<fk>) }
func CrossInnerJoin[T any, J any, U comparable, K comparable](
	getMainStructs func(derivedIds []K, isFirstGet bool) ([]T, error),
	getDerivedId func(mainStruct T) U,
	getDerivedStructs func(ids []U) (derivedStructs []J, err error),
	getMainId func(derivedStruct J) K,
	filterUnusedDerivedStruct *func(derivedStruct J, ids []U) bool,
	setDerived func(mainStruct *T, derivedStruct J) (nextMain bool),
) ([]T, error) {
	var derivedStructIds []K
	mainStructs, err := getMainStructs(derivedStructIds, true)
	if err != nil {
		return mainStructs, err
	}
	if len(mainStructs) == 0 {
		return mainStructs, nil
	}

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

	derivedStructIds = arrayhelper.Unique(
		arrayhelper.Map(derivedStructs, getMainId),
	)
	mainStructs, err = getMainStructs(derivedStructIds, false)
	if err != nil {
		return mainStructs, err
	}

	if filterUnusedDerivedStruct != nil {
		derivedStructs = arrayhelper.Filter(derivedStructs, func(derivedStruct J) bool {
			return (*filterUnusedDerivedStruct)(derivedStruct, derivedIds)
		})
	}

	for i := range mainStructs {
		for j := range derivedStructs {
			nextMain := setDerived(&mainStructs[i], derivedStructs[j])
			if nextMain {
				break
			}
		}
	}

	return mainStructs, err
}
