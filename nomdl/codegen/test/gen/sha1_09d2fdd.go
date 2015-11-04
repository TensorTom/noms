// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_sha1_09d2fdd_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("D",
			[]types.Field{
				types.Field{"structField", types.MakeTypeRef(ref.Parse("sha1-1c216c6f1d6989e4ede5f78b7689214948dabeef"), 0), false},
				types.Field{"enumField", types.MakeTypeRef(ref.Parse("sha1-1c216c6f1d6989e4ede5f78b7689214948dabeef"), 1), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("DUser",
			[]types.Field{
				types.Field{"Dfield", types.MakeTypeRef(ref.Ref{}, 0), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-1c216c6f1d6989e4ede5f78b7689214948dabeef"),
	})
	__genPackageInFile_sha1_09d2fdd_CachedRef = types.RegisterPackage(&p)
}

// D

type D struct {
	_structField S
	_enumField   E

	ref *ref.Ref
}

func NewD() D {
	return D{
		_structField: NewS(),
		_enumField:   NewE(),

		ref: &ref.Ref{},
	}
}

type DDef struct {
	StructField SDef
	EnumField   E
}

func (def DDef) New() D {
	return D{
		_structField: def.StructField.New(),
		_enumField:   def.EnumField,
		ref:          &ref.Ref{},
	}
}

func (s D) Def() (d DDef) {
	d.StructField = s._structField.Def()
	d.EnumField = s._enumField
	return
}

var __typeRefForD types.TypeRef

func (m D) TypeRef() types.TypeRef {
	return __typeRefForD
}

func init() {
	__typeRefForD = types.MakeTypeRef(__genPackageInFile_sha1_09d2fdd_CachedRef, 0)
	types.RegisterStruct(__typeRefForD, builderForD, readerForD)
}

func builderForD() chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := D{ref: &ref.Ref{}}
		s._structField = (<-c).(S)
		s._enumField = (<-c).(E)
		c <- s
	}()
	return c
}

func readerForD(v types.Value) chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := v.(D)
		c <- s._structField
		c <- s._enumField
	}()
	return c
}

func (s D) Equals(other types.Value) bool {
	return other != nil && __typeRefForD.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s D) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s D) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForD.Chunks()...)
	chunks = append(chunks, s._structField.Chunks()...)
	return
}

func (s D) StructField() S {
	return s._structField
}

func (s D) SetStructField(val S) D {
	s._structField = val
	s.ref = &ref.Ref{}
	return s
}

func (s D) EnumField() E {
	return s._enumField
}

func (s D) SetEnumField(val E) D {
	s._enumField = val
	s.ref = &ref.Ref{}
	return s
}

// DUser

type DUser struct {
	_Dfield D

	ref *ref.Ref
}

func NewDUser() DUser {
	return DUser{
		_Dfield: NewD(),

		ref: &ref.Ref{},
	}
}

type DUserDef struct {
	Dfield DDef
}

func (def DUserDef) New() DUser {
	return DUser{
		_Dfield: def.Dfield.New(),
		ref:     &ref.Ref{},
	}
}

func (s DUser) Def() (d DUserDef) {
	d.Dfield = s._Dfield.Def()
	return
}

var __typeRefForDUser types.TypeRef

func (m DUser) TypeRef() types.TypeRef {
	return __typeRefForDUser
}

func init() {
	__typeRefForDUser = types.MakeTypeRef(__genPackageInFile_sha1_09d2fdd_CachedRef, 1)
	types.RegisterStruct(__typeRefForDUser, builderForDUser, readerForDUser)
}

func builderForDUser() chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := DUser{ref: &ref.Ref{}}
		s._Dfield = (<-c).(D)
		c <- s
	}()
	return c
}

func readerForDUser(v types.Value) chan types.Value {
	c := make(chan types.Value)
	go func() {
		s := v.(DUser)
		c <- s._Dfield
	}()
	return c
}

func (s DUser) Equals(other types.Value) bool {
	return other != nil && __typeRefForDUser.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s DUser) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s DUser) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForDUser.Chunks()...)
	chunks = append(chunks, s._Dfield.Chunks()...)
	return
}

func (s DUser) Dfield() D {
	return s._Dfield
}

func (s DUser) SetDfield(val D) DUser {
	s._Dfield = val
	s.ref = &ref.Ref{}
	return s
}
