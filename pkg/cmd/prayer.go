package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"

	"github.com/ibihim/pray-go/pkg/aladhan"
	"github.com/ibihim/pray-go/pkg/api"
	"github.com/ibihim/pray-go/pkg/prayer"
	"github.com/ibihim/pray-go/pkg/store"
)

const (
	city   = "city"
	nation = "nation"
	cache  = "cache"
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
			country, err := cmd.Flags().GetString(nation)
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
	nextCmd.Flags().StringP(nation, "n", "Germany", "Country name")
	nextCmd.Flags().BoolP(cache, "s", true, "Cache the prayer times")

	rootCmd.AddCommand(nextCmd)

	return rootCmd
}

func RunNextPrayer(city, country string, cache bool) error {
	today := time.Now()

	// Check for stored prayers. TODO@ibihim make it possible to disable with a flag.
	storedPrayers, err := store.GetAll()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Couldn't get stored prayers: %v", err)
	}

	todaysPrayers, err := storedPrayers.Get(api.ToDay(today))
	if err != nil {
		return fmt.Errorf("Couldn't get today's prayers: %v", err)
	}
	if todaysPrayers != nil {
		prayer, err := prayer.Next(todaysPrayers)
		if err != nil {
			return fmt.Errorf("Couldn't get next prayer: %v", err)
		}

		fmt.Println(prayer)
	}

	monthsPrayers, err := aladhan.GetCalendarByCity(today, city, country, 2)
	if err != nil {
		return err
	}

	// Get the next month's first day's prayer times.
	// This is necessary in case that today is the last day of the month append
	// the next prayer happens tomorrow.
	year, month, _ := today.Date()
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	nextMonthsFirstPrayers, err := aladhan.GetTimingsByCity(nextMonth, city, country, 2)
	if err != nil {
		return err
	}

	klog.V(4).Infof("Timings: %v", nextMonth)

	prayers, err := aladhan.CalendarToPrayerTimes(append(monthsPrayers, nextMonthsFirstPrayers))
	if err != nil {
		return fmt.Errorf("Couldn't convert calendar to prayer times: %v", err)
	}

	prayer, err := prayer.Next(prayers)
	if err != nil {
		return err
	}

	fmt.Printf(prayer)

	return nil
}
