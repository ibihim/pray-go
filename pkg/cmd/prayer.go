package cmd

import (
	"flag"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"

	"github.com/ibihim/pray-go/pkg/client"
	"github.com/ibihim/pray-go/pkg/prayer"
)

const (
	city    = "city"
	country = "country"
	cache   = "cache"
)

func PrayerCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "prayer",
		Short: "A tool to get prayer times",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			flag.CommandLine.VisitAll(func(flag *flag.Flag) {
				klog.V(4).Infof("Flag: --%s=%q", flag.Name, flag.Value)
			})
		},
	}

	// Init klog files
	fs := flag.NewFlagSet("", flag.PanicOnError)
	klog.InitFlags(fs)
	rootCmd.PersistentFlags().AddGoFlagSet(fs)

	nextCmd := &cobra.Command{
		Use:   "next",
		Short: "Get the next prayer time",
		RunE: func(cmd *cobra.Command, args []string) error {
			city, err := cmd.Flags().GetString(city)
			if err != nil {
				return err
			}
			country, err := cmd.Flags().GetString(country)
			if err != nil {
				return err
			}
			cache, err := cmd.Flags().GetBool(cache)
			if err != nil {
				return err
			}

			return RunNextPrayer(city, country, cache)
		},
	}

	nextCmd.Flags().StringP(city, "c", "Berlin", "City name")
	nextCmd.Flags().StringP(country, "C", "Germany", "Country name")
	nextCmd.Flags().BoolP(cache, "s", true, "Cache the prayer times")

	rootCmd.AddCommand(nextCmd)

	return rootCmd
}

func RunNextPrayer(city, country string, cache bool) error {
	today := time.Now()
	todayPrayers, err := client.GetTimingsByCity(today, city, country, 2)
	if err != nil {
		return err
	}

	tomorrow := time.Now().AddDate(0, 0, 1)
	tomorrowPrayers, err := client.GetTimingsByCity(tomorrow, city, country, 2)
	if err != nil {
		return err
	}

	klog.V(4).Infof("Timings: %v", todayPrayers)

	prayer, err := prayer.Next(today, append(todayPrayers, tomorrowPrayers...))
	if err != nil {
		return err
	}

	fmt.Printf(prayer)

	return nil
}
