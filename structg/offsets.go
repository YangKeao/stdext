package structg

import (
	"debug/dwarf"
	"debug/elf"
	"errors"
	"unsafe"
)

// Offsets defines the offsets of fields in G
type Offsets struct {
	// GoID represents the offset of `goid` field
	GoID int64
}

// FieldOffsets records the offset of every field of struct G
var FieldOffsets Offsets

// initGFieldOffsets initialize the offsets of G on an offsets struct
func initGFieldOffsets(exe *elf.File) error {
	data, err := exe.DWARF()
	if err != nil {
		return err
	}

	notFoundFields := map[string]struct{}{
		"goid": {},
	}

	entryReader := data.Reader()
	for {
		entry, err := entryReader.Next()
		if err != nil {
			return err
		}
		if entry.Tag != dwarf.TagStructType {
			continue
		}
		if entry.AttrField(dwarf.AttrName).Val.(string) != "runtime.g" {
			continue
		}

		typ, err := data.Type(entry.Offset)
		if err != nil {
			return err
		}

		structTyp, ok := typ.(*dwarf.StructType)
		if !ok {
			return errors.New("runtime.g is not struct")
		}
		for _, field := range structTyp.Field {
			if field.Name == "goid" {
				FieldOffsets.GoID = field.ByteOffset
				delete(notFoundFields, "goid")

				if len(notFoundFields) == 0 {
					return nil
				}
			}
		}

		errMsg := "field "
		for field := range notFoundFields {
			errMsg += field + " "
		}
		errMsg += " not found"
		return errors.New(errMsg)
	}
}

// gOffset records the offset of G struct from TLS (FS Base)
var gOffset uint64

// getGOffset gets the offset of G struct through ELF information
// it assumes this program is compiled with CGO enabled
func getGOffsetFromELF(exe *elf.File) (uint64, error) {
	var tls *elf.Prog
	for _, prog := range exe.Progs {
		if prog.Type == elf.PT_TLS {
			tls = prog
			break
		}
	}
	var tlsg *elf.Symbol
	if tls != nil {
		symbols, err := exe.Symbols()
		if err != nil {
			return 0, err
		}
		for _, sym := range symbols {
			if sym.Name == "runtime.tlsg" {
				tlsg = &sym
				break
			}
		}
	}

	if tlsg == nil {
		// negative of pointer size
		var offset uint64
		offset = 8
		offset = ^offset
		return offset + 1, nil
	}

	// The address of G is `TLS - tls.Memsz + tlsg.Value`
	return ^(tls.Memsz) + 1 + tlsg.Value, nil
}

// GetGAddr returns the address of G struct
func GetGAddr() uintptr {
	fsBase := readTLS()
	addr := fsBase + gOffset
	return *(*uintptr)(unsafe.Pointer(uintptr(addr)))
}
