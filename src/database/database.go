package database

import (
	"sync"
	"math/rand"

)

/*
	Simple key pair value DB in memory
 */

type inMemoryDB struct {
	m map[string]int
	lck sync.RWMutex
}
/*
	The first parameter is the value and second is the status (error) if different to "".
	It must be changes to proper error handling

 */
func (db *inMemoryDB) Get (name string) (int,string) {

	db.lck.RLock()
	defer db.lck.RUnlock()

	value, ok := db.m[name]

	if ok != true {
		return  0,"Not Found"
	}

	return value,""
}
/*
	the function may return some execution code
*/

func (db *inMemoryDB) Set (name string)  {

	db.lck.Lock()
	defer db.lck.Unlock()
// 	to every string is added random value. Simulating a business process
	db.m[name]=rand.Int()


}

/*
	make a new map, thus DB
 */
func (db *inMemoryDB) NewDB()  {

	db.m = make(map[string]int)
}

func CreateNewDB() *inMemoryDB {

	return &inMemoryDB{m: make(map[string]int)}
}



