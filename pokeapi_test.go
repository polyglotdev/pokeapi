package pokeapi

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_GetPokemon(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		idOrName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pokemon
		wantErr bool
	}{
		{
			name: "Pikachu",
			fields: fields{
				httpClient: http.DefaultClient,
				baseURL:    "https://pokeapi.co/api/v2/",
			},
			args: args{
				idOrName: "pikachu",
			},
			want: &Pokemon{
				Name: "pikachu",
				ID:   25,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
				baseURL:    tt.fields.baseURL,
			}
			got, err := c.GetPokemon(tt.args.idOrName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetPokemon() = %v, want %v", got, tt.want)
			}
		})
	}
}
