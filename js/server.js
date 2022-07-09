// Path to the proto file defination
const PROTO_PATH = __dirname + '/../proto/Feed.proto';

// Imports
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { FeedModel } = require('./db');

// Loading the proto file and its defination
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true, longs: String, enums: String, defaults: true, oneofs: true}
);
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const feed_proto = protoDescriptor.feed;

// function to create the feed structure to match the proto definition
const createFeed = (feedData) => {
    return {
        feed_id: feedData.id,
        feed_title: feedData.title,
        feed_content: feedData.content,
        feed_publish_time: feedData.createdAt,
        feed_author: feedData.author
    }
}

// get feed function to send only one feed depending on the id
const getFeed = (call, callback) => {
    const feed = FeedModel.findOne({
        where: {
            id: call.request.feed_id
        }
    })
    feed.then(feedData => {
        if(!feedData) {
            callback({
                code: 400,
                message: "invalid input",
                status: grpc.status.INTERNAL
            }, null);
            return;
        }
        callback(null, createFeed(feedData));
    })
}

// function to send all feeds
const getAllFeeds = (call, callback) => {
    const feeds = FeedModel.findAll({});
    feeds.then(feedData => {
        if(!feedData) {
            callback({
                code: 400,
                message: "invalid input",
                status: grpc.status.INTERNAL
            }, null);
            return;
        }
        callback(null, {
            feeds: feedData.map(feed => createFeed(feed))
        });
    })
}

// function to save a new feed
const postFeed = (call, callback) => {
    const {
        feed_title,
        feed_content,
        feed_author
    } = call.request;

    if (!feed_title || !feed_content || !feed_author) {
        callback({
            code: 400,
            message: `invalid input you need to input feed_title, feed_content and feed_author`,
            status: grpc.status.INTERNAL
        }, null);
        return;
    }

    const feed = new FeedModel({
        title: feed_title,
        content: feed_content,
        author: feed_author,
        createdAt: new Date()
    });

    feed.save()
        .then((feedData) => {
            callback(null, createFeed(feedData));
        });
}

// function to delete a feed by id
const deleteFeed = (call, callback) => {
    const feed = FeedModel.destroy({
        where: {
            id: call.request.feed_id
        }
    })
    feed.then(feedData => {
        if(!feedData) {
            callback({
                code: 400,
                message: "invalid input",
                status: grpc.status.INTERNAL
            }, null);
            return;
        }
        callback(null, {
            message: "success"
        });
    })
}

// main server function
function main() {
    // Create a new gRPC server
    var server = new grpc.Server();

    // adding the services to the server
    server.addService(feed_proto.FeedService.service, {
        getFeed: getFeed,
        getAllFeeds: getAllFeeds,
        postFeed: postFeed,
        deleteFeed: deleteFeed
    });

    // Bind to port
    server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      server.start();
      console.log('Server running at http://0.0.0.0:50051');
    });
}

main();