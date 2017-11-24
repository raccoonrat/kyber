package edwards25519

import (
	"crypto/sha256"
	"hash"
	"io"
	"reflect"

	"github.com/dedis/fixbuf"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/internal/marshalling"
	"github.com/dedis/kyber/xof/blake"
)

// SuiteEd25519 implements some basic functionalities such as Group, HashFactory,
// CipherFactory and XOFFactory.
type SuiteEd25519 struct {
	Curve
}

// Hash return a newly instanciated sha256 hash function.
func (s *SuiteEd25519) Hash() hash.Hash {
	return sha256.New()
}

// XOF returns an XOF which is implemented via the Blake2b hash.
func (s *SuiteEd25519) XOF(key []byte) kyber.XOF {
	return blake.New(key)
}

func (s *SuiteEd25519) Read(r io.Reader, objs ...interface{}) error {
	return fixbuf.Read(r, s, objs...)
}

func (s *SuiteEd25519) Write(w io.Writer, objs ...interface{}) error {
	return fixbuf.Write(w, objs)
}

// New implements the kyber.Encoding interface
func (s *SuiteEd25519) New(t reflect.Type) interface{} {
	return marshalling.GroupNew(s, t)
}

// NewAES128SHA256Ed25519 returns a cipher suite based on AES-128, SHA-256, and
// the Ed25519 curve.
func NewAES128SHA256Ed25519() *SuiteEd25519 {
	suite := new(SuiteEd25519)
	return suite
}
