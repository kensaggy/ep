package ep

// Null is a Type representing NULL values. Use Null.Data(n) to create Data
// instances of `n` nulls
var Null = &nullType{}

var _ = Types.Register("NULL", Null)

type nullType struct{}

func (t *nullType) String() string     { return t.Name() }
func (*nullType) Data(n int) Data      { return nulls(n) }
func (*nullType) DataEmpty(n int) Data { return nulls(n) }
func (*nullType) Name() string         { return "NULL" }
func (*nullType) Is(t Type) bool {
	return t.Name() == "NULL"
}

type nulls int                         // number of nulls in the set
func (nulls) Type() Type               { return Null }
func (vs nulls) Len() int              { return int(vs) }
func (nulls) Less(int, int) bool       { return false }
func (nulls) Swap(int, int)            {}
func (nulls) Slice(i, j int) Data      { return nulls(j - i) }
func (vs nulls) Append(data Data) Data { return vs + data.(nulls) }
func (vs nulls) Duplicate(t int) Data  { return vs * nulls(t) }
func (vs nulls) Strings() []string     { return make([]string, vs) }
