package enum

import (
	"fmt"
	"reflect"

	apiv1 "github.com/fi-ts/api/go/fits/api/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// GetStringValue from a enum if the Enum is annotated with a enum_string_value
func GetStringValue[E protoreflect.Enum](e E) (*string, error) {
	if reflect.ValueOf(e).Kind() == reflect.Pointer {
		if reflect.ValueOf(e).IsNil() {
			return nil, fmt.Errorf("given enum is a nil pointer")
		}
	}
	if e.Descriptor().Values().Len() <= int(e.Number()) {
		return nil, fmt.Errorf("given enum number:%d is out of range", e.Number())
	}
	value := e.Descriptor().Values().Get(int(e.Number()))

	// Retrieve custom options
	opts := value.Options()

	// Check if the option is present
	if opts != nil && proto.HasExtension(opts, apiv1.E_EnumStringValue) {
		stringValue, ok := proto.GetExtension(opts, apiv1.E_EnumStringValue).(string)
		if !ok {
			return nil, fmt.Errorf("cast to stringvalue not possible: %v", e)
		}
		return &stringValue, nil
	}

	return nil, fmt.Errorf("unable to fetch stringvalue from %v", e)
}

// GetEnum of given Type and string if annotated with a enum_string_value
func GetEnum[E protoreflect.Enum](s string) (e E, err error) {

	var (
		realE    E
		allenums []protoreflect.EnumType
	)
	protoregistry.GlobalTypes.RangeEnums(func(et protoreflect.EnumType) bool {
		if et.Descriptor().FullName() == realE.Descriptor().FullName() {
			allenums = append(allenums, et)
		}
		return true
	})

	if len(allenums) != 1 {
		return realE, fmt.Errorf("no matching enum of type:%s for %q found", realE.Descriptor().FullName(), s)
	}

	targetEnum := allenums[0]

	for i := range realE.Descriptor().Values().Len() {
		value := realE.Descriptor().Values().Get(i)

		// Retrieve custom options
		opts := value.Options()

		// Check if the option is present
		if opts != nil && proto.HasExtension(opts, apiv1.E_EnumStringValue) {
			stringValue, ok := proto.GetExtension(opts, apiv1.E_EnumStringValue).(string)
			if !ok {
				return realE, fmt.Errorf("cast to stringvalue not possible: %v", e)
			}
			if stringValue == s {
				return targetEnum.New(protoreflect.EnumNumber(i)).(E), nil
			}
		}
	}

	return realE, fmt.Errorf("no matching enum of type:%s for %q found", realE.Descriptor().FullName(), s)
}
