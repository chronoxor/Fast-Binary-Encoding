// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.proto.fbe

// Fast Binary Encoding com.chronoxor.proto final client
@Suppress("MemberVisibilityCanBePrivate", "PropertyName")
open class FinalClient : com.chronoxor.fbe.Client, FinalReceiverListener
{
    // Client sender models accessors
    val OrderSenderModel: OrderFinalModel
    val BalanceSenderModel: BalanceFinalModel
    val AccountSenderModel: AccountFinalModel

    // Client receiver values accessors
    private val OrderReceiverValue: com.chronoxor.proto.Order
    private val BalanceReceiverValue: com.chronoxor.proto.Balance
    private val AccountReceiverValue: com.chronoxor.proto.Account

    // Client receiver models accessors
    private val OrderReceiverModel: OrderFinalModel
    private val BalanceReceiverModel: BalanceFinalModel
    private val AccountReceiverModel: AccountFinalModel

    constructor() : super(true)
    {
        OrderSenderModel = OrderFinalModel(sendBuffer)
        OrderReceiverValue = com.chronoxor.proto.Order()
        OrderReceiverModel = OrderFinalModel()
        BalanceSenderModel = BalanceFinalModel(sendBuffer)
        BalanceReceiverValue = com.chronoxor.proto.Balance()
        BalanceReceiverModel = BalanceFinalModel()
        AccountSenderModel = AccountFinalModel(sendBuffer)
        AccountReceiverValue = com.chronoxor.proto.Account()
        AccountReceiverModel = AccountFinalModel()
    }

    constructor(sendBuffer: com.chronoxor.fbe.Buffer, receiveBuffer: com.chronoxor.fbe.Buffer) : super(sendBuffer, receiveBuffer, true)
    {
        OrderSenderModel = OrderFinalModel(sendBuffer)
        OrderReceiverValue = com.chronoxor.proto.Order()
        OrderReceiverModel = OrderFinalModel()
        BalanceSenderModel = BalanceFinalModel(sendBuffer)
        BalanceReceiverValue = com.chronoxor.proto.Balance()
        BalanceReceiverModel = BalanceFinalModel()
        AccountSenderModel = AccountFinalModel(sendBuffer)
        AccountReceiverValue = com.chronoxor.proto.Account()
        AccountReceiverModel = AccountFinalModel()
    }

    @Suppress("JoinDeclarationAndAssignment")
    fun send(obj: Any): Long
    {
        when (obj)
        {
            is com.chronoxor.proto.Order -> return send(obj)
            is com.chronoxor.proto.Balance -> return send(obj)
            is com.chronoxor.proto.Account -> return send(obj)
        }

        return 0
    }

    fun send(value: com.chronoxor.proto.Order): Long
    {
        // Serialize the value into the FBE stream
        val serialized = OrderSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Order serialization failed!" }
        assert(OrderSenderModel.verify()) { "com.chronoxor.proto.Order validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.Balance): Long
    {
        // Serialize the value into the FBE stream
        val serialized = BalanceSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Balance serialization failed!" }
        assert(BalanceSenderModel.verify()) { "com.chronoxor.proto.Balance validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.Account): Long
    {
        // Serialize the value into the FBE stream
        val serialized = AccountSenderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Account serialization failed!" }
        assert(AccountSenderModel.verify()) { "com.chronoxor.proto.Account validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }

    // Send message handler
    override fun onSend(buffer: ByteArray, offset: Long, size: Long): Long { throw UnsupportedOperationException("com.chronoxor.proto.fbe.Client.onSend() not implemented!") }
    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: FinalReceiverListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            com.chronoxor.proto.fbe.OrderFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                OrderReceiverModel.attach(buffer, offset)
                assert(OrderReceiverModel.verify()) { "com.chronoxor.proto.Order validation failed!" }
                val deserialized = OrderReceiverModel.deserialize(OrderReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.proto.Order deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = OrderReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(OrderReceiverValue)
                return true
            }
            com.chronoxor.proto.fbe.BalanceFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                BalanceReceiverModel.attach(buffer, offset)
                assert(BalanceReceiverModel.verify()) { "com.chronoxor.proto.Balance validation failed!" }
                val deserialized = BalanceReceiverModel.deserialize(BalanceReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.proto.Balance deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = BalanceReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(BalanceReceiverValue)
                return true
            }
            com.chronoxor.proto.fbe.AccountFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                AccountReceiverModel.attach(buffer, offset)
                assert(AccountReceiverModel.verify()) { "com.chronoxor.proto.Account validation failed!" }
                val deserialized = AccountReceiverModel.deserialize(AccountReceiverValue)
                assert(deserialized > 0) { "com.chronoxor.proto.Account deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = AccountReceiverValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(AccountReceiverValue)
                return true
            }
        }

        return false
    }
}