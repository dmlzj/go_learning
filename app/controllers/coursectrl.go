package controllers

import (
	"encoding/json"
	"restful-api/app/models"
	"github.com/revel/revel"
	)

type CourseCtrl struct {
	GorpController
}

func (c CourseCtrl) parseCourseItem() (models.Course, error) {
	Course := models.Course{}
	err := json.NewDecoder(c.Request.GetBody()).Decode(&Course)
	return Course, err
}

func (c CourseCtrl) Add() revel.Result {
	if course, err := c.parseCourseItem(); err != nil {
		return c.RenderText("Unable to parse the BidItem from JSON.")
	} else {
		// Validate the model
		course.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your Course.")
		} else {
			if err := c.Txn.Insert(&course); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJSON(course)
			}
		}
	}
}

func (c CourseCtrl) Get(id int64) revel.Result {
	course := new(models.Course)
	err := c.Txn.SelectOne(course,
		`SELECT * FROM Courses WHERE id = ?`, id)
	if err != nil {
		return c.RenderJSON("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(course)
}

func (c CourseCtrl) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	course, err := c.Txn.Select(models.Course{},
		`SELECT * FROM courses WHERE Id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJSON(course)
}

