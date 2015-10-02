package mapper

// Bundle - set of entries for specific TLDs
type Bundle []Entry

type Entry struct {
	TLDs   []string         `json: "TLDs"`
	Fields map[string]Field `json: "fields"` // a list of fields from "01" to last number "nn" in ascending order
}

// Field - representation of one field
type Field struct {
	Key       string   `json: "key"`        // the label for the field in whois output
	Value     []string `json: "value"`      // used if the field has prearranged value
	Name      []string `json: "name"`       // the name of the field in the database, if the field is not prearranged ("value" is not defined)
	Format    string   `json: "format"`     // special instructions to indicate how to display field
	Multiple  bool     `json: "multiple"`   // if this option is set to 'true', then for each value will be repeated label in whois output
	Hide      bool     `json: "hide"`       // if this option is set to 'true', the value of the field will not shown in whois output
	Related   string   `json: "related"`    // the name of the field in the database through which the request for
	RelatedBy string   `json: "related_by"` // the name of the field in the database through which the related request for
	RelatedTo string   `json: "related_to"` // the name of the table/type in the database through which made a relation
}

func (bundle Bundle) EntryByTLD(TLD string) *Entry {
	for index := range bundle {
		for _, item := range bundle[index].TLDs {
			if item == TLD {
				return &bundle[index]
			}
		}
	}

	return nil
}
