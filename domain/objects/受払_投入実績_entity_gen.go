// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import types "techbookfest16-sample/domain/types"

type Ent受払投入実績 struct {
	*Ent受払    `json:"受払,omitempty"`
	Get投入数量   types.Quantity `json:"投入数量"`
	Get製造指図ID types.No       `json:"製造指図ID"`
}

func NewEnt受払投入実績(受払 *Ent受払, 投入数量 types.Quantity, 製造指図ID types.No) (*Ent受払投入実績, error) {
	e := &Ent受払投入実績{
		Ent受払:     受払,
		Get投入数量:   投入数量,
		Get製造指図ID: 製造指図ID,
	}
	err := e.Validate()
	return e, err
}
