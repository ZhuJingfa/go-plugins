# Disable RPC Plugin

This plugin returns a 403 for /rpc. Nothing more.

## Usage

Register the plugin before building Micro

```
package main

import (
	"micro/micro/plugin"
	rpc "micro/go-plugins/micro/disable_rpc"
)

func init() {
	plugin.Register(rpc.NewPlugin())
}
```

### Scoped to API

If you like to only apply the plugin for a specific component you can register it with that specifically. 
For example, below you'll see the plugin registered with the API.

```
package main

import (
	"micro/micro/api"
	rpc "micro/go-plugins/micro/disable_rpc"
)

func init() {
	api.Register(rpc.NewPlugin())
}
```
