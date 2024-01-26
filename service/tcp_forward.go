package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

type TcpForward struct {
	PortListener net.Listener
	SrcConn      net.Conn
	DestConn     net.Conn
}

func NewTcpForWard() *TcpForward {
	return &TcpForward{}
}

func (_self *TcpForward) doTcpForward(srcAddr string, destAddr string) (err error) {
	_self.PortListener, err = net.Listen("tcp", srcAddr)
	if err != nil {
		return err
	}
	_self.SrcConn, err = _self.PortListener.Accept()
	if err != nil {
		log.Error("Forward Accept err:", err)
		log.Error(fmt.Sprint("转发出现异常：", srcAddr, "->", destAddr))
		return err
	}

	_self.DestConn, err = net.DialTimeout("tcp", destAddr, 30*time.Second)
	if err != nil {
		log.Debug("转发出现异常 Forward to Dest Addr err:", err.Error())
		return err
	}

	for {
		go func() {
			_, err := io.Copy(_self.DestConn, _self.SrcConn)
			if err != nil {
				log.Error("客户端来源数据转发到目标端口异常：", err)
			}
		}()
		go func() {
			_, err := io.Copy(_self.SrcConn, _self.DestConn)
			if err != nil {
				log.Error("目标端口返回响应数据异常：", err)
			}
		}()

	}
	return nil
}
func (_self *TcpForward) Stop() {
	log.Debug("关闭一个连接：", _self.SrcConn.RemoteAddr(), " on ", _self.SrcConn.LocalAddr())
	_self.SrcConn.Close()
	_self.DestConn.Close()
}
