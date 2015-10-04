package mapper

// Bundle - set of entries
type Bundle []Entry

type Entry struct {
	// TLDs - list of TLDs, which accepted by specified Entry
	TLDs []string `json: "TLDs"`

	// a list of fields from "01" to last number "nn" in ascending order
	Fields map[string]Field `json: "fields"`
}

// Field - representation of one field
type Field struct {
	// a label in whois output
	Key string `json: "key"`

	// used if a field has constant value
	Value []string `json: "value"`

	// a name of the field in the database, if "value" is not defined
	Name []string `json: "name"`

	// special instructions to indicate how to display the field
	Format string `json: "format"`

	// if this option is set to 'true', each value will be repeated in whois output with the same label
	Multiple bool `json: "multiple"`

	// if this option is set to 'true', a value of the field will not shown in whois output
	Hide bool `json: "hide"`

	// a name of the field in a database through which a request for
	Related string `json: "related"`

	// a name of the field in a database through which related a request for
	RelatedBy string `json: "related_by"`

	// a name of the table/type in a database through which made a relation
	RelatedTo string `json: "related_to"`
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
