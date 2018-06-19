// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/operation.proto

package servicecontrol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the importance of the data contained in the operation.
type Operation_Importance int32

const (
	// The API implementation may cache and aggregate the data.
	// The data may be lost when rare and unexpected system failures occur.
	Operation_LOW Operation_Importance = 0
	// The API implementation doesn't cache and aggregate the data.
	// If the method returns successfully, it's guaranteed that the data has
	// been persisted in durable storage.
	Operation_HIGH Operation_Importance = 1
)

var Operation_Importance_name = map[int32]string{
	0: "LOW",
	1: "HIGH",
}
var Operation_Importance_value = map[string]int32{
	"LOW":  0,
	"HIGH": 1,
}

func (x Operation_Importance) String() string {
	return proto.EnumName(Operation_Importance_name, int32(x))
}
func (Operation_Importance) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// Represents information regarding an operation.
type Operation struct {
	// Identity of the operation. This must be unique within the scope of the
	// service that generated the operation. If the service calls
	// Check() and Report() on the same operation, the two calls should carry
	// the same id.
	//
	// UUID version 4 is recommended, though not required.
	// In scenarios where an operation is computed from existing information
	// and an idempotent id is desirable for deduplication purpose, UUID version 5
	// is recommended. See RFC 4122 for details.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId" json:"operation_id,omitempty"`
	// Fully qualified name of the operation. Reserved for future use.
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName" json:"operation_name,omitempty"`
	// Identity of the consumer who is using the service.
	// This field should be filled in for the operations initiated by a
	// consumer, but not for service-initiated operations that are
	// not related to a specific consumer.
	//
	// This can be in one of the following formats:
	//   project:<project_id>,
	//   project_number:<project_number>,
	//   api_key:<api_key>.
	ConsumerId string `protobuf:"bytes,3,opt,name=consumer_id,json=consumerId" json:"consumer_id,omitempty"`
	// Required. Start time of the operation.
	StartTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// End time of the operation.
	// Required when the operation is used in [ServiceController.Report][google.api.servicecontrol.v1.ServiceController.Report],
	// but optional when the operation is used in [ServiceController.Check][google.api.servicecontrol.v1.ServiceController.Check].
	EndTime *google_protobuf3.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Labels describing the operation. Only the following labels are allowed:
	//
	// - Labels describing monitored resources as defined in
	//   the service configuration.
	// - Default labels of metric values. When specified, labels defined in the
	//   metric value override these default.
	// - The following labels defined by Google Cloud Platform:
	//     - `cloud.googleapis.com/location` describing the location where the
	//        operation happened,
	//     - `servicecontrol.googleapis.com/user_agent` describing the user agent
	//        of the API request,
	//     - `servicecontrol.googleapis.com/service_agent` describing the service
	//        used to handle the API request (e.g. ESP),
	//     - `servicecontrol.googleapis.com/platform` describing the platform
	//        where the API is served (e.g. GAE, GCE, GKE).
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Represents information about this operation. Each MetricValueSet
	// corresponds to a metric defined in the service configuration.
	// The data type used in the MetricValueSet must agree with
	// the data type specified in the metric definition.
	//
	// Within a single operation, it is not allowed to have more than one
	// MetricValue instances that have the same metric names and identical
	// label value combinations. If a request has such duplicated MetricValue
	// instances, the entire request is rejected with
	// an invalid argument error.
	MetricValueSets []*MetricValueSet `protobuf:"bytes,7,rep,name=metric_value_sets,json=metricValueSets" json:"metric_value_sets,omitempty"`
	// Represents information to be logged.
	LogEntries []*LogEntry `protobuf:"bytes,8,rep,name=log_entries,json=logEntries" json:"log_entries,omitempty"`
	// DO NOT USE. This is an experimental field.
	Importance Operation_Importance `protobuf:"varint,11,opt,name=importance,enum=google.api.servicecontrol.v1.Operation_Importance" json:"importance,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Operation) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *Operation) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Operation) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *Operation) GetStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Operation) GetEndTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Operation) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Operation) GetMetricValueSets() []*MetricValueSet {
	if m != nil {
		return m.MetricValueSets
	}
	return nil
}

func (m *Operation) GetLogEntries() []*LogEntry {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *Operation) GetImportance() Operation_Importance {
	if m != nil {
		return m.Importance
	}
	return Operation_LOW
}

func init() {
	proto.RegisterType((*Operation)(nil), "google.api.servicecontrol.v1.Operation")
	proto.RegisterEnum("google.api.servicecontrol.v1.Operation_Importance", Operation_Importance_name, Operation_Importance_value)
}

func init() { proto.RegisterFile("google/api/servicecontrol/v1/operation.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x9d, 0xa6, 0xf9, 0xf5, 0x56, 0x63, 0x1c, 0x3c, 0x2c, 0xa1, 0x90, 0x58, 0x50, 0x72,
	0x28, 0xb3, 0x34, 0x45, 0xb0, 0x7a, 0x2b, 0x48, 0x1b, 0x8d, 0xb6, 0xac, 0xa2, 0xe2, 0x25, 0x4c,
	0x36, 0xcf, 0x65, 0x70, 0x77, 0x66, 0x99, 0x99, 0x04, 0x7a, 0xf6, 0xe2, 0x9f, 0xec, 0x51, 0x76,
	0xf6, 0x47, 0x13, 0x90, 0xb5, 0xb7, 0x7d, 0x8f, 0xef, 0xe7, 0xbb, 0xdf, 0x79, 0x6f, 0x06, 0x4e,
	0x62, 0xa5, 0xe2, 0x04, 0x03, 0x9e, 0x89, 0xc0, 0xa0, 0xde, 0x8a, 0x08, 0x23, 0x25, 0xad, 0x56,
	0x49, 0xb0, 0x3d, 0x0d, 0x54, 0x86, 0x9a, 0x5b, 0xa1, 0x24, 0xcb, 0xb4, 0xb2, 0x8a, 0x1e, 0x15,
	0x6a, 0xc6, 0x33, 0xc1, 0xf6, 0xd5, 0x6c, 0x7b, 0x3a, 0x3a, 0xda, 0xf1, 0xe2, 0x52, 0x2a, 0xeb,
	0x50, 0x53, 0xb0, 0xa3, 0xe6, 0x3f, 0x25, 0x2a, 0x5e, 0xa2, 0xb4, 0xfa, 0xb6, 0x54, 0x07, 0x8d,
	0xea, 0x14, 0xad, 0x16, 0xd1, 0x72, 0xcb, 0x93, 0x0d, 0x96, 0xc0, 0xb8, 0x04, 0x5c, 0xb5, 0xda,
	0xfc, 0x08, 0xac, 0x48, 0xd1, 0x58, 0x9e, 0x66, 0x85, 0xe0, 0xf8, 0x77, 0x1b, 0xfa, 0xd7, 0xd5,
	0x79, 0xe8, 0x33, 0x78, 0x58, 0x1f, 0x6e, 0x29, 0xd6, 0x3e, 0x99, 0x90, 0x69, 0x3f, 0xf4, 0xea,
	0xde, 0x7c, 0x4d, 0x9f, 0xc3, 0xe0, 0x4e, 0x22, 0x79, 0x8a, 0xfe, 0x81, 0x13, 0x3d, 0xaa, 0xbb,
	0x1f, 0x79, 0x8a, 0x74, 0x0c, 0x5e, 0xa4, 0xa4, 0xd9, 0xa4, 0xa8, 0x73, 0xa3, 0x96, 0xd3, 0x40,
	0xd5, 0x9a, 0xaf, 0xe9, 0x39, 0x80, 0xb1, 0x5c, 0xdb, 0x65, 0x9e, 0xc8, 0x3f, 0x9c, 0x90, 0xa9,
	0x37, 0x1b, 0xb1, 0x72, 0x92, 0x55, 0x5c, 0xf6, 0xb9, 0x8a, 0x1b, 0xf6, 0x9d, 0x3a, 0xaf, 0xe9,
	0x4b, 0xe8, 0xa1, 0x5c, 0x17, 0x60, 0xfb, 0xbf, 0x60, 0x17, 0xe5, 0xda, 0x61, 0xef, 0xa1, 0x93,
	0xf0, 0x15, 0x26, 0xc6, 0xef, 0x4c, 0x5a, 0x53, 0x6f, 0x76, 0xc6, 0x9a, 0xf6, 0xc6, 0xea, 0xa9,
	0xb0, 0x85, 0xa3, 0xde, 0xe6, 0x7b, 0x08, 0x4b, 0x0b, 0xfa, 0x0d, 0x9e, 0xec, 0x8e, 0x7b, 0x69,
	0xd0, 0x1a, 0xbf, 0xeb, 0x7c, 0x4f, 0x9a, 0x7d, 0x3f, 0x38, 0xec, 0x4b, 0x4e, 0x7d, 0x42, 0x1b,
	0x3e, 0x4e, 0xf7, 0x6a, 0x43, 0x2f, 0xc1, 0xab, 0xd6, 0x2e, 0xd0, 0xf8, 0x3d, 0xe7, 0xf9, 0xa2,
	0xd9, 0x73, 0xa1, 0xe2, 0x22, 0x1e, 0x24, 0xc5, 0x97, 0x40, 0x43, 0x43, 0x00, 0x91, 0x66, 0x4a,
	0x5b, 0x2e, 0x23, 0xf4, 0xbd, 0x09, 0x99, 0x0e, 0x66, 0xb3, 0xfb, 0x9e, 0x79, 0x5e, 0x93, 0xe1,
	0x8e, 0xcb, 0xe8, 0x1c, 0xbc, 0x9d, 0x69, 0xd0, 0x21, 0xb4, 0x7e, 0xe2, 0x6d, 0x79, 0x4d, 0xf2,
	0x4f, 0xfa, 0x14, 0xda, 0x6e, 0x20, 0xe5, 0xad, 0x28, 0x8a, 0xd7, 0x07, 0xaf, 0xc8, 0xf1, 0x18,
	0xe0, 0xce, 0x94, 0x76, 0xa1, 0xb5, 0xb8, 0xfe, 0x3a, 0x7c, 0x40, 0x7b, 0x70, 0x78, 0x35, 0xbf,
	0xbc, 0x1a, 0x92, 0x8b, 0x5f, 0x04, 0x26, 0x91, 0x4a, 0x1b, 0x13, 0x5e, 0x0c, 0xea, 0x88, 0x37,
	0xf9, 0xaa, 0x6f, 0xc8, 0xf7, 0x77, 0xa5, 0x3e, 0x56, 0x09, 0x97, 0x31, 0x53, 0x3a, 0x0e, 0x62,
	0x94, 0xee, 0x22, 0x94, 0xcf, 0x85, 0x67, 0xc2, 0xfc, 0xfb, 0xc9, 0xbc, 0xd9, 0xef, 0xfc, 0x21,
	0x64, 0xd5, 0x71, 0xe4, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x9c, 0xc2, 0x5d, 0x03,
	0x04, 0x00, 0x00,
}
