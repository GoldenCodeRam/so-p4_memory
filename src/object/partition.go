package object

type PartitionNumber int

type Partition struct {
	Number  PartitionNumber
	Size    int
	Process *Process
}
