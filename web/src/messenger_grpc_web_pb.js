/**
 * @fileoverview gRPC-Web generated client stub for messenger
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.28.3
// source: messenger.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.messenger = require('./messenger_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.messenger.MessengerServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.messenger.MessengerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.messenger.User,
 *   !proto.messenger.User>}
 */
const methodDescriptor_MessengerService_RegisterUser = new grpc.web.MethodDescriptor(
  '/messenger.MessengerService/RegisterUser',
  grpc.web.MethodType.UNARY,
  proto.messenger.User,
  proto.messenger.User,
  /**
   * @param {!proto.messenger.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.messenger.User.deserializeBinary
);


/**
 * @param {!proto.messenger.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.messenger.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.messenger.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServiceClient.prototype.registerUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/messenger.MessengerService/RegisterUser',
      request,
      metadata || {},
      methodDescriptor_MessengerService_RegisterUser,
      callback);
};


/**
 * @param {!proto.messenger.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.messenger.User>}
 *     Promise that resolves to the response
 */
proto.messenger.MessengerServicePromiseClient.prototype.registerUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/messenger.MessengerService/RegisterUser',
      request,
      metadata || {},
      methodDescriptor_MessengerService_RegisterUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.messenger.FriendRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_MessengerService_SendFriendRequest = new grpc.web.MethodDescriptor(
  '/messenger.MessengerService/SendFriendRequest',
  grpc.web.MethodType.UNARY,
  proto.messenger.FriendRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.messenger.FriendRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.messenger.FriendRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServiceClient.prototype.sendFriendRequest =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/messenger.MessengerService/SendFriendRequest',
      request,
      metadata || {},
      methodDescriptor_MessengerService_SendFriendRequest,
      callback);
};


/**
 * @param {!proto.messenger.FriendRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.messenger.MessengerServicePromiseClient.prototype.sendFriendRequest =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/messenger.MessengerService/SendFriendRequest',
      request,
      metadata || {},
      methodDescriptor_MessengerService_SendFriendRequest);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.messenger.Chat,
 *   !proto.messenger.Chat>}
 */
const methodDescriptor_MessengerService_CreateChat = new grpc.web.MethodDescriptor(
  '/messenger.MessengerService/CreateChat',
  grpc.web.MethodType.UNARY,
  proto.messenger.Chat,
  proto.messenger.Chat,
  /**
   * @param {!proto.messenger.Chat} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.messenger.Chat.deserializeBinary
);


/**
 * @param {!proto.messenger.Chat} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.messenger.Chat)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.messenger.Chat>|undefined}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServiceClient.prototype.createChat =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/messenger.MessengerService/CreateChat',
      request,
      metadata || {},
      methodDescriptor_MessengerService_CreateChat,
      callback);
};


/**
 * @param {!proto.messenger.Chat} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.messenger.Chat>}
 *     Promise that resolves to the response
 */
proto.messenger.MessengerServicePromiseClient.prototype.createChat =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/messenger.MessengerService/CreateChat',
      request,
      metadata || {},
      methodDescriptor_MessengerService_CreateChat);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.messenger.Message,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_MessengerService_SendMessage = new grpc.web.MethodDescriptor(
  '/messenger.MessengerService/SendMessage',
  grpc.web.MethodType.UNARY,
  proto.messenger.Message,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.messenger.Message} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.messenger.Message} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServiceClient.prototype.sendMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/messenger.MessengerService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_MessengerService_SendMessage,
      callback);
};


/**
 * @param {!proto.messenger.Message} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.messenger.MessengerServicePromiseClient.prototype.sendMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/messenger.MessengerService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_MessengerService_SendMessage);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.messenger.GetMessagesRequest,
 *   !proto.messenger.Message>}
 */
const methodDescriptor_MessengerService_GetMessages = new grpc.web.MethodDescriptor(
  '/messenger.MessengerService/GetMessages',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.messenger.GetMessagesRequest,
  proto.messenger.Message,
  /**
   * @param {!proto.messenger.GetMessagesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.messenger.Message.deserializeBinary
);


/**
 * @param {!proto.messenger.GetMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.messenger.Message>}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServiceClient.prototype.getMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/messenger.MessengerService/GetMessages',
      request,
      metadata || {},
      methodDescriptor_MessengerService_GetMessages);
};


/**
 * @param {!proto.messenger.GetMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.messenger.Message>}
 *     The XHR Node Readable Stream
 */
proto.messenger.MessengerServicePromiseClient.prototype.getMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/messenger.MessengerService/GetMessages',
      request,
      metadata || {},
      methodDescriptor_MessengerService_GetMessages);
};


module.exports = proto.messenger;
