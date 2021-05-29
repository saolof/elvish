package file

import (
	"os"
	"testing"

	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/errs"
	. "src.elv.sh/pkg/eval/evaltest"
	"src.elv.sh/pkg/testutil"
)

// A number that exceeds the range of int64
const z = "100000000000000000000"

func TestFile(t *testing.T) {
	setup := func(ev *eval.Evaler) {
		ev.AddGlobal(eval.NsBuilder{}.AddNs("file", Ns).Ns())
	}
	_, cleanup := testutil.InTestDir()
	defer cleanup()
	TestWithSetup(t, setup,
		That(
			"echo haha > out3", "f = (file:open out3)",
			"slurp < $f", "file:close $f").Puts("haha\n"),

		That(`p = (file:pipe)`, `echo haha > $p `, `pwclose $p`,
			`slurp < $p`, `prclose $p`).Puts("haha\n"),

		That(`p = (file:pipe)`, `echo Zeppelin > $p`, `file:pwclose $p`,
			`echo Sabbath > $p`, `slurp < $p`, `file:prclose $p`).Puts("Zeppelin\n"),

		That(`p = (file:pipe)`, `echo Legolas > $p`, `file:prclose $p`,
			`slurp < $p`).Throws(AnyError),

		// Side effect checked below
		That("echo > file100", "file:truncate file100 100").DoesNothing(),

		// Should also test the case where the argument doesn't fit in an int
		// but does in a *big.Int, but this could consume too much disk

		That("file:truncate bad -1").Throws(errs.OutOfRange{
			What:     "size argument to file:truncate",
			ValidLow: "0", ValidHigh: "2^64-1", Actual: "-1",
		}),

		That("file:truncate bad "+z).Throws(errs.OutOfRange{
			What:     "size argument to file:truncate",
			ValidLow: "0", ValidHigh: "2^64-1", Actual: z,
		}),

		That("file:truncate bad 1.5").Throws(errs.BadValue{
			What:  "size argument to file:truncate",
			Valid: "integer", Actual: "non-integer",
		}),
	)

	fi, err := os.Stat("file100")
	if err != nil {
		t.Errorf("stat file100: %v", err)
	}
	if size := fi.Size(); size != 100 {
		t.Errorf("got file100 size %v, want 100", size)
	}
}
