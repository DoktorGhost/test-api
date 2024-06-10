package validator

import (
	"testing"
)

func Test_validateAge(t *testing.T) {
	type args struct {
		age uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{name: "Тест №1 валидный возраст",
			args:    args{age: 55},
			want:    55,
			wantErr: false,
		},
		{name: "Тест №2 невалидный возраст",
			args:    args{age: 155},
			want:    0,
			wantErr: true,
		},
		{name: "Тест №3 невалидный возраст",
			args:    args{age: 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAge(tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateAge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Тест №1 валидный email",
			args:    args{email: "ff@ss.ru"},
			want:    "ff@ss.ru",
			wantErr: false,
		},
		{name: "Тест №2 невалидный email",
			args:    args{email: "ffss.ru"},
			want:    "",
			wantErr: true,
		},
		{name: "Тест №3 невалидный email",
			args:    args{email: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateLastname(t *testing.T) {
	type args struct {
		surname string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Тест №1 валидная фамилия",
			args:    args{surname: "сиМонОв"},
			want:    "Симонов",
			wantErr: false,
		},
		{name: "Тест №2 невалидная фамилия",
			args:    args{surname: "Симонов55"},
			want:    "",
			wantErr: true,
		},
		{name: "Тест №3 невалидная фамилия",
			args:    args{surname: "4534523"},
			want:    "",
			wantErr: true,
		},
		{name: "Тест №4 невалидная фамилия",
			args:    args{surname: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateLastname(tt.args.surname)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateLastname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateLastname() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Тест №1 валидное имя",
			args:    args{name: "алеКсаНДр"},
			want:    "Александр",
			wantErr: false,
		},
		{name: "Тест №2 невалидное имя",
			args:    args{name: "алеК55саНДр"},
			want:    "",
			wantErr: true,
		},
		{name: "Тест №3 невалидное имя",
			args:    args{name: "234234"},
			want:    "",
			wantErr: true,
		},
		{name: "Тест №4 невалидное имя",
			args:    args{name: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
