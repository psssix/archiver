package compression

type Packer interface {
	Pack(string) ([]byte, error)
}

type Unpacker interface {
	Unpack([]byte) (string, error)
}
