package cas

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CASer interface {
	CasEnforcer() *casbin.Enforcer
	LoadPolicy()
	Enforce(role, sub, obj string) bool
	SavePolicy()
}

type CAS struct {
	needReload bool
	casEnforcer *casbin.Enforcer
}

func (cas *CAS) CasEnforcer() *casbin.Enforcer {
	return cas.casEnforcer
}

func (cas *CAS) LoadPolicy() {
	cas.casEnforcer.LoadPolicy()
}

func (cas *CAS) Enforce(role, sub, object string) bool {
	if cas.needReload {
		cas.LoadPolicy()
	}
	return cas.casEnforcer.Enforce(role, sub, object)
}

func (cas *CAS) SavePolicy() {
	cas.needReload = true
	cas.casEnforcer.SavePolicy()
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(dir, "\\", "/", -1)
}

var Auth CASer

func LoadPolicy() {
	Auth.LoadPolicy()
}

func init() {
	if Auth == nil {
		cas := new(CAS)
		cas.casEnforcer = casbin.NewEnforcer(fmt.Sprintf("%s/resource/rpa_model.conf", getCurrentDirectory()), mongodbadapter.NewAdapter("mongodb:27017/unite"))
		cas.casEnforcer.LoadPolicy()
		//cas.casEnforcer.AddPolicy("admin", "/*", "*")
		//cas.casEnforcer.AddPolicy("anonymous", "/login", "*")
		//cas.SavePolicy()
		Auth = cas
	}
}
