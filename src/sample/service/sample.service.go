package SampleSupabaseService

import (
	SampleDTO "go_echo_example/src/sample/dtos"
	SupabaseImplementation "go_echo_example/src/sample/repository/implement"
)

func Insert(data *SampleDTO.SampleDTO) *SampleDTO.SampleDTO {
	r := SupabaseImplementation.InsertSample(data)

	return r
}
