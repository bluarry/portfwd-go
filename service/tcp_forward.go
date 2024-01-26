package service

import (
	"net"
)

type TcpForward struct {
	PortListener net.Listener
}

func NewTcpForWard() *TcpForward {
	return &TcpForward{}
}

//func (_self *TcpForward) doTcpForward(srcAddr string, destAddr string) (err error) {
//	_self.PortListener, err = net.Listen("tcp", srcAddr)
//	if err != nil {
//		return err
//	}
//
//	for {
//		realClientConn, err := _self.PortListener.Accept()
//		if err != nil {
//			logs.Error("Forward Accept err:", err.Error())
//			logs.Error(fmt.Sprint("转发出现异常：", _self.Config.SrcAddr, ":", _self.Config.SrcPort, "->", destAddr))
//			_self.StopJob()
//			break
//		}
//
//		if ForWardDebug == true {
//			logs.Info("新用户 ", realClientConn.RemoteAddr().String(), " 数据转发规则：", fmt.Sprint(_self.Config.SrcAddr, ":", _self.Config.SrcPort), "->", destAddr)
//		}
//
//		var destConn net.Conn
//		if _self.Config.Protocol == "UDP" {
//			//destConn, err = Common.DialKcpTimeout(destAddr, 100)
//			destConn, err = net.DialTimeout("UDP", destAddr, 30*time.Second)
//		} else {
//			destConn, err = net.DialTimeout("tcp", destAddr, 30*time.Second)
//		}
//
//		if err != nil {
//			if ForWardDebug == true {
//				logs.Warn("转发出现异常 Forward to Dest Addr err:", err.Error())
//			}
//
//			//break
//			continue
//
//		}
//
//		forwardClient := &ForWardClient{realClientConn, destConn, nil, _self.ClosedCallBack}
//
//		if Utils.IsNotEmpty(_self.Config.Others) {
//			var dispatchConns []io.Writer
//			//分发方式
//			dispatchTargets := Utils.Split(_self.Config.Others, ";")
//
//			for _, dispatchTarget := range dispatchTargets {
//				logs.Debug("分发到：", dispatchTarget)
//				dispatchTargetConn, err := net.DialTimeout("tcp", dispatchTarget, 30*time.Second)
//				if err == nil {
//					dispatchConns = append(dispatchConns, dispatchTargetConn)
//				}
//
//			}
//
//			forwardClient.DispatchConns = dispatchConns
//
//			go forwardClient.DispatchData(dispatchConns)
//		} else {
//			go forwardClient.StartForward()
//		}
//
//		_self.RegistryClient(_self.GetClientId(realClientConn), forwardClient)
//		//_self.RegistryClient(fmt.Sprint(sourceAddr, "_", "TCP", "_", id), forwardClient)
//
//	}
//	return nil
//}
