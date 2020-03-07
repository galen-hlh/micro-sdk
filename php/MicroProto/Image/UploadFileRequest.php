<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: image/image.proto

namespace MicroProto\Image;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>image.UploadFileRequest</code>
 */
class UploadFileRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bytes file = 1;</code>
     */
    private $file = '';
    /**
     * Generated from protobuf field <code>string fileName = 2;</code>
     */
    private $fileName = '';
    /**
     * Generated from protobuf field <code>uint32 fileType = 3;</code>
     */
    private $fileType = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $file
     *     @type string $fileName
     *     @type int $fileType
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Image\Image::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bytes file = 1;</code>
     * @return string
     */
    public function getFile()
    {
        return $this->file;
    }

    /**
     * Generated from protobuf field <code>bytes file = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setFile($var)
    {
        GPBUtil::checkString($var, False);
        $this->file = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string fileName = 2;</code>
     * @return string
     */
    public function getFileName()
    {
        return $this->fileName;
    }

    /**
     * Generated from protobuf field <code>string fileName = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setFileName($var)
    {
        GPBUtil::checkString($var, True);
        $this->fileName = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint32 fileType = 3;</code>
     * @return int
     */
    public function getFileType()
    {
        return $this->fileType;
    }

    /**
     * Generated from protobuf field <code>uint32 fileType = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setFileType($var)
    {
        GPBUtil::checkUint32($var);
        $this->fileType = $var;

        return $this;
    }

}
