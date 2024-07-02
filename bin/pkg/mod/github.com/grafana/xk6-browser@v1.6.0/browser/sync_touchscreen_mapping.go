package browser

import (
	"github.com/grafana/sobek"

	"github.com/grafana/xk6-browser/common"
	"github.com/grafana/xk6-browser/k6ext"
)

// syncMapTouchscreen is like mapTouchscreen but returns synchronous functions.
func syncMapTouchscreen(vu moduleVU, ts *common.Touchscreen) mapping {
	return mapping{
		"tap": func(x float64, y float64) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (result any, reason error) {
				return nil, ts.Tap(x, y) //nolint:wrapcheck
			})
		},
	}
}
