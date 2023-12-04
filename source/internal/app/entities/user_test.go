package entities_test

import (
	"testing"
	"time"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/entities"
)

func Test_age_gte(t *testing.T) {
	t.Parallel()

	type args struct {
		birthday time.Time
		at       time.Time
		age      int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "more 18",
			args: args{
				birthday: time.Date(2005, 11, 30, 0, 0, 0, 0, time.Local),
				at:       time.Date(2023, 11, 30, 0, 0, 0, 0, time.Local),
				age:      18,
			},
			want: true,
		},
		{
			name: "less 18",
			args: args{
				birthday: time.Date(2005, 11, 30, 0, 0, 0, 0, time.Local),
				at:       time.Date(2023, 11, 29, 0, 0, 0, 0, time.Local),
				age:      18,
			},
			want: false,
		},
	}

	for i := 0; i < len(tests); i++ {
		i := i
		t.Run(tests[i].name, func(t *testing.T) {
			t.Parallel()
			if got := entities.AgeGte(tests[i].args.birthday, tests[i].args.at, tests[i].args.age); got != tests[i].want {
				t.Errorf("age_gte() = %v, want %v", got, tests[i].want)
			}
		})
	}
}
