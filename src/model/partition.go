package model

import "so-p4_memory/src/object"

type Partition struct {
    *object.Partition

	Process *Process
}
