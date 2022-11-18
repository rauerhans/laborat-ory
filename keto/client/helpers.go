//source & inspiration: https://github.com/ory/keto/blob/6c0e1ba87f4d3a355cebd0ea77f28319be2dd606/cmd/relationtuple/output.go

package client

import (
	"encoding/json"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	Collection struct {
		relations []*rts.RelationTuple
	}
	OutputTuple struct {
		*rts.RelationTuple
	}
)

func NewCollection(rels []*rts.RelationTuple) (*Collection, error) {
	r := &Collection{relations: rels}
	//for i, rel := range rels {
	//	var err error
	//	r.relations[i], err = &rts.RelationTuple{}).FromDataProvider(rel)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	return r, nil
}

func NewAPICollection(rels []*rts.RelationTuple) *Collection {
	return &Collection{relations: rels}
}

func (r *Collection) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *Collection) Table() [][]string {
	ir := r.relations

	data := make([][]string, len(ir))
	for i, rel := range ir {
		var sub string
		if rel.Subject != nil {
			sub = rel.Subject.String()
		} else {
			sub = ""
		}

		data[i] = []string{rel.Namespace, rel.Object, rel.Relation, sub}
	}

	return data
}

func (r *Collection) Interface() interface{} {
	return r.relations
}

func (r *Collection) MarshalJSON() ([]byte, error) {
	ir := r.relations
	return json.Marshal(ir)
}

func (r *Collection) UnmarshalJSON(raw []byte) error {
	return json.Unmarshal(raw, &r.relations)
}

func (r *Collection) Len() int {
	return len(r.relations)
}

func (r *Collection) IDs() []string {
	ts := r.relations
	ids := make([]string, len(ts))
	for i, rt := range ts {
		ids[i] = rt.String()
	}
	return ids
}

func (r *OutputTuple) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT ID",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *OutputTuple) Columns() []string {
	return []string{
		r.Namespace,
		r.Object,
		r.Relation,
		r.Subject.String(),
	}
}
