package goutils

import (
	"fmt"
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

func (ts *TestSuite) TestReadByteInt64() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	result, err := ReadByteInt64(b)
	assert.NoError(ts.T(), err, "Binary.read int64 in bytes should not return an error, %v", err)
	fmt.Printf("TestReadByteInt64 : Result = %d\n", result)
}

//func (ts *TestSuite) TestSessionTable() {
//assert.NoError(ts.T(), err, "Submitting VerifCode job should not return an error")
////assert.NoError(s.T(), err, "Creating Hbase store should not return an error, %v", err)
////assert.Exactly(s.T(), s.session.Uid, result.Uid, "Mismatched error returned: [expected=%v, got=%v]", s.session.Uid, result.Uid)
////assert.IsType(s.T(), reflect.TypeOf(s.session), reflect.TypeOf(result), "Result should be of type *Session")
//}
