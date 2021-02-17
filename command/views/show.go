package views

import (
	"fmt"

	"github.com/hashicorp/terraform/command/arguments"
	"github.com/hashicorp/terraform/plans"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/terraform"
)

type Show interface {
	Plan(plan *plans.Plan, baseState *states.State, schemas *terraform.Schemas)
}

func NewShow(vt arguments.ViewType, view *View) Show {
	switch vt {
	case arguments.ViewHuman:
		return &ShowHuman{View: *view}
	default:
		panic(fmt.Sprintf("unknown view type %v", vt))
	}
}

type ShowHuman struct {
	View
}

var _ Show = (*ShowHuman)(nil)

func (v *ShowHuman) Plan(plan *plans.Plan, baseState *states.State, schemas *terraform.Schemas) {
	renderPlan(plan, baseState, schemas, &v.View)
}
