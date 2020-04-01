package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "WeChat robot plugin"
	app.Usage = "WeChat robot plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "The key of robot to send message",
			EnvVar: "PLUGIN_KEY",
		},
		cli.StringFlag{
			Name:   "msgtype",
			Usage:  "The type of message, either text, markdown, image, news",
			EnvVar: "PLUGIN_MSGTYPE",
			Value:  "text",
		},
		cli.StringFlag{
			Name:   "content",
			Usage:  "The message content to send",
			EnvVar: "PLUGIN_CONTENT",
		},
		cli.StringSliceFlag{
			Name:   "mentioned_list",
			Usage:  "The mentioned user list, eg: @all,kaynewang",
			EnvVar: "PLUGIN_MENTIONED_LIST",
		},
		cli.StringSliceFlag{
			Name:   "mentioned_mobile_list",
			Usage:  "The mentioned mobile list, eg: @all,kaynewang",
			EnvVar: "PLUGIN_MENTIONED_MOBILE_LIST",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit",
			Usage:  "git commit",
			EnvVar: "DRONE_COMMIT",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:        c.String("build.tag"),
			Number:     c.Int("build.number"),
			Event:      c.String("build.event"),
			Status:     c.String("build.status"),
			Commit:     c.String("commit"),
			CommitLink: c.String("commit.link"),
			Ref:        c.String("commit.ref"),
			Branch:     c.String("commit.branch"),
			Author:     c.String("commit.author"),
			Message:    c.String("commit.message"),
			Link:       c.String("build.link"),
			Started:    c.Int64("build.started"),
			Created:    c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Key:                 c.String("key"),
			MsgType:             c.String("msgtype"),
			Content:             c.String("content"),
			MentionedList:       c.StringSlice("mentioned_list"),
			MentionedMobileList: c.StringSlice("mentioned_mobile_list"),
		},
	}
	return plugin.Exec()
}
