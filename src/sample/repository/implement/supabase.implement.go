package SupabaseImplementation

import (
	"go_echo_example/configs/database/supabase"
	SampleDTO "go_echo_example/src/sample/dtos"
	SampleModel "go_echo_example/src/sample/repository/models"
)

func InsertSample(data *SampleDTO.SampleDTO) *SampleDTO.SampleDTO {
	supabase := supabase.DB()

	sample := &SampleModel.Sample{
		Title: data.Title,
	}

	if err := supabase.Create(&sample).Error; err != nil {
		panic(err)
	}
	return data
}
