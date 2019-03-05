// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.2.0.0

@file:Suppress("UnusedImport", "unused")

package protoex.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import protoex.*

// Fast Binary Encoding protoex final proxy
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalProxy : fbe.Receiver
{
    // Imported proxy
    var protoProxy: proto.fbe.FinalProxy? = null

    // Proxy models accessors
    private val OrderModel: OrderFinalModel
    private val BalanceModel: BalanceFinalModel
    private val AccountModel: AccountFinalModel

    constructor() : super(true)
    {
        protoProxy = proto.fbe.FinalProxy(buffer)
        OrderModel = OrderFinalModel()
        BalanceModel = BalanceFinalModel()
        AccountModel = AccountFinalModel()
    }

    constructor(buffer: Buffer) : super(buffer, true)
    {
        protoProxy = proto.fbe.FinalProxy(buffer)
        OrderModel = OrderFinalModel()
        BalanceModel = BalanceFinalModel()
        AccountModel = AccountFinalModel()
    }

    // Proxy handlers
    protected open fun onProxy(model: OrderFinalModel, type: Long, buffer: ByteArray, offset: Long, size: Long) {}
    protected open fun onProxy(model: BalanceFinalModel, type: Long, buffer: ByteArray, offset: Long, size: Long) {}
    protected open fun onProxy(model: AccountFinalModel, type: Long, buffer: ByteArray, offset: Long, size: Long) {}

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            protoex.fbe.OrderFinalModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                OrderModel.attach(buffer, offset)
                assert(OrderModel.verify()) { "protoex.Order validation failed!" }

                // Call proxy handler
                onProxy(OrderModel, type, buffer, offset, size)
                return true
            }
            protoex.fbe.BalanceFinalModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                BalanceModel.attach(buffer, offset)
                assert(BalanceModel.verify()) { "protoex.Balance validation failed!" }

                // Call proxy handler
                onProxy(BalanceModel, type, buffer, offset, size)
                return true
            }
            protoex.fbe.AccountFinalModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                AccountModel.attach(buffer, offset)
                assert(AccountModel.verify()) { "protoex.Account validation failed!" }

                // Call proxy handler
                onProxy(AccountModel, type, buffer, offset, size)
                return true
            }
        }

        if ((protoProxy != null) && protoProxy!!.onReceive(type, buffer, offset, size))
            return true

        return false
    }
}