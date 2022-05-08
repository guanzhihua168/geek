package context

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"parent-api-go/global"
	"parent-api-go/pkg/response"

	"reflect"
)

var DefaultAppContext *AppContext

func init() {
	DefaultAppContext = &AppContext{
		Log:     logrus.WithField("defaualt", "default"),
		Context: &gin.Context{},
	}
}

type ReqHandleFunc func(*AppContext)

func Handle(r ReqHandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ac, ok := c.Get(global.AppContext); ok {
			if act, ok := ac.(*AppContext); ok {
				r(act)
			}
		}
	}
}

type AppContext struct {
	*gin.Context
	Log    *logrus.Entry
	AuthId uint32
}

func (c *AppContext) CtxInto(val interface{}) error {

	rtyp := reflect.TypeOf(val)
	rval := reflect.ValueOf(val)

	if rval.Kind() != reflect.Ptr {
		c.Log.Error(val, "not ptr ")
		return errors.New("not ptr ")
	}

	if field, ok := rtyp.Elem().FieldByName("Ctx"); ok {
		if field.Type != reflect.TypeOf(&AppContext{}) {
			c.Log.Error(val, "not AppContext ")
			return errors.New("not AppContext")
		}
		rval.Elem().FieldByName("Ctx").Set(reflect.ValueOf(c))
	} else {
		c.Log.Error(val, "not Ctx ")
		return errors.New("not Ctx ")
	}

	if m := rval.MethodByName("Init"); m.IsValid() {
		m.Call([]reflect.Value{})
	}
	return nil
}

func (c *AppContext) GetCtxInfo(val interface{}) interface{} {

	rtyp := reflect.TypeOf(val)
	rval := reflect.ValueOf(val)

	if rval.Kind() != reflect.Ptr {
		c.Log.Error(val, "not ptr ")
		return errors.New("not ptr ")
	}

	if field, ok := rtyp.Elem().FieldByName("Ctx"); ok {
		if field.Type != reflect.TypeOf(&AppContext{}) {
			c.Log.Error(val, "not AppContext ")
			return errors.New("not AppContext")
		}
		rval.Elem().FieldByName("Ctx").Set(reflect.ValueOf(c))
	} else {
		c.Log.Error(val, "not Ctx ")
		return errors.New("not Ctx ")
	}

	if m := rval.MethodByName("Init"); m.IsValid() {
		m.Call([]reflect.Value{})
	}

	return rval.Interface()
}

func (c *AppContext) Okay(params ...interface{}) {
	c.JSON(response.Okay(params...))
}
