package cmds

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"pkg.mattglei.ch/terminal/internal/lcp"
	"pkg.mattglei.ch/terminal/internal/output"
	"pkg.mattglei.ch/terminal/internal/util"
)

func workouts(s ssh.Session, styles output.Styles) {
	headers := []string{"", "NAME", "STARTED", "DURATION", "DISTANCE", "TYPE", "AVG. HEART RATE"}

	var data [][]string
	activities, err := lcp.FetchActivities()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load strava activities from lcp"))
		return
	}
	for i, a := range activities.Data {
		switch a.SportType {
		case "GravelRide":
			a.SportType = "Gravel Ride"
		case "MountainBikeRide":
			a.SportType = "Mountain Bike Ride"
		case "WeightTraining":
			a.SportType = "Weight Training"
		case "":
			a.SportType = "Workout"
		}
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				a.Name,
				util.RenderExactFromNow(a.StartDate),
				util.RenderDuration(int(a.MovingTime)),
				fmt.Sprintf("%.2f mi [%.2f km]", a.Distance*0.000621371, a.Distance*0.001),
				a.SportType,
				fmt.Sprintf("%.2f bpm", a.AverageHeartrate),
			},
		)
	}

	table := output.Table(styles).Headers(headers...).Rows(data...).Render()
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().
			Width(lipgloss.Width(table)+10).
			Render("\nOne of my favorite things in the world is staying active and enjoying the outdoors. I grew up in New Hampshire hiking, biking, snowshoeing, and traveling with my family. Out of all of those things I especially love cycling mainly through gravel cycling, road cycling, and mountain biking. Below are some of my most recent Strava activities:"),
	)
	fmt.Fprintln(s)
	fmt.Fprintln(s, table)
	output.LiveFrom(s, styles, table, activities.Updated)
	fmt.Fprintln(s)
}
