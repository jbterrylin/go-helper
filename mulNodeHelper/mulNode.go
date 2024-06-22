package mulnodehelper

func NewMulNode[T any, J any](initInfos []T, initFunc func(T) (key string, node J, err error)) (nodeMap map[string]J, err error) {
	nodeMap = map[string]J{}
	for _, initInfo := range initInfos {
		var key string
		var node J
		key, node, err = initFunc(initInfo)
		if err != nil {
			return
		}
		nodeMap[key] = node
	}
	return
}

// usage example

// var gormMap map[string]*gorm.DB
// var gormDefaultKey string

// type DbInfo struct {
// 	KeyName string
// 	Conn    string
// }

// func Gorm(key ...string) *gorm.DB {
// 	if len(key) == 0 {
// 		return gormMap[gormDefaultKey]
// 	}
// 	return gormMap[key[0]]
// }

// func main() {
// 	dbInfos := []DbInfo{
// 		{KeyName: "Master", Conn: "string1"},
// 		{KeyName: "Slave", Conn: "string2"},
// 	}
// 	if len(dbInfos) > 0 {
// 		gormDefaultKey = dbInfos[0].KeyName
// 	}
// 	var err error
// 	gormMap, err = NewMulNode(dbInfos, func(a DbInfo) (key string, node *gorm.DB, err error) {
// 		key = a.KeyName
// 		node, err = gorm.Open(sqlite.Open(a.Conn), &gorm.Config{})
// 		return
// 	})
// 	if err != nil {
// 		fmt.Println("err")
// 	}
// 	fmt.Println(gormMap)
// 	fmt.Println(Gorm()) // mean using gormDefaultKey
// 	fmt.Println(Gorm("Slave"))
// }
