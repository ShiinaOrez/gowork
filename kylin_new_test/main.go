package main

import (
	"fmt"
	"time"
)

type Argument map[string]interface{}

type Result map[string]interface{}

type Mission interface {
	Start(Argument) *chan int
	ReadResult() Result
}

type FirstMission struct {
	Content map[string]interface{}
}

func (mission FirstMission) Start(args Argument) *chan int {
	ch := make(chan int)
	go func() {
		fmt.Println(args["str"].(string))
		time.Sleep(5 * time.Second)
		mission.Content["ret"] = "first_mission_completed"
		mission.Content["id"] = args["id"]
		mission.Content["str"] = "next mission is running"
		ch<- 1
	}()
	return &ch
}

func (mission FirstMission) ReadResult() Result {
	return mission.Content
}

type SecondMission struct {
	Content map[string]interface{}
}

func (mission SecondMission) Start(args Argument) *chan int {
	ch := make(chan int)
	go func() {
		fmt.Println(args["str"].(string))
		time.Sleep(3 * time.Second)
		mission.Content["ret"] = "second_mission_completed"
		mission.Content["id"] = args["id"]
		ch<- 1
	}()
	return &ch
}

func (mission SecondMission) ReadResult() Result {
	return mission.Content
}

func dispather(startMissions []string, args []Argument, missionMap map[string]Mission, logicMap map[string][]string, inMap map[string]int) {
	argsMap := make(map[string]Argument)
	channels := make(map[string]*chan int)
	for index, missionID := range startMissions {
		channels[missionID] = missionMap[missionID].Start(args[index])
	}
	for len(channels) > 0 {
		for missionID, ch := range channels {
			select {
			case status := <-(*ch): {
				close(*ch)
				delete(channels, missionID)
				fmt.Println("Mission:", missionID, "status:", status)
				nextMissions := logicMap[missionID]
				for _, nextMissionID := range nextMissions {
					inMap[nextMissionID] -= 1
					result := missionMap[missionID].ReadResult()
					if _, ok := argsMap[nextMissionID]; !ok {
						argsMap[nextMissionID] = make(map[string]interface{})
					}
					for k, v := range result {
						argsMap[nextMissionID][k] = v
					}
					if in, _ := inMap[nextMissionID]; in == 0 {
						channels[nextMissionID] = missionMap[nextMissionID].Start(argsMap[nextMissionID])
					}
				}
			}
			default:
			}
		}
	}
	return
}

func main() {
	var missionMap = make(map[string]Mission)
	missionMap["1"] = FirstMission{Content:make(map[string]interface{})}
	missionMap["2"] = SecondMission{Content:make(map[string]interface{})}
	missionMap["3"] = SecondMission{Content:make(map[string]interface{})}
	var logicMap = make(map[string][]string)
	logicMap["1"] = []string{"2", "3"}
	var inMap = make(map[string]int)
	inMap["1"] = 0
	inMap["2"] = 1
	inMap["3"] = 1
	var startArg = map[string]interface{} {
		"str": "Mission First Is Running.",
	}

	dispather([]string{"1"}, []Argument{startArg}, missionMap, logicMap, inMap)
}
