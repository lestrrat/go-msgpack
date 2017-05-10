package msgpack

// Auto-generated by internal/cmd/genmap/genmap.go. DO NOT EDIT!

import "github.com/pkg/errors"

func (e *Encoder) encodeMapBool(in interface{}) error {
	for k, v := range in.(map[string]bool) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapInt(in interface{}) error {
	for k, v := range in.(map[string]int) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapInt8(in interface{}) error {
	for k, v := range in.(map[string]int8) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapInt16(in interface{}) error {
	for k, v := range in.(map[string]int16) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapInt32(in interface{}) error {
	for k, v := range in.(map[string]int32) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapInt64(in interface{}) error {
	for k, v := range in.(map[string]int64) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapUint(in interface{}) error {
	for k, v := range in.(map[string]uint) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapUint8(in interface{}) error {
	for k, v := range in.(map[string]uint8) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapUint16(in interface{}) error {
	for k, v := range in.(map[string]uint16) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapUint32(in interface{}) error {
	for k, v := range in.(map[string]uint32) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapUint64(in interface{}) error {
	for k, v := range in.(map[string]uint64) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapFloat32(in interface{}) error {
	for k, v := range in.(map[string]float32) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapFloat64(in interface{}) error {
	for k, v := range in.(map[string]float64) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}

func (e *Encoder) encodeMapString(in interface{}) error {
	for k, v := range in.(map[string]string) {
		if err := e.EncodeString(k); err != nil {
			return errors.Wrap(err, `failed to encode key`)
		}
		if err := e.Encode(v); err != nil {
			return errors.Wrapf(err, `failed to encode value for %s`, k)
		}
	}
	return nil
}