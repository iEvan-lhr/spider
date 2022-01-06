package process

import "spider/Persistence"

type Scheduler struct {
	works    []Persistence.Work
	missions map[string]byte
}

//value systemQueue struct {
//	queues int
//	cpuNum int
//}

func (s *Scheduler) GenerateWork(mission string) {
	if len(s.missions) == 0 {
		s.missions = map[string]byte{mission: 1}
	} else {
		if _, ok := s.missions[mission]; ok {
			s.missions[mission] += 1
		} else {
			s.missions[mission] = 1
		}
	}
}

func (s *Scheduler) ReleaseWork() {

}

func (s *Scheduler) CleanWork() {

}

func (s *Scheduler) SplitWork() {

}
