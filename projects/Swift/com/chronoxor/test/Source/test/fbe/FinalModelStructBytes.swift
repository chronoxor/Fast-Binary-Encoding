// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0


import fbe

import proto

// Fast Binary Encoding StructBytes final model
public class FinalModelStructBytes: FinalModel {
    public var _buffer: Buffer
    public var _offset: Int

    let f1: FinalModelData
    let f2: FinalModelOptionalData
    let f3: FinalModelOptionalData

    // Field type
    public var fbeType: Int = fbeTypeConst

    public static let fbeTypeConst: Int = 120

    public required init(buffer: Buffer, offset: Int) {
        _buffer = buffer
        _offset = offset

        f1 = FinalModelData(buffer: buffer, offset: 0)
        f2 = FinalModelOptionalData(buffer: buffer, offset: 0)
        f3 = FinalModelOptionalData(buffer: buffer, offset: 0)
    }

    // Get the allocation size
    public func fbeAllocationSize(value fbeValue: StructBytes) -> Int {
        return 0
            + f1.fbeAllocationSize(value: fbeValue.f1)
            + f2.fbeAllocationSize(value: fbeValue.f2)
            + f3.fbeAllocationSize(value: fbeValue.f3)
    }

    // Check if the struct value is valid
    public func verify() -> Int {
        _buffer.shift(offset: fbeOffset)
        let fbeResult = verifyFields()
        _buffer.unshift(offset: fbeOffset)
        return fbeResult
    }

    // Check if the struct fields are valid
    public func verifyFields() -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeFieldSize: Int = 0

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f1.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f2.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f3.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        return fbeCurrentOffset
    }

    // Get the struct value
    public func get(size: inout Size) -> StructBytes {
        var fbeValue = StructBytes()
        return get(size: &size, fbeValue: &fbeValue)
    }

    // Get the struct value
    public func get(size: inout Size, fbeValue: inout StructBytes) -> StructBytes {
        _buffer.shift(offset: fbeOffset)
        size.value = getFields(fbeValue: &fbeValue)
        _buffer.unshift(offset: fbeOffset)
        return fbeValue
    }

    // Get the struct fields values
    public func getFields(fbeValue: inout StructBytes) -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeCurrentSize: Int = 0
        var fbeFieldSize = Size()

        f1.fbeOffset = fbeCurrentOffset
        fbeValue.f1 = f1.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeValue.f2 = f2.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeValue.f3 = f3.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }

    // Set the struct value
    public func set(value fbeValue: StructBytes) throws -> Int {
        _buffer.shift(offset: fbeOffset)
        let fbeSize = try setFields(fbeValue: fbeValue)
        _buffer.unshift(offset: fbeOffset)
        return fbeSize
    }

    // Set the struct fields values
    public func setFields(fbeValue: StructBytes) throws -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeCurrentSize: Int = 0
        let fbeFieldSize = Size()

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try f1.set(value: fbeValue.f1)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try f2.set(value: fbeValue.f2)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try f3.set(value: fbeValue.f3)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }
}