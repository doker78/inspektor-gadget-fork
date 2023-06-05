// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type execsnoopEvent struct {
	MntnsId   uint64
	Timestamp uint64
	Pid       uint32
	Ppid      uint32
	Uid       uint32
	Gid       uint32
	Retval    int32
	ArgsCount int32
	ArgsSize  uint32
	Comm      [16]uint8
	Args      [7680]uint8
	_         [4]byte
}

// loadExecsnoop returns the embedded CollectionSpec for execsnoop.
func loadExecsnoop() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ExecsnoopBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load execsnoop: %w", err)
	}

	return spec, err
}

// loadExecsnoopObjects loads execsnoop and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*execsnoopObjects
//	*execsnoopPrograms
//	*execsnoopMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadExecsnoopObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadExecsnoop()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// execsnoopSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopSpecs struct {
	execsnoopProgramSpecs
	execsnoopMapSpecs
}

// execsnoopSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopProgramSpecs struct {
	IgExecveE *ebpf.ProgramSpec `ebpf:"ig_execve_e"`
	IgExecveX *ebpf.ProgramSpec `ebpf:"ig_execve_x"`
}

// execsnoopMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopMapSpecs struct {
	Events               *ebpf.MapSpec `ebpf:"events"`
	Execs                *ebpf.MapSpec `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.MapSpec `ebpf:"gadget_mntns_filter_map"`
}

// execsnoopObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopObjects struct {
	execsnoopPrograms
	execsnoopMaps
}

func (o *execsnoopObjects) Close() error {
	return _ExecsnoopClose(
		&o.execsnoopPrograms,
		&o.execsnoopMaps,
	)
}

// execsnoopMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopMaps struct {
	Events               *ebpf.Map `ebpf:"events"`
	Execs                *ebpf.Map `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.Map `ebpf:"gadget_mntns_filter_map"`
}

func (m *execsnoopMaps) Close() error {
	return _ExecsnoopClose(
		m.Events,
		m.Execs,
		m.GadgetMntnsFilterMap,
	)
}

// execsnoopPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopPrograms struct {
	IgExecveE *ebpf.Program `ebpf:"ig_execve_e"`
	IgExecveX *ebpf.Program `ebpf:"ig_execve_x"`
}

func (p *execsnoopPrograms) Close() error {
	return _ExecsnoopClose(
		p.IgExecveE,
		p.IgExecveX,
	)
}

func _ExecsnoopClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed execsnoop_bpfel_x86.o
var _ExecsnoopBytes []byte
