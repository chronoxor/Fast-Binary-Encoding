// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.2.0.0

package proto.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.ByteBuffer;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import proto.*;

// Fast Binary Encoding proto proxy
public class Proxy extends fbe.Receiver
{
    // Proxy models accessors
    private final OrderModel OrderModel;
    private final BalanceModel BalanceModel;
    private final AccountModel AccountModel;

    public Proxy()
    {
        super(false);
        OrderModel = new OrderModel();
        BalanceModel = new BalanceModel();
        AccountModel = new AccountModel();
    }
    public Proxy(Buffer buffer)
    {
        super(buffer, false);
        OrderModel = new OrderModel();
        BalanceModel = new BalanceModel();
        AccountModel = new AccountModel();
    }

    // Proxy handlers
    protected void onProxy(OrderModel model, long type, byte[] buffer, long offset, long size) {}
    protected void onProxy(BalanceModel model, long type, byte[] buffer, long offset, long size) {}
    protected void onProxy(AccountModel model, long type, byte[] buffer, long offset, long size) {}

    @Override
    public boolean onReceive(long type, byte[] buffer, long offset, long size)
    {
        switch ((int)type)
        {
            case (int)proto.fbe.OrderModel.fbeTypeConst:
            {
                // Attach the FBE stream to the proxy model
                OrderModel.attach(buffer, offset);
                assert OrderModel.verify() : "proto.Order validation failed!";

                // Call proxy handler
                onProxy(OrderModel, type, buffer, offset, size);
                return true;
            }
            case (int)proto.fbe.BalanceModel.fbeTypeConst:
            {
                // Attach the FBE stream to the proxy model
                BalanceModel.attach(buffer, offset);
                assert BalanceModel.verify() : "proto.Balance validation failed!";

                // Call proxy handler
                onProxy(BalanceModel, type, buffer, offset, size);
                return true;
            }
            case (int)proto.fbe.AccountModel.fbeTypeConst:
            {
                // Attach the FBE stream to the proxy model
                AccountModel.attach(buffer, offset);
                assert AccountModel.verify() : "proto.Account validation failed!";

                // Call proxy handler
                onProxy(AccountModel, type, buffer, offset, size);
                return true;
            }
        }

        return false;
    }
}
