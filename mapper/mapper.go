package mapper

type MapperRecord struct {
	Fields map[string]MapperField
}

type MapperField struct {
	Key       string   // the prompt for the field in the whoisd output
	Value     []string // if the field has prearranged value (not use any field from the storage)
	Name      []string // field name in the storage, if the field is not prearranged ("value" is not defined)
	Format    string   // special instructions to indicate how to display field
	Multiple  bool     // if this option is set to 'true', then for each value will be repeated prompt
	Hide      bool     // if this option is set to 'true', the value of the field will not shown in the whois output
	Related   string   // field name in the storage through which the request for
	RelatedBy string   // field name in the storage through which the related request for
	RelatedTo string   // table/type name in the storage through which made relation
}
