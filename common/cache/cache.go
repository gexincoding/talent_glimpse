package cache

var (
	RecruitmentCache = "recruitment_"
	LoginCache       = "session_"
)

type Cache struct {
	Type string //缓存类型
}

func (Cache) Put(k string, v interface{}) error {
	return nil
}

func (Cache) Get(k string) {

}

//
//
//func InitCache(){
//	reCach :=
//}
