package observability

import (
	"strconv"

	"github.com/beepfd/server/models"
	"github.com/cen-ngc5139/BeePF/loader/lib/src/meta"
	lib "github.com/cen-ngc5139/BeePF/loader/lib/src/observability/topology"
	"github.com/cilium/ebpf"
	"github.com/pkg/errors"
)

type Topo struct{}

func NewTopo() *Topo {
	return &Topo{}
}

func (t *Topo) GetTopo() (meta.Topology, error) {
	topology, err := lib.MergeTopology()
	if err != nil {
		return topology, errors.Wrap(err, "获取程序失败")
	}

	return topology, nil
}

func (t *Topo) ListProgs() ([]models.ProgramInfoWrapper, error) {
	progs, err := lib.ListAllPrograms()
	if err != nil {
		return nil, errors.Wrap(err, "获取程序失败")
	}

	programInfos := make([]models.ProgramInfoWrapper, 0, len(progs))
	for _, info := range progs {
		wrapper, err := ConvertProgToWrapper(info)
		if err != nil {
			continue
		}

		programInfos = append(programInfos, wrapper)
	}

	return programInfos, nil
}

func (t *Topo) GetProgDetail(progID string) (*models.ProgramDetail, error) {
	progIDInt, err := strconv.ParseUint(progID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "转换程序 ID 失败")
	}

	prog, err := lib.GetProgInfo(ebpf.ProgramID(progIDInt))
	if err != nil {
		return nil, errors.Wrap(err, "获取程序失败")
	}

	wrapper, err := ConvertProgToWrapper(prog)
	if err != nil {
		return nil, errors.Wrap(err, "转换程序失败")
	}

	detail := models.ProgramDetail{
		ProgramInfoWrapper: wrapper,
	}

	for _, v := range wrapper.Maps {
		mapInfo, err := lib.GetMapInfo(v)
		if err != nil {
			return nil, errors.Wrap(err, "获取 map 失败")
		}

		mapWrapper, err := ConvertMapToWrapper(mapInfo)
		if err != nil {
			return nil, errors.Wrap(err, "转换 map 失败")
		}

		detail.MapsDetail = append(detail.MapsDetail, mapWrapper)
	}

	return &detail, nil
}

func (t *Topo) DumpXlated(progID string) ([]byte, error) {
	progIDInt, err := strconv.ParseUint(progID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "转换程序 ID 失败")
	}

	dump, err := lib.GetProgDumpXlated(ebpf.ProgramID(progIDInt))
	if err != nil {
		return nil, errors.Wrap(err, "获取程序 dump 失败")
	}

	return dump, nil
}

func (t *Topo) DumpJited(progID string) ([]byte, error) {
	progIDInt, err := strconv.ParseUint(progID, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "转换程序 ID 失败")
	}

	dump, err := lib.GetProgDumpJited(ebpf.ProgramID(progIDInt))
	if err != nil {
		return nil, errors.Wrap(err, "获取程序 dump 失败")
	}

	return dump, nil
}
