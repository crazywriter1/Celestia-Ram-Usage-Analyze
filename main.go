package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/mem"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	p.Title.Text = "System RAM Usage (6 Hours)"
	p.Y.Label.Text = "RAM usage (%)"

	bars := make(plotter.Values, 0)

	for i := 0; i < 6; i++ {
		v, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		usage := v.UsedPercent
		bars = append(bars, usage)

		fmt.Printf("RAM usage at %d hour(s): %.2f%%\n", i+1, usage)

		if i < 5 {
			time.Sleep(time.Hour)
		}
	}

	barChart, err := plotter.NewBarChart(bars, vg.Points(20))
	if err != nil {
		panic(err)
	}

	barChart.LineStyle.Width = vg.Length(0)
	barChart.Color = plotutil.Color(0)
	barChart.Offset = 0

	p.Add(barChart)
	p.NominalX("1", "2", "3", "4", "5", "6")

	if err := p.Save(10*vg.Inch, 4*vg.Inch, "ram_usage_chart.png"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
