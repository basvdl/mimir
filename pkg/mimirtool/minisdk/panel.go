// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/grafana-tools/sdk/blob/master/panel.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: 2016 Alexander I.Grafov <grafov@gmail.com>.
// Provenance-includes-copyright: 2016-2019 The Grafana SDK authors

package minisdk

import (
	"encoding/json"
	"fmt"
)

// Each panel may be one of these types.
const (
	CustomType panelType = iota
	DashlistType
	GraphType
	TableType
	TextType
	PluginlistType
	AlertlistType
	SinglestatType
	StatType
	RowType
	BarGaugeType
	HeatmapType
	TimeseriesType
	GaugeType
	PieChartType
	BarChartType
	CanvasType
	XyChartType
	StateTimelineType
	StatusHistoryType
	HistogramType
	GeomapType
)

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		CommonPanel
		// Should be initialized only one type of panels.
		// OfType field defines which of types below will be used.
		*GraphPanel
		*TablePanel
		*TextPanel
		*SinglestatPanel
		*StatPanel
		*DashlistPanel
		*PluginlistPanel
		*RowPanel
		*AlertlistPanel
		*BarGaugePanel
		*HeatmapPanel
		*TimeseriesPanel
		*GaugePanel
		*PieChartPanel
		*BarChartPanel
		*CanvasPanel
		*XyChartPanel
		*StateTimelinePanel
		*StatusHistoryPanel
		*HistogramPanel
		*GeomapPanel
		*CustomPanel
	}
	panelType   int8
	CommonPanel struct {
		Datasource *DatasourceRef `json:"datasource,omitempty"` // metrics
		ID         uint           `json:"id"`
		OfType     panelType      `json:"-"`     // it required for defining type of the panel
		Title      string         `json:"title"` // general
		Type       string         `json:"type"`
	}
	GraphPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	TablePanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	TextPanel       struct{}
	SinglestatPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	StatPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	DashlistPanel   struct{}
	PluginlistPanel struct{}
	AlertlistPanel  struct{}
	BarGaugePanel   struct {
		Targets []Target `json:"targets,omitempty"`
	}
	RowPanel struct {
		Panels []Panel `json:"panels"`
	}
	HeatmapPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	TimeseriesPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	GaugePanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	PieChartPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	BarChartPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	CanvasPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	XyChartPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	StateTimelinePanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	StatusHistoryPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	HistogramPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	GeomapPanel struct {
		Targets []Target `json:"targets,omitempty"`
	}
	CustomPanel map[string]interface{}
)

// for an any panel
type Target struct {
	Datasource *DatasourceRef `json:"datasource,omitempty"`
	Expr       string         `json:"expr,omitempty"`
}

// GetTargets is iterate over all panel targets. It just returns nil if
// no targets defined for panel of concrete type.
func (p *Panel) GetTargets() *[]Target {
	switch p.OfType {
	case GraphType:
		return &p.GraphPanel.Targets
	case SinglestatType:
		return &p.SinglestatPanel.Targets
	case StatType:
		return &p.StatPanel.Targets
	case TableType:
		return &p.TablePanel.Targets
	case BarGaugeType:
		return &p.BarGaugePanel.Targets
	case HeatmapType:
		return &p.HeatmapPanel.Targets
	case TimeseriesType:
		return &p.TimeseriesPanel.Targets
	case GaugeType:
		return &p.GaugePanel.Targets
	case PieChartType:
		return &p.PieChartPanel.Targets
	case BarChartType:
		return &p.BarChartPanel.Targets
	case CanvasType:
		return &p.CanvasPanel.Targets
	case XyChartType:
		return &p.XyChartPanel.Targets
	case StateTimelineType:
		return &p.StateTimelinePanel.Targets
	case StatusHistoryType:
		return &p.StatusHistoryPanel.Targets
	case HistogramType:
		return &p.HistogramPanel.Targets
	case GeomapType:
		return &p.GeomapPanel.Targets
	default:
		return nil
	}
}

