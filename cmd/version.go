package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var bannerBase64 = "ICAgICAgICBDQ0NDQ0NDQ0NDQ0NDICAgICBPT09PT09PT08gICAgIEREREREREREREREREQgICAgICBJSUlJSUlJSUlJTk5OTk5OTk4gICAgICAgIE5OTk5OTk5OICAgICAgICBHR0dHR0dHR0dHR0dHCiAgICAgQ0NDOjo6Ojo6Ojo6Ojo6QyAgIE9POjo6Ojo6Ojo6T08gICBEOjo6Ojo6Ojo6Ojo6REREICAgSTo6Ojo6Ojo6SU46Ojo6Ojo6TiAgICAgICBOOjo6Ojo6TiAgICAgR0dHOjo6Ojo6Ojo6Ojo6RwogICBDQzo6Ojo6Ojo6Ojo6Ojo6OkMgT086Ojo6Ojo6Ojo6Ojo6T08gRDo6Ojo6Ojo6Ojo6Ojo6OkREIEk6Ojo6Ojo6OklOOjo6Ojo6OjpOICAgICAgTjo6Ojo6Ok4gICBHRzo6Ojo6Ojo6Ojo6Ojo6OkcKICBDOjo6OjpDQ0NDQ0NDQzo6OjpDTzo6Ojo6OjpPT086Ojo6Ojo6T0RERDo6Ojo6REREREQ6Ojo6OkRJSTo6Ojo6OklJTjo6Ojo6Ojo6Ok4gICAgIE46Ojo6OjpOICBHOjo6OjpHR0dHR0dHRzo6OjpHCiBDOjo6OjpDICAgICAgIENDQ0NDQ086Ojo6OjpPICAgTzo6Ojo6Ok8gIEQ6Ojo6OkQgICAgRDo6Ojo6RCBJOjo6OkkgIE46Ojo6Ojo6Ojo6TiAgICBOOjo6Ojo6TiBHOjo6OjpHICAgICAgIEdHR0dHRwpDOjo6OjpDICAgICAgICAgICAgICBPOjo6OjpPICAgICBPOjo6OjpPICBEOjo6OjpEICAgICBEOjo6OjpESTo6OjpJICBOOjo6Ojo6Ojo6OjpOICAgTjo6Ojo6Ok5HOjo6OjpHICAgICAgICAgICAgICAKQzo6Ojo6QyAgICAgICAgICAgICAgTzo6Ojo6TyAgICAgTzo6Ojo6TyAgRDo6Ojo6RCAgICAgRDo6Ojo6REk6Ojo6SSAgTjo6Ojo6OjpOOjo6Ok4gIE46Ojo6OjpORzo6Ojo6RyAgICAgICAgICAgICAgCkM6Ojo6OkMgICAgICAgICAgICAgIE86Ojo6Ok8gICAgIE86Ojo6Ok8gIEQ6Ojo6OkQgICAgIEQ6Ojo6OkRJOjo6OkkgIE46Ojo6OjpOIE46Ojo6TiBOOjo6Ojo6Tkc6Ojo6OkcgICAgR0dHR0dHR0dHRwpDOjo6OjpDICAgICAgICAgICAgICBPOjo6OjpPICAgICBPOjo6OjpPICBEOjo6OjpEICAgICBEOjo6OjpESTo6OjpJICBOOjo6Ojo6TiAgTjo6OjpOOjo6Ojo6Ok5HOjo6OjpHICAgIEc6Ojo6Ojo6OkcKQzo6Ojo6QyAgICAgICAgICAgICAgTzo6Ojo6TyAgICAgTzo6Ojo6TyAgRDo6Ojo6RCAgICAgRDo6Ojo6REk6Ojo6SSAgTjo6Ojo6Ok4gICBOOjo6Ojo6Ojo6OjpORzo6Ojo6RyAgICBHR0dHRzo6OjpHCkM6Ojo6OkMgICAgICAgICAgICAgIE86Ojo6Ok8gICAgIE86Ojo6Ok8gIEQ6Ojo6OkQgICAgIEQ6Ojo6OkRJOjo6OkkgIE46Ojo6OjpOICAgIE46Ojo6Ojo6Ojo6Tkc6Ojo6OkcgICAgICAgIEc6Ojo6RwogQzo6Ojo6QyAgICAgICBDQ0NDQ0NPOjo6Ojo6TyAgIE86Ojo6OjpPICBEOjo6OjpEICAgIEQ6Ojo6OkQgSTo6OjpJICBOOjo6Ojo6TiAgICAgTjo6Ojo6Ojo6Ok4gRzo6Ojo6RyAgICAgICBHOjo6OkcKICBDOjo6OjpDQ0NDQ0NDQzo6OjpDTzo6Ojo6OjpPT086Ojo6Ojo6T0RERDo6Ojo6REREREQ6Ojo6OkRJSTo6Ojo6OklJTjo6Ojo6Ok4gICAgICBOOjo6Ojo6OjpOICBHOjo6OjpHR0dHR0dHRzo6OjpHCiAgIENDOjo6Ojo6Ojo6Ojo6Ojo6QyBPTzo6Ojo6Ojo6Ojo6OjpPTyBEOjo6Ojo6Ojo6Ojo6Ojo6REQgSTo6Ojo6Ojo6SU46Ojo6OjpOICAgICAgIE46Ojo6Ojo6TiAgIEdHOjo6Ojo6Ojo6Ojo6Ojo6RwogICAgIENDQzo6Ojo6Ojo6Ojo6OkMgICBPTzo6Ojo6Ojo6Ok9PICAgRDo6Ojo6Ojo6Ojo6OkRERCAgIEk6Ojo6Ojo6OklOOjo6Ojo6TiAgICAgICAgTjo6Ojo6Ok4gICAgIEdHRzo6Ojo6OkdHRzo6OkcKICAgICAgICBDQ0NDQ0NDQ0NDQ0NDICAgICBPT09PT09PT08gICAgIEREREREREREREREREQgICAgICBJSUlJSUlJSUlJTk5OTk5OTk4gICAgICAgICBOTk5OTk5OICAgICAgICBHR0dHR0cgICBHR0dHCg=="

var versionTpl = `%c[%d;%d;%dm%s%c[0m
Name: dora
Version: %s
BuildDate: %s
Arch: %s
CommitID: %s\n
`

var (
	Version   string
	BuildDate string
	CommitID  string
	Arch      string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
		fmt.Printf(versionTpl, 0x1B, 0, 0, 36, banner, 0x1B, Version, BuildDate, runtime.GOOS+"/"+runtime.GOARCH, CommitID)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	rootCmd.AddCommand(versionCmd)
}
