package moegoapis_test

import (
	"os"
	"strings"
	"testing"

	appointmentpb "github.com/MoeGolibrary/moegoapis/genproto/go/business/appointment/v1/appointmentpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestAppointmentMutationSendAutoMessageContract(t *testing.T) {
	t.Parallel()

	file := appointmentpb.File_moego_business_appointment_v1_appointment_service_proto
	tests := []struct {
		message protoreflect.Name
		number  protoreflect.FieldNumber
	}{
		{message: "CreateAppointmentRequest", number: 6},
		{message: "RescheduleAppointmentRequest", number: 4},
		{message: "CancelAppointmentRequest", number: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.message), func(t *testing.T) {
			message := file.Messages().ByName(tt.message)
			if message == nil {
				t.Fatalf("message %s not found", tt.message)
			}

			field := message.Fields().ByName("send_auto_message")
			if field == nil {
				t.Fatalf("send_auto_message not found in %s", tt.message)
			}
			if field.Number() != tt.number {
				t.Fatalf("send_auto_message number = %d, want %d", field.Number(), tt.number)
			}
			if field.Kind() != protoreflect.BoolKind || !field.HasOptionalKeyword() {
				t.Fatalf("send_auto_message must be an optional bool; kind=%s optional=%t",
					field.Kind(), field.HasOptionalKeyword())
			}
			if old := message.Fields().ByName("auto_message_methods"); old != nil {
				t.Fatalf("legacy auto_message_methods still exists in %s", tt.message)
			}
		})
	}

	statusRequest := file.Messages().ByName("UpdateAppointmentStatusRequest")
	if statusRequest == nil || statusRequest.Fields().ByName("message_method") == nil {
		t.Fatal("UpdateAppointmentStatusRequest.message_method must remain in the contract")
	}
}

func TestAppointmentAutoMessageLegacyArtifactsRemoved(t *testing.T) {
	t.Parallel()

	for _, path := range []string{
		"moego/business/message/v1/message.proto",
		"genproto/go/business/message/v1/messagepb/message.pb.go",
	} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("legacy artifact %s still exists or could not be checked: %v", path, err)
		}
	}

	protoSource, err := os.ReadFile("moego/business/appointment/v1/appointment_service.proto")
	if err != nil {
		t.Fatal(err)
	}
	const comment = `// bool, Whether to send the corresponding auto message for this appointment.
  // Delivery methods (e.g. MSG, EMAIL) follow the business's auto-message settings
  // and the customer's communication preferences.
  // When false or not set, no auto message will be sent.`
	if got := strings.Count(string(protoSource), comment); got != 3 {
		t.Fatalf("frozen send_auto_message comment count = %d, want 3", got)
	}

	swagger, err := os.ReadFile("docs/openapi/moego/business/appointment/v1/appointment_service.swagger.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, legacy := range []string{"autoMessageMethods", "MessageDeliveryMethod"} {
		if strings.Contains(string(swagger), legacy) {
			t.Fatalf("generated Swagger still contains %q", legacy)
		}
	}
}
