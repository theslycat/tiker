package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
)

// Config represents the top-level YAML structure
type Config struct {
	Name  string `mapstructure:"Name"`
	Tasks []Task `mapstructure:"Tasks"`
}

// Task represents each task in the Tasks array
type Task struct {
	ID      int    `mapstructure:"id"`
	Content string `mapstructure:"content"`
	Start   string `mapstructure:"start"`
	End     string `mapstructure:"end"`
}

func calculateDuration(start string, end string) (int, error) {
	var startH, startM, endH, endM int
	var duration int
	splittedStart := strings.Split(start, ":")
	splittedEnd := strings.Split(end, ":")
	startH, err := strconv.Atoi(splittedStart[0])
	if err != nil {
		return -1, err
	}
	startM, err = strconv.Atoi(splittedStart[1])
	if err != nil {
		return -1, err
	}

	endH, err = strconv.Atoi(splittedEnd[0])
	endM, err = strconv.Atoi(splittedEnd[1])

	duration = 3600*(endH-startH) + 60*(endM-startM)
	return duration, nil
}
func tickSingleTask(t Task) error {
	fmt.Println("NOW:", t.Content)
	// calculate duration
	duration, err := calculateDuration(t.Start, t.End)
	if err != nil {
		return err
	}
	//log.Println(duration)
	for i := duration; i >= 0; i-- {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d:%d:%d Left\r", i/3600, (i-(i/3600)*3600)/60, (i-(i/3600)*3600)-((i-(i/3600)*3600)/60)*60)
	}
	return nil
}
func main() {
	logger := log.Default()
	TaskList := []Task{}
	// INIT CONFIG
	viper.SetDefault("tiker-style", "normal")

	// END
	cmd := cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "load-from-file",
				Value: "",
				Usage: "where to load ur schedule from",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			/*
				If specified a schedule file to load from, then load it and parse it.
			*/
			if c.String("load-from-file") != "" {
				filePath := c.String("load-from-file")
				// load the file
				logger.Println("loading from " + filePath)
				viper.SetConfigFile(filePath)
				viper.SetConfigType("yaml")
				viper.ReadInConfig()
				// parse the file
				logger.Println("parsing schedule file " + string(viper.GetString("Name")))
				var config Config
				viper.Unmarshal(&config)
				//fmt.Println(config.Tasks[0].ID)
				for _, task := range config.Tasks {
					TaskList = append(TaskList, task)
				}
			}
			// clean the screen
			screen.Clear()
			screen.MoveTopLeft()
			/*
				tick a task
				then calculate rest time (only if next task exists)
			*/
			for n, task := range TaskList {
				screen.Clear()
				screen.MoveTopLeft()
				if err := tickSingleTask(task); err != nil {
					logger.Println("ERROR:" + err.Error())
				}
				if n != len(TaskList)-1 {
					next := TaskList[n+1]
					// culculate rest time
					restDuration, err := calculateDuration(task.End, next.Start)
					if err != nil {
						return err
					}
					if restDuration == 0 {
						continue
					}
					screen.Clear()
					screen.MoveTopLeft()
					fmt.Println("BE BACK SOOOOON:")
					for i := restDuration; i > 1; i-- {
						time.Sleep(time.Second * 1)
						fmt.Printf("%d:%d:%d\r", i/3600, (i-(i/3600)*3600)/60, (i-(i/3600)*3600)-((i-(i/3600)*3600)/60)*60)
					}
				}
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
