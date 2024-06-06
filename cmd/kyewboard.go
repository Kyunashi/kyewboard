package main

import (
	"context"
	"net/http"

	"kyewboard/pkg/db"
	"kyewboard/pkg/view"

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewPlayer() db.Player {
    stats := map[string]int{
        "Vitality": 0,
        "Strength": 0,
        "Inteligence": 0,
        "Sense": 0,
        "Agility": 0,
    }
    dev := db.Skill{Title: "Development", Category: "IT", Level: 1, Experience: 1,}
    sec := db.Skill{Title: "IT Security", Category: "IT", Level: 1, Experience: 1,}
    skate := db.Skill{Title: "Skateboarding", Category: "Sport", Level: 1, Experience: 1,}
    garden := db.Skill{Title: "Gardening", Category: "Biology", Level: 1, Experience: 1,}
    rocketleauge := db.Skill{Title: "Rocketleague", Category: "Esport", Level: 1, Experience: 1,}
    
    skills := map[string]db.Skill{
        "Development": dev,
        "IT Security": sec,
        "Skateboarding": skate,
        "Gardening": garden,
        "Rocketleague": rocketleauge,
    }
    return db.Player{
        Stats : stats,
        Skills: skills,
        Experience : 0,
        Level : 1,
        Id : 1,
        Name : "Kyew",
    }
}


func main() {
	e := echo.New()
    e.Use(middleware.Logger())

    rewards := []string{"+1000 GO Exp", "+1000 Html Exp"}
    objectives := []string{"Setup GO Server", "Setup Templ", "Setup Air"}
    quest := db.Quest{Id: 1, Message: "Kyewboard setup quest", Status: "Pending",Objectives: objectives, Rewards: rewards, Assignee: "kyew"}
    player := NewPlayer()
	component := view.Index(quest, player)
    body := view.Body(quest, player)
	// quests := make([]Quest)

	// quests := append(quests, quest)
	
	e.Static("/static", "/assets")

	e.GET("/", func(c echo.Context) error {
        return component.Render(context.Background(), c.Response().Writer)
	})

    e.POST("/accepted", func(c echo.Context) error {
        quest.Status = "Accepted"
        completebtn := view.CompleteDiv()
        return completebtn.Render(context.Background(),c.Response().Writer)

    })

    e.POST("/declined", func(c echo.Context) error {
        quest.Status = "Declined"
        return c.String(http.StatusOK, quest.Status)
    })
    
    e.POST("/completed", func(c echo.Context) error {
        return body.Render(context.Background(), c.Response().Writer)
    })
	e.Logger.Fatal(e.Start(":42069"))
}
