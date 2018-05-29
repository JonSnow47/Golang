/*
 * Revision History:
 *     Initial: 2018/05/02        Chen Yanchen
 */

package main

import (
	"gopkg.in/go-playground/validator.v9"
)

type Context struct {
	Validate *validator.Validate
}

func (c *Context) Validator(s interface{}) error {
	return c.Validate.Struct(s)
}
