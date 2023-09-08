package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type response struct {
	SlackName     string    `json:"slack_name"`
	CurrentDay    string    `json:"current_day"`
	UtcTime       time.Time `json:"utc_time"`
	Track         string    `json:"track"`
	GithubUrl     string    `json:"github_file_url"`
	GithubRepoUrl string    `json:"github_repo_url"`
	StatusCode    int       `json:"status_code"`
}

func MainController(c *fiber.Ctx) error {
	data := response{
		SlackName:     c.Query("slack_name"),
		CurrentDay:    time.Now().Format("Monday"),
		UtcTime:       time.Now().UTC().Round(time.Second * 120),
		Track:         c.Query("track"),
		GithubUrl:     "",
		GithubRepoUrl: "",
		StatusCode:    200,
	}
	return c.JSON(data)
}
