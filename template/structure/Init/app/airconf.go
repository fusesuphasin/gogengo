package app

import "html/template"

var AirConfTemplate = template.Must(template.New("").Parse(
`# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "." 
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use make as well.
cmd = "go build -o ./bin/doesify ."
# Binary file yields from cmd.
# bin = "bin/doesify"
# Customize binary.
full_bin = "./bin/doesify server"
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = []
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = []
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 2500 # ms
# Stop to run old binary when build errors occur.
stop_on_error = true
# This log file places in your tmp_dir.
log = "air_errors.log"

[log]
# Show log time
time = false

[color]
# Customize ech part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true
`))

