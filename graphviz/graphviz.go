// Package graphviz provides some dot(1) utilities.
package graphviz

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
)

// lazy.
var path = "/tmp/graph.html"

// html template.
var html = `<html>
<body>
  <script src="https://cdn.rawgit.com/anvaka/panzoom/v4.0.0/dist/panzoom.min.js"></script>
    <div id="graph">{graph}</div>
    <script>
      panzoom(document.querySelector('#graph'))
    </script>
</body>
</html>`

// render template.
func render(graph []byte) []byte {
	return bytes.Replace([]byte(html), []byte("{graph}"), graph, -1)
}

// OpenDot opens the given reader as zoomable SVG in the browser.
func OpenDot(r io.Reader) error {
	cmd := exec.Command("dot", "-Tsvg")
	cmd.Stdin = r

	b, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "executing")
	}

	err = ioutil.WriteFile(path, render(b), 0755)
	if err != nil {
		return errors.Wrap(err, "writing")
	}

	return browser.OpenURL("file://" + path)
}

// Must helper.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
