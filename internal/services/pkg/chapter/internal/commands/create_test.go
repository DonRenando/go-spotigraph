package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/commands"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

func TestChapter_Create(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Invalid request type",
			req:     &struct{}{},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &chapterv1.CreateRequest{},
			wantErr: true,
		},
		{
			name: "Empty label",
			req: &chapterv1.CreateRequest{
				Label: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &chapterv1.CreateRequest{
				Label: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing chapter",
			req: &chapterv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				t1 := models.NewChapter("Foo")
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter",
			req: &chapterv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter with database error",
			req: &chapterv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, chapters)
			}

			// Prepare handler
			underTest := commands.CreateHandler(chapters)

			// Do the query
			got, err := underTest.Handle(ctx, testCase.req)

			// assert results expectations
			if testCase.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
			}
		})
	}
}
