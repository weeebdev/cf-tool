package cmd

import (
	"fmt"
	"os"
	"time"

	"cf-tool/client"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// CList command
func CList() (err error) {
	cln := client.Instance
	info := Args.Info
	contests, err := cln.CList(info)
	if err != nil {
		return err
	}

	color.Cyan("Total contests: %v", len(contests))
	color.Cyan("Showing top 10")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Name", "Phase", "Duration", "Date"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	for i, cont := range contests {
		if i > 9 {
			break
		}
		duration := time.Duration(cont.DurationSeconds) * time.Second
		date := time.Now().Add(duration).Format("2006-01-02 15:04:05")
		table.Append([]string{
			fmt.Sprint(cont.ID),
			cont.Name,
			cont.Phase,
			// convert seconds to minutes
			fmt.Sprintf("%v min", cont.DurationSeconds / 60),
			// convert seconds (int) to date
			date,
		})
	}
	table.Render()

	return
}
