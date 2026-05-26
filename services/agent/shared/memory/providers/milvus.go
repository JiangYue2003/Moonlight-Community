package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/milvus-io/milvus/client/v2/column"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"

	"github.com/zhiguang/zhiguang-go/services/agent/shared/memory"
)

type MilvusFactVectorStore struct {
	enabled     bool
	cli         *milvusclient.Client
	collection  string
	vectorField string
	vectorDim   int
}

func NewMilvusFactVectorStore(enabled bool, cli *milvusclient.Client, collection, vectorField string, vectorDim int) *MilvusFactVectorStore {
	return &MilvusFactVectorStore{
		enabled:     enabled,
		cli:         cli,
		collection:  collection,
		vectorField: vectorField,
		vectorDim:   vectorDim,
	}
}

func (s *MilvusFactVectorStore) UpsertFactVectors(ctx context.Context, userID int64, vectors []memory.FactVector) error {
	if !s.enabled || s.cli == nil || userID <= 0 || len(vectors) == 0 {
		return nil
	}

	ids := make([]string, 0, len(vectors))
	uids := make([]int64, 0, len(vectors))
	factIDs := make([]string, 0, len(vectors))
	subjects := make([]string, 0, len(vectors))
	predicates := make([]string, 0, len(vectors))
	objects := make([]string, 0, len(vectors))
	sources := make([]string, 0, len(vectors))
	versions := make([]string, 0, len(vectors))
	statuses := make([]string, 0, len(vectors))
	confidences := make([]float32, 0, len(vectors))
	embeddings := make([][]float32, 0, len(vectors))

	for _, v := range vectors {
		if strings.TrimSpace(v.Fact.FactID) == "" || len(v.Vector) == 0 {
			continue
		}
		ids = append(ids, fmt.Sprintf("u%d-f%s", userID, v.Fact.FactID))
		uids = append(uids, userID)
		factIDs = append(factIDs, v.Fact.FactID)
		subjects = append(subjects, trim(v.Fact.Subject, 255))
		predicates = append(predicates, trim(v.Fact.Predicate, 255))
		objects = append(objects, trim(v.Fact.ObjectValue, 2000))
		sources = append(sources, trim(v.Fact.SourceRef, 255))
		versions = append(versions, trim(v.Fact.Version, 128))
		status := strings.TrimSpace(v.Fact.Status)
		if status == "" {
			status = "active"
		}
		statuses = append(statuses, status)
		confidences = append(confidences, float32(v.Fact.Confidence))
		embeddings = append(embeddings, v.Vector)
	}
	if len(ids) == 0 {
		return nil
	}

	opt := milvusclient.NewColumnBasedInsertOption(s.collection).
		WithVarcharColumn("id", ids).
		WithInt64Column("user_id", uids).
		WithVarcharColumn("fact_id", factIDs).
		WithVarcharColumn("subject", subjects).
		WithVarcharColumn("predicate", predicates).
		WithVarcharColumn("object_value", objects).
		WithVarcharColumn("source_ref", sources).
		WithVarcharColumn("version", versions).
		WithVarcharColumn("status", statuses).
		WithFloatVectorColumn(s.vectorField, s.vectorDim, embeddings)
	opt.WithColumns(column.NewColumnFloat("confidence", confidences))
	_, err := s.cli.Upsert(ctx, opt)
	return err
}

func (s *MilvusFactVectorStore) SearchFactVectors(ctx context.Context, q memory.Query) ([]memory.ScoredFact, error) {
	if !s.enabled || s.cli == nil || q.UserID <= 0 || q.TopK <= 0 || len(q.Vector) == 0 {
		return nil, nil
	}
	filter := fmt.Sprintf("user_id == %d and status == \"active\"", q.UserID)
	resultSets, err := s.cli.Search(ctx,
		milvusclient.NewSearchOption(s.collection, q.TopK, []entity.Vector{entity.FloatVector(q.Vector)}).
			WithANNSField(s.vectorField).
			WithFilter(filter).
			WithOutputFields("fact_id", "subject", "predicate", "object_value", "source_ref", "version", "status", "confidence"),
	)
	if err != nil {
		return nil, err
	}

	out := make([]memory.ScoredFact, 0, q.TopK)
	for _, rs := range resultSets {
		factIDCol := rs.GetColumn("fact_id")
		subCol := rs.GetColumn("subject")
		preCol := rs.GetColumn("predicate")
		objCol := rs.GetColumn("object_value")
		srcCol := rs.GetColumn("source_ref")
		verCol := rs.GetColumn("version")
		stCol := rs.GetColumn("status")
		confCol := rs.GetColumn("confidence")
		if factIDCol == nil || subCol == nil || preCol == nil || objCol == nil || srcCol == nil || verCol == nil || stCol == nil || confCol == nil {
			continue
		}
		for i := 0; i < rs.ResultCount; i++ {
			fid, e1 := factIDCol.GetAsString(i)
			sub, e2 := subCol.GetAsString(i)
			pre, e3 := preCol.GetAsString(i)
			obj, e4 := objCol.GetAsString(i)
			src, e5 := srcCol.GetAsString(i)
			ver, e6 := verCol.GetAsString(i)
			st, e7 := stCol.GetAsString(i)
			conf, e8 := confCol.GetAsDouble(i)
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil || e7 != nil || e8 != nil {
				continue
			}
			out = append(out, memory.ScoredFact{
				Fact: memory.Fact{
					FactID:      fid,
					Subject:     sub,
					Predicate:   pre,
					ObjectValue: obj,
					SourceRef:   src,
					Confidence:  conf,
					Version:     ver,
					Status:      st,
				},
				Source: "memory_milvus",
				Rank:   len(out) + 1,
			})
		}
	}
	return out, nil
}

func trim(s string, n int) string {
	r := []rune(strings.TrimSpace(s))
	if n <= 0 || len(r) <= n {
		return string(r)
	}
	return string(r[:n])
}

var _ memory.VectorStore = (*MilvusFactVectorStore)(nil)
