package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	//sid string
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupSuite() {
}

// Uint64ToByte converts an uint64 to bytes array in BigEndian
func (ts *TestSuite) TestUint64ToByte() {
	data := uint64(1111)
	bData := Uint64ToByte(data)
	newData := ByteToUint64(bData)
	assert.Equal(ts.T(), data, newData, "Uint64 to []byte conversion should not return error. [expect: %V, got %v]", data, newData)
}

//func (ts *TestSuite) TestSessionTable() {
//assert.NoError(ts.T(), err, "Submitting VerifCode job should not return an error")
////assert.NoError(s.T(), err, "Creating Hbase store should not return an error, %v", err)
////assert.Exactly(s.T(), s.session.Uid, result.Uid, "Mismatched error returned: [expected=%v, got=%v]", s.session.Uid, result.Uid)
////assert.IsType(s.T(), reflect.TypeOf(s.session), reflect.TypeOf(result), "Result should be of type *Session")
//}
