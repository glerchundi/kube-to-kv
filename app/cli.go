package main

import (
	"os"
	"reflect"

	"github.com/codegangsta/cli"
    "github.com/glerchundi/kube2kv"
	"github.com/glerchundi/kube2kv/config"
)

var (
	globalConfig *config.GlobalConfig = nil
)

func handleGlobalAndTemplateOpts(globalFlags []Flag, c *cli.Context) (err error) {
	globalConfig = config.NewGlobalConfig()
	overwriteWithCliFlags(globalFlags, c, true, globalConfig)
	return
}

func handleBackendOptsAndRun(backendFlags []Flag, c *cli.Context) {
	// set requested command/backend
	backend := c.Command.Name

	// backend config
	backendConfig := config.NewBackendConfig(backend)
	overwriteWithCliFlags(backendFlags, c, false, backendConfig)

	// run!
	kube2kv.Run(globalConfig, backendConfig)
}

func main() {
	// lookup flags
	globalFlags := getFlagsFromType(reflect.TypeOf((*config.GlobalConfig)(nil)).Elem())
	consulFlags := getFlagsFromType(reflect.TypeOf((*config.ConsulBackendConfig)(nil)).Elem())
	etcdFlags := getFlagsFromType(reflect.TypeOf((*config.EtcdBackendConfig)(nil)).Elem())

	// app
	app := cli.NewApp()
	app.Name = "kube2kv"
	app.Version = "0.1.0"
	app.Usage = "listen to kubernetes service events and save them inside a kv store (etcd,consul)."
	app.Flags = getCliFlags(globalFlags)
	app.Before = func(c *cli.Context) error {
		return handleGlobalAndTemplateOpts(globalFlags, c)
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "consul",
			Flags:  getCliFlags(consulFlags),
			Action: func(c *cli.Context) { handleBackendOptsAndRun(consulFlags, c) },
		},
		cli.Command{
			Name:   "etcd",
			Flags:  getCliFlags(etcdFlags),
			Action: func(c *cli.Context) { handleBackendOptsAndRun(etcdFlags, c) },
		},
	}
	app.Run(os.Args)
}
