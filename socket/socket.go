package socket

import (
	"fmt"
	"net"
	"strconv"
	"syscall"
)

type Socket struct {
	FileDescriptor int
}

func (socket *Socket) Read(bytes []byte) (int, error) {
	read, err := syscall.Read(socket.FileDescriptor, bytes)
	if err != nil {
		return 0, err
	}
	return read, err
}

func (socket *Socket) Write(bytes []byte) (int, error) {
	write, err := syscall.Write(socket.FileDescriptor, bytes)
	if err != nil {
		return 0, err
	}
	return write, err
}

func (socket *Socket) Close() error {
	return syscall.Close(socket.FileDescriptor)
}

func (soc *Socket) String() string {
	return strconv.Itoa(soc.FileDescriptor)
}

func Listen(ip string, port int) (*Socket, error) {
	socket := &Socket{}

	handle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to create socket (%v)", err)
	}

	socket.FileDescriptor = int(handle)

	socketAddress := &syscall.SockaddrInet4{Port: port}
	copy(socketAddress.Addr[:], net.ParseIP(ip))

	if err = syscall.Bind(socket.FileDescriptor, socketAddress); err != nil {
		return nil, fmt.Errorf("failed to bind socket (%v)", err)
	}

	if err = syscall.Listen(socket.FileDescriptor, syscall.SOMAXCONN); err != nil {
		return nil, fmt.Errorf("failed to bind socket (%v)", err)
	}

	return socket, nil
}
