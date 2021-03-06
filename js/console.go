/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

 package js

 import (
	 "context"
	 "reflect"
	 "strconv"
 
	 "github.com/dop251/goja"
	 log "github.com/sirupsen/logrus"
 )
 
 var byteType = reflect.TypeOf((*[]byte)(nil))
 
 type Console struct {
	 Logger *log.Logger
 }
 
 func NewConsole() *Console {
	 return &Console{log.StandardLogger()}
 }
 
 func (c Console) log(ctx *context.Context, level log.Level, msgobj goja.Value, args ...goja.Value) {
	 if ctx != nil && *ctx != nil {
		 select {
		 case <-(*ctx).Done():
			 return
		 default:
		 }
	 }
 
	 fields := make(log.Fields)
	 for i, arg := range args {
		 if arg.ExportType() == byteType.Elem() {
			 fields[strconv.Itoa(i)] = string(arg.Export().([]byte))
		 } else {
			 fields[strconv.Itoa(i)] = arg.String()
		 }
	 }
	 msg := msgobj.String()
	 if msgobj.ExportType() == byteType.Elem() {
		 msg = string(msgobj.Export().([]byte))
	 }
	 e := c.Logger.WithFields(fields)
	 switch level {
	 case log.DebugLevel:
		 e.Debug(msg)
	 case log.InfoLevel:
		 e.Info(msg)
	 case log.WarnLevel:
		 e.Warn(msg)
	 case log.ErrorLevel:
		 e.Error(msg)
	 }
 }
 
 func (c Console) Log(ctx *context.Context, msg goja.Value, args ...goja.Value) {
	 c.Info(ctx, msg, args...)
 }
 
 func (c Console) Debug(ctx *context.Context, msg goja.Value, args ...goja.Value) {
	 c.log(ctx, log.DebugLevel, msg, args...)
 }
 
 func (c Console) Info(ctx *context.Context, msg goja.Value, args ...goja.Value) {
	 c.log(ctx, log.InfoLevel, msg, args...)
 }
 
 func (c Console) Warn(ctx *context.Context, msg goja.Value, args ...goja.Value) {
	 c.log(ctx, log.WarnLevel, msg, args...)
 }
 
 func (c Console) Error(ctx *context.Context, msg goja.Value, args ...goja.Value) {
	 c.log(ctx, log.ErrorLevel, msg, args...)
 }
 