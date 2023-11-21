package models

type AdventurerWise struct {
	Record WiseRecord
	Tests  struct {
		IAmWise struct {
			Pass int
			Fail int
		}
		DeeperUnderstanding int
		OfCourse            int
	}
}
