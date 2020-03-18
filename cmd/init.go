package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gojav/config"
	"log"
	"os"
)

func Setup() {
	app := &cli.App{
		Name:  "jav",
		Usage: "Crawl javbus Magnet",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "parallel",
				Aliases:     []string{"p"},
				Usage:       "设置每秒抓取请求数",
				Value:       2,
				Destination: &config.Cfg.Parallel,
			},
			&cli.IntFlag{
				Name:        "timeout",
				Aliases:     []string{"t"},
				Usage:       "自定义连接超时时间(毫秒)",
				Value:       5000,
				Destination: &config.Cfg.Timeout,
			},
			&cli.IntFlag{
				Name:        "limit",
				Aliases:     []string{"l"},
				Usage:       "设置抓取影片的数量上限，0为抓取全部影片",
				Value:       0,
				Destination: &config.Cfg.Limit,
			},
			&cli.StringFlag{
				Name:        "proxy",
				Aliases:     []string{"x"},
				Usage:       "使用代理服务器, 例：-x http://127.0.0.1:8087",
				Value:       "127.0.0.1:1087",
				Destination: &config.Cfg.Proxy,
			},
			&cli.StringFlag{
				Name:        "search",
				Aliases:     []string{"s"},
				Usage:       "搜索关键词，可只抓取搜索结果的磁链或封面",
				Value:       "",
				Destination: &config.Cfg.Search,
			},
			&cli.StringFlag{
				Name:        "base",
				Aliases:     []string{"b"},
				Usage:       "自定义抓取的起始页",
				Value:       config.BaseUrl,
				Destination: &config.Cfg.Base,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "设置磁链和封面抓取结果的保存位置，默认为当前主目录下的 magnets 文件夹",
				Value:       "magnets",
				Destination: &config.Cfg.Output,
			},
			&cli.BoolFlag{
				Name:        "nomag",
				Aliases:     []string{"n"},
				Usage:       "是否抓取尚无磁链的影片",
				Value:       false,
				Destination: &config.Cfg.Nomag,
			},
			&cli.BoolFlag{
				Name:        "allmag",
				Aliases:     []string{"a"},
				Usage:       "是否抓取影片的所有磁链(默认只抓取文件体积最大的磁链",
				Value:       false,
				Destination: &config.Cfg.Nomag,
			},
			&cli.BoolFlag{
				Name:        "nopic",
				Aliases:     []string{"N"},
				Usage:       "不抓取图片",
				Value:       false,
				Destination: &config.Cfg.Nopic,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("========== 获取资源站点：%s ==========\n", config.BaseUrl)
			fmt.Printf("并行链接数: %d, 连接超时设置：%d \n", config.Cfg.Parallel, config.Cfg.Timeout)
			fmt.Printf("磁链保存位置: %s, 代理服务器: %s \n", config.Cfg.Output, config.Cfg.Proxy)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

