//go:build mini
// +build mini

package GoScriptCode

import (
	"github.com/qtgolang/SunnyNet/src/Interface"
)

type GoScriptTypeHTTP func(Interface.ConnHTTPScriptCall)
type GoScriptTypeWS func(Interface.ConnWebSocketScriptCall)
type GoScriptTypeTCP func(Interface.ConnTCPScriptCall)
type GoScriptTypeUDP func(Interface.ConnUDPScriptCall)

type LogFuncInterface func(info ...any)
type SaveFuncInterface func(code []byte)

func RunCode(UserScriptCode []byte, log LogFuncInterface) (resError string, h GoScriptTypeHTTP, w GoScriptTypeWS, t GoScriptTypeTCP, u GoScriptTypeUDP) {
	return "DLL不支持脚本代码", nil, nil, nil, nil
}
