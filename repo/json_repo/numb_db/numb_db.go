package numb_db

/* A simple JSON Database
 * Scheme:
 * {
 *    "count": 0,  -> PRIMARY KEY counter
 *    [
 *       // DATA
 *    ]
 * }
 */

type NumbDB struct {
	path string
}

func NewNumbDB(path string) *NumbDB {
	ret := new(NumbDB)
	ret.path = path
	return ret
}

func (self *NumbDB) Query(query string, params ...any) (any, error) {
	return nil, nil
}

func (self *NumbDB) Exec(query string, params ...any) error {
	return nil
}
