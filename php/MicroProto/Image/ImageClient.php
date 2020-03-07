<?php
// GENERATED CODE -- DO NOT EDIT!

namespace MicroProto\Image;

/**
 */
class ImageClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function UploadFile($metadata = [], $options = []) {
        return $this->_bidiRequest('/image.Image/UploadFile',
        ['\MicroProto\Image\UploadFileResponse','decode'],
        $metadata, $options);
    }

}
