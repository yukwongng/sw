package main

import (
	"log"

	"github.com/pensando/sw/nic/delphi/gosdk"
	"github.com/pensando/sw/nic/delphi/proto/delphi"
)

type service struct {
	name string
}

func (s *service) OnMountComplete() {
	log.Printf("OnMountComplete() done for %s\n", s.name)
}

func (s *service) Name() string {
	return s.name
}

func (s *service) OnUpgRespCreate(obj *UpgResp) {
	log.Printf("OnUpgRespCreate called %d\n", obj.GetUpgRespVal())
}

func (s *service) OnUpgRespUpdate(old, obj *UpgResp) {
	log.Printf("OnUpgRespUpdate called %d\n", obj.GetUpgRespVal())
}

func (s *service) OnUpgRespDelete(obj *UpgResp) {
	log.Printf("OnUpgRespDelete called %d\n", obj.GetUpgRespVal())
}

func main() {
	s1 := &service{
		name: "example",
	}
	c1, err := gosdk.NewClient(s1)
	if err != nil {
		panic(err)
	}
	UpgReqMount(c1, delphi.MountMode_ReadWriteMode)
	UpgRespMount(c1, delphi.MountMode_ReadMode)
	UpgRespWatch(c1, s1)

	// run the event loop
	c1.Run()

	// create an object
	u := &UpgReq{
		Key:       10,
		UpgReqCmd: UpgReqType_UpgStart,
	}
	c1.SetObject(u)
	a := make(chan struct{})
	_ = <-a
}
