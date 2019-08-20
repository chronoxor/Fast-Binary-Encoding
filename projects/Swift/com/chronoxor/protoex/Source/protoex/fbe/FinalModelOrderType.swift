// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0


import fbe

// Fast Binary Encoding OrderType final model
public class FinalModelOrderType: FinalModel {

    public var _buffer: Buffer
    public var _offset: Int

    // Final size
    public let fbeSize: Int = 1

    public init(buffer: Buffer = Buffer(), offset: Int = 0) {
        _buffer = buffer
        _offset = offset
    }

    // Get the allocation size
    public func fbeAllocationSize(value: OrderType) -> Int { fbeSize }

    public func verify() -> Int  {
        if _buffer.offset + fbeOffset + fbeSize > _buffer.size {
            return Int.max
        }

        return fbeSize
    }

    // Get the value
    public func get(size: inout Size) -> OrderType {
        if _buffer.offset + fbeOffset + fbeSize > _buffer.size {
            return OrderType()
        }

        size.value = fbeSize
        return OrderType(value: readByte(offset: fbeOffset))
    }

    // Set the value
    public func set(value: OrderType) -> Int {
        if _buffer.offset + fbeOffset + fbeSize > _buffer.size {
            assertionFailure("Model is broken!")
            return 0
        }

        write(offset: fbeOffset, value: value.raw)
        return fbeSize
    }
}