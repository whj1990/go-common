package enum

const (
	DataModelDimensionType = "0"
	DataModelMeasureType   = "1"

	IndexDefinitionDoNotCountCalculateType       = "0"
	IndexDefinitionCustomExpressionCalculateType = "1"
	IndexDefinitionSumCalculateType              = "2"
	IndexDefinitionAvgCalculateType              = "3"
	IndexDefinitionMaxCalculateType              = "4"
	IndexDefinitionMinCalculateType              = "5"
	IndexDefinitionCountCalculateType            = "6"
	IndexDefinitionDistinctCalculateType         = "7"

	IndexDefinitionFilterCustomExpressionCalculateType = "0"
	IndexDefinitionFilterEqualCalculateType            = "1"
	IndexDefinitionFilterNotEqualCalculateType         = "2"
	IndexDefinitionFilterGreaterCalculateType          = "3"
	IndexDefinitionFilterGreaterEqualCalculateType     = "4"
	IndexDefinitionFilterLessThenCalculateType         = "5"
	IndexDefinitionFilterLessThenEqualCalculateType    = "6"

	IndexDefinitionAtomType   = "0"
	IndexDefinitionDeriveType = "1"
)
