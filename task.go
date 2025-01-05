package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Task struct {
	Title       string
	Category    string
	Status      TaskStatus
	CreatedAt   time.Time
	CompletedAt time.Time
}

func NewTask(title, category string) Task {
	return Task{
		Title:     title,
		Category:  category,
		Status:    Pending,
		CreatedAt: time.Now(),
	}
}

func (t *Task) title() string {
	return t.Title
}

func (t *Task) category() string {
	return t.Category
}

func (t *Task) status() TaskStatus {
	return t.Status
}

func (t *Task) createdAt() time.Time {
	return t.CreatedAt
}

func (t *Task) completedAt() time.Time {
	return t.CompletedAt
}

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Completed
)

func (ts TaskStatus) String() string {
	switch ts {
	case Pending:
		return "Pending"
	case InProgress:
		return "InProgress"
	case Completed:
		return "Completed"
	default:
		log.Fatal("error: invalid task status. Choose between Pending, InProgress, or Completed")
		return "Unknown"
	}
}

func ParseTaskStatus(status string) (TaskStatus, error) {
	switch strings.ToLower(status) {
	case "pending":
		return Pending, nil
	case "inprogress":
		return InProgress, nil
	case "completed":
		return Completed, nil
	default:
		return TaskStatus(0), fmt.Errorf("error: invalid task status: %s. Choose between Pending, InProgress, or Completed", status)
	}
}
