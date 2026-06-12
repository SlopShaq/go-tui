package tuigen

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"regexp"
	"testing"
)

// mountIndexRe extracts the index expression from generated app.Mount /
// app.MountPersistent calls.
var mountIndexRe = regexp.MustCompile(`app\.Mount(?:Persistent)?\(c, (.+), func\(\) tui\.Component \{`)

// evalMountKey evaluates a generated mount index expression after substituting
// concrete iteration values for loop index variables.
func evalMountKey(t *testing.T, expr string, vars map[string]int) int64 {
	t.Helper()
	for name, val := range vars {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(name) + `\b`)
		expr = re.ReplaceAllString(expr, fmt.Sprintf("%d", val))
	}
	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, expr)
	if err != nil {
		t.Fatalf("evaluating mount key %q: %v", expr, err)
	}
	v, ok := constant.Int64Val(tv.Value)
	if !ok {
		t.Fatalf("mount key %q did not evaluate to an integer constant", expr)
	}
	return v
}

// TestGenerator_MountKeysDoNotCollide reproduces issue #88: a component
// mounted inside a for loop must never share a runtime mount key with a
// component mounted at a different call site. With the colliding keys, the
// mount cache returns the wrong component type at runtime (e.g. a Markdown
// where a TextArea was requested).
func TestGenerator_MountKeysDoNotCollide(t *testing.T) {
	type tc struct {
		input    string
		loopVars []string // loop index variables expected in generated keys
		iters    int      // iterations to simulate per loop variable
		wantKeys int      // expected number of mount call sites
	}

	tests := map[string]tc{
		"loop component followed by standalone component": {
			input: `package x

type app struct{}

templ (c *app) Render() {
	<div>
		for i, item := range c.items {
			<markdown source={item} />
		}
		<textarea onSubmit={c.submit} />
	</div>
}`,
			loopVars: []string{"i"},
			iters:    3,
			wantKeys: 2,
		},
		"loop with discarded index uses synthetic variable": {
			input: `package x

type app struct{}

templ (c *app) Render() {
	<div>
		for _, item := range c.items {
			<markdown source={item} />
		}
		<textarea onSubmit={c.submit} />
	</div>
}`,
			loopVars: []string{"__idx_0"},
			iters:    3,
			wantKeys: 2,
		},
		"loop struct component call followed by standalone component": {
			input: `package x

type app struct{}

templ (c *app) Render() {
	<div>
		for i, item := range c.items {
			@Widget(item)
		}
		<textarea onSubmit={c.submit} />
	</div>
}`,
			loopVars: []string{"i"},
			iters:    3,
			wantKeys: 2,
		},
		"nested loop component followed by standalone component": {
			input: `package x

type app struct{}

templ (c *app) Render() {
	<div>
		for i, row := range c.rows {
			for j, item := range row {
				<markdown source={item} />
			}
		}
		<textarea onSubmit={c.submit} />
	</div>
}`,
			loopVars: []string{"i", "j"},
			iters:    3,
			wantKeys: 2,
		},
		"two sibling loops with components": {
			input: `package x

type app struct{}

templ (c *app) Render() {
	<div>
		for i, item := range c.items {
			<markdown source={item} />
		}
		for j, other := range c.others {
			<markdown source={other} />
		}
	</div>
}`,
			loopVars: []string{"i", "j"},
			iters:    3,
			wantKeys: 2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := parseAndGenerateSkipImports("test.gsx", tt.input)
			if err != nil {
				t.Fatalf("generation failed: %v", err)
			}
			code := string(output)

			matches := mountIndexRe.FindAllStringSubmatch(code, -1)
			if len(matches) != tt.wantKeys {
				t.Fatalf("expected %d mount calls, found %d\nGot:\n%s", tt.wantKeys, len(matches), code)
			}

			// For each call site, expand the key expression over simulated
			// loop iterations and record every runtime key it can produce.
			seen := map[int64]int{} // runtime key -> call site index
			for site, m := range matches {
				expr := m[1]

				// Find which loop variables this expression references.
				var refs []string
				for _, v := range tt.loopVars {
					if regexp.MustCompile(`\b` + regexp.QuoteMeta(v) + `\b`).MatchString(expr) {
						refs = append(refs, v)
					}
				}

				// Enumerate all combinations of iteration values for the
				// referenced loop variables.
				combos := [][]int{{}}
				for range refs {
					var next [][]int
					for _, c := range combos {
						for it := 0; it < tt.iters; it++ {
							next = append(next, append(append([]int{}, c...), it))
						}
					}
					combos = next
				}

				for _, combo := range combos {
					vars := map[string]int{}
					for vi, v := range refs {
						vars[v] = combo[vi]
					}
					key := evalMountKey(t, expr, vars)
					if prev, ok := seen[key]; ok && prev != site {
						t.Errorf("mount key collision: call site %d expr %q produces key %d already produced by call site %d (vars %v)",
							site, expr, key, prev, vars)
					}
					seen[key] = site
				}
			}
		})
	}
}
