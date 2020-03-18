<?php
// GENERATED CODE -- DO NOT EDIT!

namespace MicroProto\IdProduce;

/**
 */
class IdProduceClient extends \Grpc\BaseStub {

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
    public function GetDistributeId($metadata = [], $options = []) {
        return $this->_bidiRequest('/idProduce.IdProduce/GetDistributeId',
        ['\MicroProto\IdProduce\IdProduceResponse','decode'],
        $metadata, $options);
    }

}
