// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.2.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Proxy StructSimple interface
type OnProxyStructSimple interface {
    OnProxyStructSimple(model *StructSimpleModel, fbeType int, buffer []byte)
}

// Proxy StructSimple function
type OnProxyStructSimpleFunc func(model *StructSimpleModel, fbeType int, buffer []byte)
func (f OnProxyStructSimpleFunc) OnProxyStructSimple(model *StructSimpleModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructOptional interface
type OnProxyStructOptional interface {
    OnProxyStructOptional(model *StructOptionalModel, fbeType int, buffer []byte)
}

// Proxy StructOptional function
type OnProxyStructOptionalFunc func(model *StructOptionalModel, fbeType int, buffer []byte)
func (f OnProxyStructOptionalFunc) OnProxyStructOptional(model *StructOptionalModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructNested interface
type OnProxyStructNested interface {
    OnProxyStructNested(model *StructNestedModel, fbeType int, buffer []byte)
}

// Proxy StructNested function
type OnProxyStructNestedFunc func(model *StructNestedModel, fbeType int, buffer []byte)
func (f OnProxyStructNestedFunc) OnProxyStructNested(model *StructNestedModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructBytes interface
type OnProxyStructBytes interface {
    OnProxyStructBytes(model *StructBytesModel, fbeType int, buffer []byte)
}

// Proxy StructBytes function
type OnProxyStructBytesFunc func(model *StructBytesModel, fbeType int, buffer []byte)
func (f OnProxyStructBytesFunc) OnProxyStructBytes(model *StructBytesModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructArray interface
type OnProxyStructArray interface {
    OnProxyStructArray(model *StructArrayModel, fbeType int, buffer []byte)
}

// Proxy StructArray function
type OnProxyStructArrayFunc func(model *StructArrayModel, fbeType int, buffer []byte)
func (f OnProxyStructArrayFunc) OnProxyStructArray(model *StructArrayModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructVector interface
type OnProxyStructVector interface {
    OnProxyStructVector(model *StructVectorModel, fbeType int, buffer []byte)
}

// Proxy StructVector function
type OnProxyStructVectorFunc func(model *StructVectorModel, fbeType int, buffer []byte)
func (f OnProxyStructVectorFunc) OnProxyStructVector(model *StructVectorModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructList interface
type OnProxyStructList interface {
    OnProxyStructList(model *StructListModel, fbeType int, buffer []byte)
}

// Proxy StructList function
type OnProxyStructListFunc func(model *StructListModel, fbeType int, buffer []byte)
func (f OnProxyStructListFunc) OnProxyStructList(model *StructListModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructSet interface
type OnProxyStructSet interface {
    OnProxyStructSet(model *StructSetModel, fbeType int, buffer []byte)
}

// Proxy StructSet function
type OnProxyStructSetFunc func(model *StructSetModel, fbeType int, buffer []byte)
func (f OnProxyStructSetFunc) OnProxyStructSet(model *StructSetModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructMap interface
type OnProxyStructMap interface {
    OnProxyStructMap(model *StructMapModel, fbeType int, buffer []byte)
}

// Proxy StructMap function
type OnProxyStructMapFunc func(model *StructMapModel, fbeType int, buffer []byte)
func (f OnProxyStructMapFunc) OnProxyStructMap(model *StructMapModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructHash interface
type OnProxyStructHash interface {
    OnProxyStructHash(model *StructHashModel, fbeType int, buffer []byte)
}

// Proxy StructHash function
type OnProxyStructHashFunc func(model *StructHashModel, fbeType int, buffer []byte)
func (f OnProxyStructHashFunc) OnProxyStructHash(model *StructHashModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Proxy StructHashEx interface
type OnProxyStructHashEx interface {
    OnProxyStructHashEx(model *StructHashExModel, fbeType int, buffer []byte)
}

// Proxy StructHashEx function
type OnProxyStructHashExFunc func(model *StructHashExModel, fbeType int, buffer []byte)
func (f OnProxyStructHashExFunc) OnProxyStructHashEx(model *StructHashExModel, fbeType int, buffer []byte) {
    f(model, fbeType, buffer)
}

// Fast Binary Encoding test proxy
type Proxy struct {
    *fbe.Receiver
    protoProxy *proto.Proxy
    structSimpleModel *StructSimpleModel
    structOptionalModel *StructOptionalModel
    structNestedModel *StructNestedModel
    structBytesModel *StructBytesModel
    structArrayModel *StructArrayModel
    structVectorModel *StructVectorModel
    structListModel *StructListModel
    structSetModel *StructSetModel
    structMapModel *StructMapModel
    structHashModel *StructHashModel
    structHashExModel *StructHashExModel

    // Proxy StructSimple handler
    HandlerOnProxyStructSimple OnProxyStructSimple
    // Proxy StructOptional handler
    HandlerOnProxyStructOptional OnProxyStructOptional
    // Proxy StructNested handler
    HandlerOnProxyStructNested OnProxyStructNested
    // Proxy StructBytes handler
    HandlerOnProxyStructBytes OnProxyStructBytes
    // Proxy StructArray handler
    HandlerOnProxyStructArray OnProxyStructArray
    // Proxy StructVector handler
    HandlerOnProxyStructVector OnProxyStructVector
    // Proxy StructList handler
    HandlerOnProxyStructList OnProxyStructList
    // Proxy StructSet handler
    HandlerOnProxyStructSet OnProxyStructSet
    // Proxy StructMap handler
    HandlerOnProxyStructMap OnProxyStructMap
    // Proxy StructHash handler
    HandlerOnProxyStructHash OnProxyStructHash
    // Proxy StructHashEx handler
    HandlerOnProxyStructHashEx OnProxyStructHashEx
}

// Create a new test proxy with an empty buffer
func NewProxy() *Proxy {
    return NewProxyWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new test proxy with the given buffer
func NewProxyWithBuffer(buffer *fbe.Buffer) *Proxy {
    proxy := &Proxy{
        fbe.NewReceiver(buffer, false),
        proto.NewProxyWithBuffer(buffer),
        NewStructSimpleModel(buffer),
        NewStructOptionalModel(buffer),
        NewStructNestedModel(buffer),
        NewStructBytesModel(buffer),
        NewStructArrayModel(buffer),
        NewStructVectorModel(buffer),
        NewStructListModel(buffer),
        NewStructSetModel(buffer),
        NewStructMapModel(buffer),
        NewStructHashModel(buffer),
        NewStructHashExModel(buffer),
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
    }
    proxy.SetupHandlerOnReceive(proxy)
    proxy.SetupHandlerOnProxyStructSimpleFunc(func(model *StructSimpleModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructOptionalFunc(func(model *StructOptionalModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructNestedFunc(func(model *StructNestedModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructBytesFunc(func(model *StructBytesModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructArrayFunc(func(model *StructArrayModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructVectorFunc(func(model *StructVectorModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructListFunc(func(model *StructListModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructSetFunc(func(model *StructSetModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructMapFunc(func(model *StructMapModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructHashFunc(func(model *StructHashModel, fbeType int, buffer []byte) {})
    proxy.SetupHandlerOnProxyStructHashExFunc(func(model *StructHashExModel, fbeType int, buffer []byte) {})
    return proxy
}

// Imported proxy

// Get the proto proxy
func (p *Proxy) ProtoProxy() *proto.Proxy { return p.protoProxy }
// Set the proto proxy
func (p *Proxy) SetProtoProxy(proxy *proto.Proxy) { p.protoProxy = proxy }

// Setup handlers
func (p *Proxy) SetupHandlers(handlers interface{}) {
    p.Receiver.SetupHandlers(handlers)
    p.protoProxy.SetupHandlers(handlers)
    if handler, ok := handlers.(OnProxyStructSimple); ok {
        p.SetupHandlerOnProxyStructSimple(handler)
    }
    if handler, ok := handlers.(OnProxyStructOptional); ok {
        p.SetupHandlerOnProxyStructOptional(handler)
    }
    if handler, ok := handlers.(OnProxyStructNested); ok {
        p.SetupHandlerOnProxyStructNested(handler)
    }
    if handler, ok := handlers.(OnProxyStructBytes); ok {
        p.SetupHandlerOnProxyStructBytes(handler)
    }
    if handler, ok := handlers.(OnProxyStructArray); ok {
        p.SetupHandlerOnProxyStructArray(handler)
    }
    if handler, ok := handlers.(OnProxyStructVector); ok {
        p.SetupHandlerOnProxyStructVector(handler)
    }
    if handler, ok := handlers.(OnProxyStructList); ok {
        p.SetupHandlerOnProxyStructList(handler)
    }
    if handler, ok := handlers.(OnProxyStructSet); ok {
        p.SetupHandlerOnProxyStructSet(handler)
    }
    if handler, ok := handlers.(OnProxyStructMap); ok {
        p.SetupHandlerOnProxyStructMap(handler)
    }
    if handler, ok := handlers.(OnProxyStructHash); ok {
        p.SetupHandlerOnProxyStructHash(handler)
    }
    if handler, ok := handlers.(OnProxyStructHashEx); ok {
        p.SetupHandlerOnProxyStructHashEx(handler)
    }
}

// Setup proxy StructSimple handler
func (p *Proxy) SetupHandlerOnProxyStructSimple(handler OnProxyStructSimple) { p.HandlerOnProxyStructSimple = handler }
// Setup proxy StructSimple handler function
func (p *Proxy) SetupHandlerOnProxyStructSimpleFunc(function func(model *StructSimpleModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructSimple = OnProxyStructSimpleFunc(function) }
// Setup proxy StructOptional handler
func (p *Proxy) SetupHandlerOnProxyStructOptional(handler OnProxyStructOptional) { p.HandlerOnProxyStructOptional = handler }
// Setup proxy StructOptional handler function
func (p *Proxy) SetupHandlerOnProxyStructOptionalFunc(function func(model *StructOptionalModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructOptional = OnProxyStructOptionalFunc(function) }
// Setup proxy StructNested handler
func (p *Proxy) SetupHandlerOnProxyStructNested(handler OnProxyStructNested) { p.HandlerOnProxyStructNested = handler }
// Setup proxy StructNested handler function
func (p *Proxy) SetupHandlerOnProxyStructNestedFunc(function func(model *StructNestedModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructNested = OnProxyStructNestedFunc(function) }
// Setup proxy StructBytes handler
func (p *Proxy) SetupHandlerOnProxyStructBytes(handler OnProxyStructBytes) { p.HandlerOnProxyStructBytes = handler }
// Setup proxy StructBytes handler function
func (p *Proxy) SetupHandlerOnProxyStructBytesFunc(function func(model *StructBytesModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructBytes = OnProxyStructBytesFunc(function) }
// Setup proxy StructArray handler
func (p *Proxy) SetupHandlerOnProxyStructArray(handler OnProxyStructArray) { p.HandlerOnProxyStructArray = handler }
// Setup proxy StructArray handler function
func (p *Proxy) SetupHandlerOnProxyStructArrayFunc(function func(model *StructArrayModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructArray = OnProxyStructArrayFunc(function) }
// Setup proxy StructVector handler
func (p *Proxy) SetupHandlerOnProxyStructVector(handler OnProxyStructVector) { p.HandlerOnProxyStructVector = handler }
// Setup proxy StructVector handler function
func (p *Proxy) SetupHandlerOnProxyStructVectorFunc(function func(model *StructVectorModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructVector = OnProxyStructVectorFunc(function) }
// Setup proxy StructList handler
func (p *Proxy) SetupHandlerOnProxyStructList(handler OnProxyStructList) { p.HandlerOnProxyStructList = handler }
// Setup proxy StructList handler function
func (p *Proxy) SetupHandlerOnProxyStructListFunc(function func(model *StructListModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructList = OnProxyStructListFunc(function) }
// Setup proxy StructSet handler
func (p *Proxy) SetupHandlerOnProxyStructSet(handler OnProxyStructSet) { p.HandlerOnProxyStructSet = handler }
// Setup proxy StructSet handler function
func (p *Proxy) SetupHandlerOnProxyStructSetFunc(function func(model *StructSetModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructSet = OnProxyStructSetFunc(function) }
// Setup proxy StructMap handler
func (p *Proxy) SetupHandlerOnProxyStructMap(handler OnProxyStructMap) { p.HandlerOnProxyStructMap = handler }
// Setup proxy StructMap handler function
func (p *Proxy) SetupHandlerOnProxyStructMapFunc(function func(model *StructMapModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructMap = OnProxyStructMapFunc(function) }
// Setup proxy StructHash handler
func (p *Proxy) SetupHandlerOnProxyStructHash(handler OnProxyStructHash) { p.HandlerOnProxyStructHash = handler }
// Setup proxy StructHash handler function
func (p *Proxy) SetupHandlerOnProxyStructHashFunc(function func(model *StructHashModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructHash = OnProxyStructHashFunc(function) }
// Setup proxy StructHashEx handler
func (p *Proxy) SetupHandlerOnProxyStructHashEx(handler OnProxyStructHashEx) { p.HandlerOnProxyStructHashEx = handler }
// Setup proxy StructHashEx handler function
func (p *Proxy) SetupHandlerOnProxyStructHashExFunc(function func(model *StructHashExModel, fbeType int, buffer []byte)) { p.HandlerOnProxyStructHashEx = OnProxyStructHashExFunc(function) }

// Receive message handler
func (p *Proxy) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    case p.structSimpleModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structSimpleModel.Buffer().Attach(buffer)
        if !p.structSimpleModel.Verify() {
            return false, errors.New("test.StructSimple validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructSimple.OnProxyStructSimple(p.structSimpleModel, fbeType, buffer)
        return true, nil
    case p.structOptionalModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structOptionalModel.Buffer().Attach(buffer)
        if !p.structOptionalModel.Verify() {
            return false, errors.New("test.StructOptional validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructOptional.OnProxyStructOptional(p.structOptionalModel, fbeType, buffer)
        return true, nil
    case p.structNestedModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structNestedModel.Buffer().Attach(buffer)
        if !p.structNestedModel.Verify() {
            return false, errors.New("test.StructNested validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructNested.OnProxyStructNested(p.structNestedModel, fbeType, buffer)
        return true, nil
    case p.structBytesModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structBytesModel.Buffer().Attach(buffer)
        if !p.structBytesModel.Verify() {
            return false, errors.New("test.StructBytes validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructBytes.OnProxyStructBytes(p.structBytesModel, fbeType, buffer)
        return true, nil
    case p.structArrayModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structArrayModel.Buffer().Attach(buffer)
        if !p.structArrayModel.Verify() {
            return false, errors.New("test.StructArray validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructArray.OnProxyStructArray(p.structArrayModel, fbeType, buffer)
        return true, nil
    case p.structVectorModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structVectorModel.Buffer().Attach(buffer)
        if !p.structVectorModel.Verify() {
            return false, errors.New("test.StructVector validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructVector.OnProxyStructVector(p.structVectorModel, fbeType, buffer)
        return true, nil
    case p.structListModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structListModel.Buffer().Attach(buffer)
        if !p.structListModel.Verify() {
            return false, errors.New("test.StructList validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructList.OnProxyStructList(p.structListModel, fbeType, buffer)
        return true, nil
    case p.structSetModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structSetModel.Buffer().Attach(buffer)
        if !p.structSetModel.Verify() {
            return false, errors.New("test.StructSet validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructSet.OnProxyStructSet(p.structSetModel, fbeType, buffer)
        return true, nil
    case p.structMapModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structMapModel.Buffer().Attach(buffer)
        if !p.structMapModel.Verify() {
            return false, errors.New("test.StructMap validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructMap.OnProxyStructMap(p.structMapModel, fbeType, buffer)
        return true, nil
    case p.structHashModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structHashModel.Buffer().Attach(buffer)
        if !p.structHashModel.Verify() {
            return false, errors.New("test.StructHash validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructHash.OnProxyStructHash(p.structHashModel, fbeType, buffer)
        return true, nil
    case p.structHashExModel.FBEType():
        // Attach the FBE stream to the proxy model
        p.structHashExModel.Buffer().Attach(buffer)
        if !p.structHashExModel.Verify() {
            return false, errors.New("test.StructHashEx validation failed")
        }

        // Call proxy handler
        p.HandlerOnProxyStructHashEx.OnProxyStructHashEx(p.structHashExModel, fbeType, buffer)
        return true, nil
    }

    if p.protoProxy != nil {
        if ok, err := p.protoProxy.OnReceive(fbeType, buffer); ok {
            return ok, err
        }
    }

    return false, nil
}
