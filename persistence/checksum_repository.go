package persistence

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/cloudsonic/sonic-server/scanner"
	"github.com/siddontang/ledisdb/ledis"
)

var (
	checkSumKeyName = []byte("checksums")
)

type checkSumRepository struct {
	data map[string]string
}

func NewCheckSumRepository() scanner.CheckSumRepository {
	r := &checkSumRepository{}
	r.loadData()
	return r
}

func (r *checkSumRepository) loadData() {
	r.data = make(map[string]string)

	pairs, err := Db().HGetAll(checkSumKeyName)
	if err != nil {
		beego.Error("Error loading CheckSums:", err)
	}
	for _, p := range pairs {
		r.data[string(p.Field)] = string(p.Value)
	}
	beego.Debug("Loaded", len(r.data), "checksums")
}

func (r *checkSumRepository) Put(id, sum string) error {
	if id == "" {
		return errors.New("Id is required")
	}
	_, err := Db().HSet(checkSumKeyName, []byte(id), []byte(sum))
	return err
}

func (r *checkSumRepository) Get(id string) (string, error) {
	return r.data[id], nil
}

func (r *checkSumRepository) SetData(newSums map[string]string) error {
	Db().HClear(checkSumKeyName)
	pairs := make([]ledis.FVPair, len(newSums))
	r.data = make(map[string]string)
	i := 0
	for id, sum := range newSums {
		p := ledis.FVPair{Field: []byte(id), Value: []byte(sum)}
		pairs[i] = p
		r.data[id] = sum
		i++
	}
	return Db().HMset(checkSumKeyName, pairs...)
}
