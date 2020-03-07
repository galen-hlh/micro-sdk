<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: idProduce/idProduce.proto

namespace MicroProto\IdProduce;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>idProduce.IdProduceResponse</code>
 */
class IdProduceResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>uint32 code = 1;</code>
     */
    private $code = 0;
    /**
     * Generated from protobuf field <code>repeated uint64 ids = 2;</code>
     */
    private $ids;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $code
     *     @type int[]|string[]|\Google\Protobuf\Internal\RepeatedField $ids
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\IdProduce\IdProduce::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>uint32 code = 1;</code>
     * @return int
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>uint32 code = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkUint32($var);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated uint64 ids = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getIds()
    {
        return $this->ids;
    }

    /**
     * Generated from protobuf field <code>repeated uint64 ids = 2;</code>
     * @param int[]|string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setIds($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::UINT64);
        $this->ids = $arr;

        return $this;
    }

}