type probePanel struct {
	CommonPanel
	//	json.RawMessage
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var probe probePanel
	if err = json.Unmarshal(b, &probe); err != nil {
		return err
	}

	p.CommonPanel = probe.CommonPanel
	switch probe.Type {
	case "graph":
		var graph GraphPanel
		p.OfType = GraphType
		if err = json.Unmarshal(b, &graph); err == nil {
			p.GraphPanel = &graph
		}
	case "table":
		var table TablePanel
		p.OfType = TableType
		if err = json.Unmarshal(b, &table); err == nil {
			p.TablePanel = &table
		}
	case "text":
		var text TextPanel
		p.OfType = TextType
		if err = json.Unmarshal(b, &text); err == nil {
			p.TextPanel = &text
		}
	case "singlestat":
		var singlestat SinglestatPanel
		p.OfType = SinglestatType
		if err = json.Unmarshal(b, &singlestat); err == nil {
			p.SinglestatPanel = &singlestat
		}
	case "stat":
		var stat StatPanel
		p.OfType = StatType
		if err = json.Unmarshal(b, &stat); err == nil {
			p.StatPanel = &stat
		}
	case "dashlist":
		var dashlist DashlistPanel
		p.OfType = DashlistType
		if err = json.Unmarshal(b, &dashlist); err == nil {
			p.DashlistPanel = &dashlist
		}
	case "bargauge":
		var bargauge BarGaugePanel
		p.OfType = BarGaugeType
		if err = json.Unmarshal(b, &bargauge); err == nil {
			p.BarGaugePanel = &bargauge
		}
	case "heatmap":
		var heatmap HeatmapPanel
		p.OfType = HeatmapType
		if err = json.Unmarshal(b, &heatmap); err == nil {
			p.HeatmapPanel = &heatmap
		}
	case "timeseries":
		var timeseries TimeseriesPanel
		p.OfType = TimeseriesType
		if err = json.Unmarshal(b, &timeseries); err == nil {
			p.TimeseriesPanel = &timeseries
		}
	case "row":
		var rowpanel RowPanel
		p.OfType = RowType
		if err = json.Unmarshal(b, &rowpanel); err == nil {
			p.RowPanel = &rowpanel
		}
	case "gauge":
		var gauge GaugePanel
		p.OfType = GaugeType
		if err = json.Unmarshal(b, &gauge); err == nil {
			p.GaugePanel = &gauge
		}
	case "piechart":
		var piechart PieChartPanel
		p.OfType = PieChartType
		if err = json.Unmarshal(b, &piechart); err == nil {
			p.PieChartPanel = &piechart
		}
	case "barchart":
		var barchart BarChartPanel
		p.OfType = BarChartType
		if err = json.Unmarshal(b, &barchart); err == nil {
			p.BarChartPanel = &barchart
		}
	case "canvas":
		var canvas CanvasPanel
		p.OfType = CanvasType
		if err = json.Unmarshal(b, &canvas); err == nil {
			p.CanvasPanel = &canvas
		}
	case "xychart":
		var xychart XyChartPanel
		p.OfType = XyChartType
		if err = json.Unmarshal(b, &xychart); err == nil {
			p.XyChartPanel = &xychart
		}
	case "state-timeline":
		var statetimeline StateTimelinePanel
		p.OfType = StateTimelineType
		if err = json.Unmarshal(b, &statetimeline); err == nil {
			p.StateTimelinePanel = &statetimeline
		}
	case "status-history":
		var statushistory StatusHistoryPanel
		p.OfType = StatusHistoryType
		if err = json.Unmarshal(b, &statushistory); err == nil {
			p.StatusHistoryPanel = &statushistory
		}
	case "histogram":
		var histogram HistogramPanel
		p.OfType = HistogramType
		if err = json.Unmarshal(b, &histogram); err == nil {
			p.HistogramPanel = &histogram
		}
	case "geomap":
		var geomap GeomapPanel
		p.OfType = GeomapType
		if err = json.Unmarshal(b, &geomap); err == nil {
			p.GeomapPanel = &geomap
		}
	default:
		var custom = make(CustomPanel)
		p.OfType = CustomType
		if err = json.Unmarshal(b, &custom); err == nil {
			p.CustomPanel = &custom
		}
	}

	if err != nil && (probe.Title != "" || probe.Type != "") {
		err = fmt.Errorf("%w (panel %q of type %q)", err, probe.Title, probe.Type)
	}

	return err
}
