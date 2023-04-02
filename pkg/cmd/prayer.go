package cmd

import (
	"flag"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"

	"github.com/ibihim/pray-go/pkg/aladhan"
	"github.com/ibihim/pray-go/pkg/api"
	"github.com/ibihim/pray-go/pkg/prayer"
	"github.com/ibihim/pray-go/pkg/store"
)

const (
	city      = "city"
	nation    = "nation"
	nocache   = "nocache"
	listFiles = "list-files"
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
			nocache, err := cmd.Flags().GetBool(nocache)
			if err != nil {
				return err
			}
			listFiles, err := cmd.Flags().GetBool(listFiles)
			if err != nil {
				return err
			}
			if listFiles {
				fmt.Println(store.GetCacheFilePath())
				return nil
			}

			return RunNextPrayer(city, country, !nocache)
		},
	}

	nextCmd.Flags().StringP(city, "c", "Berlin", "City name")
	nextCmd.Flags().StringP(nation, "n", "Germany", "Country name")
	nextCmd.Flags().BoolP(nocache, "s", false, "Do not cache the prayer times")
	nextCmd.Flags().BoolP(listFiles, "l", false, "List all files")

	rootCmd.AddCommand(nextCmd)

	return rootCmd
}

func RunNextPrayer(city, country string, cache bool) error {
	today := time.Now()

	klog.V(4).Infof("Getting prayer for %s, %s with cache: ", city, country, cache)

	if cache {
		nextPrayer, err := getPrayerFromFile()
		if err == nil {
			klog.V(4).Info("Got prayer from file")

			fmt.Printf(nextPrayer.ClockString())
			return nil
		}

		klog.V(4).Infof("Failed to get prayer from file: %v", err)
	}

	prayers, err := aladhan.GetPrayers(today, city, country, 2)
	if err != nil {
		return err
	}

	klog.V(4).Infof("Got prayers from API: %v", prayers)

	nextPrayers, err := prayer.NextPrayers(prayers)
	if err != nil {
		return err
	}

	klog.V(4).Infof("Got next prayers: %v", nextPrayers)

	fmt.Printf(nextPrayers[0].ClockString())

	if !cache {
		return nil
	}

	klog.V(4).Infof("Storing prayers: %v", prayers)

	if err := store.Store(prayers); err != nil {
		return fmt.Errorf("failed to store prayers: %w", err)
	}

	return nil
}

func getPrayerFromFile() (*api.Prayer, error) {
	prayers, err := store.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get stored prayers: %w", err)
	}

	if len(prayers) == 0 {
		return nil, fmt.Errorf("no stored prayers found")
	}

	return prayer.NextPrayer(prayers)
}
