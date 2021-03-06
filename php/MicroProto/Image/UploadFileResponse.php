<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: image/image.proto

namespace MicroProto\Image;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>image.UploadFileResponse</code>
 */
class UploadFileResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>uint32 code = 1;</code>
     */
    private $code = 0;
    /**
     * Generated from protobuf field <code>string filePath = 2;</code>
     */
    private $filePath = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $code
     *     @type string $filePath
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Image\Image::initOnce();
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
     * Generated from protobuf field <code>string filePath = 2;</code>
     * @return string
     */
    public function getFilePath()
    {
        return $this->filePath;
    }

    /**
     * Generated from protobuf field <code>string filePath = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setFilePath($var)
    {
        GPBUtil::checkString($var, True);
        $this->filePath = $var;

        return $this;
    }

}

