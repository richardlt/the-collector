package main

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/pkg/errors"
	"github.com/richardlt/the-collector/server"
	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

func main() {
	app := cli.NewApp()
	app.Name = "the-collector"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "config.yml",
			EnvVar: "THE_COLLECTOR_CONFIG",
		},
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "app-uri",
			Value:  "https://localhost:8081",
			EnvVar: "THE_COLLECTOR_APP_URI",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "jwt-secret",
			Value:  "secret-key",
			EnvVar: "THE_COLLECTOR_JWT_SECRET",
		}),
		altsrc.NewBoolFlag(cli.BoolFlag{
			Name:   "debug",
			EnvVar: "THE_COLLECTOR_DEBUG",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "database-uri",
			Value:  "localhost:27017",
			EnvVar: "THE_COLLECTOR_DATABASE_URI",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "database-name",
			Value:  "the-collector",
			EnvVar: "THE_COLLECTOR_DATABASE_NAME",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "facebook-app-id",
			EnvVar: "THE_COLLECTOR_FACEBOOK_APP_ID",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "facebook-app-secret",
			EnvVar: "THE_COLLECTOR_FACEBOOK_APP_SECRET",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "minio-uri",
			Value:  "localhost:9000",
			EnvVar: "THE_COLLECTOR_MINIO_URI",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "minio-access-key",
			EnvVar: "THE_COLLECTOR_MINIO_ACCESS_KEY",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "minio-secret-key",
			EnvVar: "THE_COLLECTOR_MINIO_SECRET_KEY",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "minio-bucket",
			Value:  "files",
			EnvVar: "THE_COLLECTOR_MINIO_BUCKET",
		}),
		altsrc.NewBoolFlag(cli.BoolFlag{
			Name:   "minio-ssl",
			EnvVar: "THE_COLLECTOR_MINIO_SSL",
		}),
		altsrc.NewBoolFlag(cli.BoolFlag{
			Name:   "sentry",
			EnvVar: "THE_COLLECTOR_SENTRY",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:   "sentry-dsn",
			Value:  "",
			EnvVar: "THE_COLLECTOR_SENTRY_DSN",
		}),
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the app",
			Flags: flags,
			Before: altsrc.InitInputSourceWithContext(flags,
				func(c *cli.Context) (altsrc.InputSourceContext, error) {
					i, err := altsrc.NewYamlSourceFromFlagFunc("config")(c)
					if err == nil {
						return i, err
					}

					return &altsrc.MapInputSource{}, nil
				},
			),
			Action: func(c *cli.Context) {
				if c.Bool("sentry") {
					hook, err := logrus_sentry.NewSentryHook(
						c.String("sentry-dsn"),
						[]logrus.Level{
							logrus.PanicLevel,
							logrus.FatalLevel,
							logrus.ErrorLevel,
							logrus.WarnLevel,
						},
					)
					if err != nil {
						logrus.Fatal(errors.WithStack(err))
					}

					logrus.AddHook(hook)
				}

				if err := server.Start(
					c.String("app-uri"), c.String("jwt-secret"), c.Bool("debug"),
					c.String("database-uri"), c.String("database-name"),
					c.String("facebook-app-id"), c.String("facebook-app-secret"),
					c.String("minio-uri"), c.String("minio-access-key"),
					c.String("minio-secret-key"), c.String("minio-bucket"),
					c.Bool("minio-ssl"),
				); err != nil {
					logrus.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
