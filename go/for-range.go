package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type lesson struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}
type section struct {
	Title               string   `json:"title"`
	Position            int      `json:"position"`
	ResetLessonPosition bool     `json:"reset_lesson_position"`
	Lessons             []lesson `json:"lessons"`
}

func main() {
	sections := []section{
		{
			Title:               "Getting started",
			ResetLessonPosition: false,
			Lessons: []lesson{
				{Name: "Welcome"},
				{Name: "Installation"},
			},
		},

		{
			Title:               "Basic operator",
			ResetLessonPosition: false,
			Lessons: []lesson{
				{Name: "Addition / Subtraction"},
				{Name: "Multiplication / Division"},
			},
		},

		{
			Title:               "Advanced topics",
			ResetLessonPosition: true,
			Lessons: []lesson{
				{Name: "Mutability"},
				{Name: "Immutability"},
			},
		},
	}

	sectionCounter := 1
	lessonCounter := 1

	for i := range sections {
		if sections[i].ResetLessonPosition {
			lessonCounter = 1
		}

		sections[i].Position = sectionCounter
		sectionCounter++

		for j := range sections[i].Lessons {
			sections[i].Lessons[j].Position = lessonCounter
			lessonCounter++
		}
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	encoder.Encode(sections)

	fmt.Println(buffer.String())

}
