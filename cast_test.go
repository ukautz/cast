package cast

type testString struct{}

func (s *testString) String() string {
	return "123"
}

var (
	testStringer1 = &testString{}
	testStringer2 = &testStringer1
	testStringer3 = &testStringer2
)
