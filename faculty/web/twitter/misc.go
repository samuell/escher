// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package twitter installs a faculty for access to the Twitter API.
package twitter

import (
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/think"
)

// AnswerMaterializer ...
type AnswerMaterializer struct{}

func (AnswerMaterializer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Name", "Sentence")
}

// ConsumerMaterializer ...
type ConsumerMaterializer struct{}

func (ConsumerMaterializer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Key", "Secret")
}

// AccessMaterializer ...
type AccessMaterializer struct{}

func (AccessMaterializer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "Token", "Secret")
}

// UserTimelineQueryMaterializer ...
type UserTimelineQueryMaterializer struct{}

func (UserTimelineQueryMaterializer) Materialize() think.Reflex {
	return basic.MaterializeConjunction("_", "UserId", "ScreenName", "AfterId", "NotAfterId", "Count")
}