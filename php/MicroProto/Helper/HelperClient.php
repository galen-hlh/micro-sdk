<?php
// GENERATED CODE -- DO NOT EDIT!

namespace MicroProto\Helper;

/**
 */
class HelperClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \MicroProto\Helper\IdRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetDistributeId(\MicroProto\Helper\IdRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/Helper/GetDistributeId',
        $argument,
        ['\MicroProto\Helper\IdResponse', 'decode'],
        $metadata, $options);
    }

}
