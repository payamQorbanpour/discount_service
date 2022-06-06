package usecase

import (
	"context"
	"testing"
	"time"

	"discount_service/usecase/dto"

	"github.com/stretchr/testify/assert"
)

var testsRepositoryInsertDiscount = []struct {
	name      string
	entry     repo
	want      []dto.Discount
	wantError bool
}{
	{
		name: "insert discount success",
		entry: repo{
			db: map[string][]dto.Discount{
				"09123456789": {
					{
						Amount:    1,
						CreatedAt: time.Time(time.Now()),
					},
				},
			},
		},
		want: []dto.Discount{
			{
				Amount:    1,
				CreatedAt: time.Time(time.Now()),
			},
			{
				Amount:    2,
				CreatedAt: time.Time(time.Now()),
			},
		},
		wantError: false,
	},
}

func Test_repository_InsertDiscount(t *testing.T) {
	for _, tt := range testsRepositoryInsertDiscount {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.entry.InsertDiscount(context.Background(), "09123456789", 2)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			tt.want[1].CreatedAt = tt.entry.db["09123456789"][1].CreatedAt
			assert.Equal(t, tt.want[1], tt.entry.db["09123456789"][1])
		})
	}
}

var testsRepositoryGetDiscountByID = []struct {
	name      string
	entry     repo
	want      []dto.Discount
	wantError bool
}{
	{
		name: "insert discount success",
		entry: repo{
			db: map[string][]dto.Discount{
				"09123456789": {
					{
						Amount:    1,
						CreatedAt: time.Time(time.Now()),
					},
					{
						Amount:    2,
						CreatedAt: time.Time(time.Now()),
					},
				},
				"09123456780": {
					{
						Amount:    4,
						CreatedAt: time.Time(time.Now()),
					},
				},
			},
		},
		want: []dto.Discount{
			{
				Amount:    1,
				CreatedAt: time.Time(time.Now()),
			},
			{
				Amount:    2,
				CreatedAt: time.Time(time.Now()),
			},
		},
		wantError: false,
	},
}

func Test_repository_GetDiscountByID(t *testing.T) {
	for _, tt := range testsRepositoryGetDiscountByID {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.GetDiscountsByID(context.Background(), "09123456789")

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			for i := range tt.want {
				tt.want[i].CreatedAt = got[i].CreatedAt
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

var testsRepositoryGetDiscounts = []struct {
	name      string
	entry     repo
	want      map[string][]dto.Discount
	wantError bool
}{
	{
		name: "insert discount success",
		entry: repo{
			db: map[string][]dto.Discount{
				"09123456789": {
					{
						Amount:    1,
						CreatedAt: time.Time(time.Now()),
					},
					{
						Amount:    2,
						CreatedAt: time.Time(time.Now()),
					},
				},
				"09123456780": {
					{
						Amount:    4,
						CreatedAt: time.Time(time.Now()),
					},
				},
			},
		},
		want: map[string][]dto.Discount{
			"09123456789": {
				{
					Amount:    1,
					CreatedAt: time.Time(time.Now()),
				},
				{
					Amount:    2,
					CreatedAt: time.Time(time.Now()),
				},
			},
			"09123456780": {
				{
					Amount:    4,
					CreatedAt: time.Time(time.Now()),
				},
			},
		},
		wantError: false,
	},
}

func Test_repository_GetDiscounts(t *testing.T) {
	for _, tt := range testsRepositoryGetDiscounts {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.entry.GetDiscounts(context.Background())

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			for k := range tt.want {
				for i := range tt.want[k] {
					tt.want[k][i].CreatedAt = got[k][i].CreatedAt
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
