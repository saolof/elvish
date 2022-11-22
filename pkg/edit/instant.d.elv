# Binding for the instant mode.
#doc:show-unstable
var -instant:binding

# Starts the instant mode. In instant mode, any text entered at the command
# line is evaluated immediately, with the output displayed.
#
# **WARNING**: Beware of unintended consequences when using destructive
# commands. For example, if you type `sudo rm -rf /tmp/*` in the instant mode,
# Elvish will attempt to evaluate `sudo rm -rf /` when you typed that far.
#doc:show-unstable
fn -instant:start { }
