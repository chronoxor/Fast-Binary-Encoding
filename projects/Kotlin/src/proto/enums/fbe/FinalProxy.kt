// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.2.0.0

@file:Suppress("UnusedImport", "unused")

package enums.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import enums.*

// Fast Binary Encoding enums final proxy
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalProxy : fbe.Receiver
{
    // Proxy models accessors
    private val EnumsModel: EnumsFinalModel

    constructor() : super(true)
    {
        EnumsModel = EnumsFinalModel()
    }

    constructor(buffer: Buffer) : super(buffer, true)
    {
        EnumsModel = EnumsFinalModel()
    }

    // Proxy handlers
    protected open fun onProxy(model: EnumsFinalModel, type: Long, buffer: ByteArray, offset: Long, size: Long) {}

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            enums.fbe.EnumsFinalModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                EnumsModel.attach(buffer, offset)
                assert(EnumsModel.verify()) { "enums.Enums validation failed!" }

                // Call proxy handler
                onProxy(EnumsModel, type, buffer, offset, size)
                return true
            }
        }

        return false
    }
}
