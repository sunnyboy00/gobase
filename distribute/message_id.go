package distribute

/**************************************************************************************
Code Description    : 该文件主要是用于生成分布式消息ID
					  但是按照reactor的模式来的管道,下面代码还要很多事情没有做
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

import (
	"sync/atomic"
	"time"
)

var (
	startUnixTimestamp int64
	nodeID             int64
	sequence           int64
)

const (
	nodeIDShfitBits = 24
	timeShiftBits   = 32
	nodeIDMask      = 0xff
	sequenceMask    = 0xffffff
)

// InitDistribute 用于初始化分布式模块,nodeID目前仅仅是支持
func InitDistribute(nID int64) {
	if nodeID > 128 || nodeID < -127 {
		panic("sorry,current only support -127 to 128")
	}
	nodeID = nID << nodeIDShfitBits
	startUnixTimestamp = time.Now().Unix()
}

// GetUniqueID 用于获取唯一ID
func GetUniqueID() int64 {
	return ((time.Now().Unix() - startUnixTimestamp) << timeShiftBits) | nodeID | (atomic.AddInt64(&sequence, 1) & sequenceMask)
}

// GetNodeID 用于获取结点ID
func GetNodeID(id int64) int64 {
	return id >> nodeIDShfitBits & nodeIDMask
}
