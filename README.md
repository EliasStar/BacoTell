<img align="right" width="256" height="256" src="icon.png">

# BacoTell
> Pluggable bot client and associated framework for Discord

## Usage
Download the latest release from [GitHub Releases](https://github.com/EliasStar/BacoTell/releases/latest) for your platform. Then rename the executable and run the following command:
``` sh
bacotell -n <bot_name> -t <bot_token> -p <path_to_plugin_dir>
```


## Developing Plugins
To use this bot framework, you must take the following steps. An example can be seen in `cmd/example_plugin/`.

1. For each application command, message component and/or modal create a struct and implement the respective interface.
2. Either register the implementations using the setters in `bacotell_plugin` or create a custom implementation of the `Plugin` interface.
3. Either call `bacotell_plugin.Run()` or `bacotell_plugin.RunCustom()` depending on your choice in step 2. During development you can use `bacotell_plugin.Debug()` and `bacotell_plugin.DebugCustom()` respectively.

That's it!
For complete API documentation, see [GoDoc](https://pkg.go.dev/github.com/EliasStar/BacoTell).


## Building
Download the latest `protoc` release from [GitHub Releases](https://github.com/protocolbuffers/protobuf/releases/latest) for your platform and add it to your `PATH`. More information can be found on [their GitHub repo](https://github.com/protocolbuffers/protobuf).

Then install the compiler extensions for Go using the following commands. Make sure that your Go tools are also in your `PATH`.
``` sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

If you made changes to the protobuf definitions, build them with the following command. After that you can adjust source code referencing the generated protobuf code.
``` sh
protoc --go_out=. --go_opt=module=github.com/EliasStar/BacoTell --go-grpc_out=. --go-grpc_opt=module=github.com/EliasStar/BacoTell proto/*
```

To compile the BacoTell executable, use this command:
``` sh
go build ./cmd/bacotell/
```


## License
BacoTell - Pluggable bot client and associated framework for Discord<br>
Copyright (C) 2023 Elias*

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
