/**
 * @fileoverview gRPC-Web generated client stub for diver
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.diver = require('./server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.diver.addNumServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.diver.addNumServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.diver.addNumParams,
 *   !proto.diver.totalNum>}
 */
const methodInfo_addNumService_addNum = new grpc.web.AbstractClientBase.MethodInfo(
  proto.diver.totalNum,
  /** @param {!proto.diver.addNumParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.diver.totalNum.deserializeBinary
);


/**
 * @param {!proto.diver.addNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.diver.totalNum)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.diver.totalNum>|undefined}
 *     The XHR Node Readable Stream
 */
proto.diver.addNumServiceClient.prototype.addNum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/diver.addNumService/addNum',
      request,
      metadata || {},
      methodInfo_addNumService_addNum,
      callback);
};


/**
 * @param {!proto.diver.addNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.diver.totalNum>}
 *     A native promise that resolves to the response
 */
proto.diver.addNumServicePromiseClient.prototype.addNum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/diver.addNumService/addNum',
      request,
      metadata || {},
      methodInfo_addNumService_addNum);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.diver.getTotalNumParams,
 *   !proto.diver.totalNum>}
 */
const methodInfo_addNumService_getTotalNum = new grpc.web.AbstractClientBase.MethodInfo(
  proto.diver.totalNum,
  /** @param {!proto.diver.getTotalNumParams} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.diver.totalNum.deserializeBinary
);


/**
 * @param {!proto.diver.getTotalNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.diver.totalNum)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.diver.totalNum>|undefined}
 *     The XHR Node Readable Stream
 */
proto.diver.addNumServiceClient.prototype.getTotalNum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/diver.addNumService/getTotalNum',
      request,
      metadata || {},
      methodInfo_addNumService_getTotalNum,
      callback);
};


/**
 * @param {!proto.diver.getTotalNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.diver.totalNum>}
 *     A native promise that resolves to the response
 */
proto.diver.addNumServicePromiseClient.prototype.getTotalNum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/diver.addNumService/getTotalNum',
      request,
      metadata || {},
      methodInfo_addNumService_getTotalNum);
};


module.exports = proto.diver;

