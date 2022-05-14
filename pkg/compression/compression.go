package compression

type Codec interface {
	Packer
	Unpacker
}

type Packer interface {
	Pack(string) ([]byte, error)
}

type Unpacker interface {
	Unpack([]byte) (string, error)
}
