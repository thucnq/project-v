package nestedsets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetsFromNote(t *testing.T) {
	tests := []struct {
		name string
		want []Set
	}{
		{
			name: "suceess",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				root := Node{
					ID:       "1",
					ParentID: "",
					Childs: []*Node{
						{
							ID:       "2",
							ParentID: "1",
							Childs: []*Node{
								{
									ID:       "17",
									ParentID: "2",
									Childs: []*Node{
										{
											ID:       "6",
											ParentID: "17",
										},
									},
								},
							},
						},
						{
							ID:       "3",
							ParentID: "1",
							Childs: []*Node{
								{
									ID:       "7",
									ParentID: "3",
								},
								{
									ID:       "8",
									ParentID: "3",
									Childs: []*Node{
										{
											ID:       "12",
											ParentID: "8",
										},
										{
											ID:       "13",
											ParentID: "8",
										},
									},
								},
								{
									ID:       "20",
									ParentID: "3",
								},
							},
						},
						{
							ID:       "39",
							ParentID: "1",
							Childs: []*Node{
								{
									ID:       "18",
									ParentID: "39",
								},
								{
									ID:       "21",
									ParentID: "39",
								},
							},
						},
						{
							ID:       "40",
							ParentID: "1",
							Childs:   []*Node{},
						},
					},
				}
				if got := SetsFromNote(root); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("SetsFromNote() = %+v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNodeFromSets(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "success",
			args: args{
				[]Set{
					{
						ID:            "1",
						LeftBowerInt:  1,
						RightBowerInt: 28,
					},
					{
						ID:            "2",
						LeftBowerInt:  2,
						RightBowerInt: 7,
					},
					{
						ID:            "17",
						LeftBowerInt:  3,
						RightBowerInt: 6,
					},
					{
						ID:            "6",
						LeftBowerInt:  4,
						RightBowerInt: 5,
					},
					{
						ID:            "3",
						LeftBowerInt:  8,
						RightBowerInt: 19,
					},
					{
						ID:            "7",
						LeftBowerInt:  9,
						RightBowerInt: 10,
					},
					{
						ID:            "8",
						LeftBowerInt:  11,
						RightBowerInt: 16,
					},
					{
						ID:            "12",
						LeftBowerInt:  12,
						RightBowerInt: 13,
					},
					{
						ID:            "13",
						LeftBowerInt:  14,
						RightBowerInt: 15,
					},
					{
						ID:            "20",
						LeftBowerInt:  17,
						RightBowerInt: 18,
					},
					{
						ID:            "39",
						LeftBowerInt:  20,
						RightBowerInt: 25,
					},
					{
						ID:            "18",
						LeftBowerInt:  21,
						RightBowerInt: 22,
					},
					{
						ID:            "21",
						LeftBowerInt:  23,
						RightBowerInt: 24,
					},
					{
						ID:            "40",
						LeftBowerInt:  26,
						RightBowerInt: 27,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NodeFromSets(tt.args.sets); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf(
						"NodeFromSets() = %+v, want %v", got.Childs, tt.want,
					)
					if got.Childs != nil {
						for _, item := range got.Childs {
							// fmt.Printf("h1 %+v\n", *item)
							for _, childItem := range item.Childs {
								// fmt.Printf("h2 %+v\n", *childItem)
								for _, grandchildItem := range childItem.Childs {
									fmt.Printf("h3 %+v\n", *grandchildItem)
								}
							}
						}

					}
				}
			},
		)
	}
}
