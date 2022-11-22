# Construct a callable value for the external program `$program`. Example:
#
# ```elvish-transcript
# ~> var x = (external man)
# ~> $x ls # opens the manpage for ls
# ```
#
# @cf has-external search-external
fn external {|program| }

# Test whether `$command` names a valid external command. Examples (your output
# might differ):
#
# ```elvish-transcript
# ~> has-external cat
# ▶ $true
# ~> has-external lalala
# ▶ $false
# ```
#
# @cf external search-external
fn has-external {|command| }

# Output the full path of the external `$command`. Throws an exception when not
# found. Example (your output might vary):
#
# ```elvish-transcript
# ~> search-external cat
# ▶ /bin/cat
# ```
#
# @cf external has-external
fn search-external {|command| }

# Exit the Elvish process with `$status` (defaulting to 0).
fn exit {|status?| }
