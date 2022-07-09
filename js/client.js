// Path to the proto file defination
const PROTO_PATH = __dirname + '/../proto/Feed.proto';

// Imports
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

// Load the proto file and its defination
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true, longs: String, enums: String, defaults: true, oneofs: true}
);
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const feed_proto = protoDescriptor.feed;

// Target server to hit. This is better to be environment variable.
const target = "localhost:50051";

/**
 * 
 * @param {any} client 
 * 
 * @param {JSON} feedData {
 *    feed_id: number,
 *  }
 * 
 * @description get feeds by id from a gRPC server
 * 
 * @returns {void}
 */
const getFeed = (client, feedData) => {
    client.getFeed(feedData, (err, response) => {
        if (err) {
            console.log(err.message);
            return;
        }
        console.log(response);
    })
}

/**
 * 
 * @param {any} client 
 * 
 * @description get all feeds
 * 
 * @returns {void}
 */
const getAllFeeds = (client) => {
    client.getAllFeeds({}, (err, response) => {
        if (err) {
            console.log(err.message);
            return;
        }
        console.log(response);
    })
}

/**
 * 
 * @param {*} client 
 * @param {JSON} feedData {
 *   feed_title: string,
 *   feed_content: string,
 *   feed_author: string,
 * }
 * 
 * @description post a feed
 * 
 * @returns {void}
 */
const postFeed = (client, feedData) => {
    client.postFeed(feedData, (err, response) => {
        if (err) {
            console.log(err.message);
            return;
        }
        console.log(response);
    })
}

/**
 * 
 * @param {*} client 
 * @param {JSON} feedData {
 *   feed_id: number,
 * }
 * 
 * @description delete a feed by id
 * 
 * @returns {void}
 */
const deleteFeed = (client, feedData) => {
    client.deleteFeed(feedData, (err, response) => {
        if (err) {
            console.log(err.message);
            return;
        }
        console.log(response);
    })
}

function main() {
    const client = new feed_proto.FeedService(target, grpc.credentials.createInsecure());

    /**
     * @description get a feed by id
     */
    // getFeed(client, {
    //     feed_id: 2
    // });

    /**
     * @description get all feeds
     */
    getAllFeeds(client);

    /**
     * @description delete a feed by id
     * @param {JSON} feedData {
     *   feed_id: number,
     * }
     */
    // deleteFeed(client, {
    //     feed_id: 2
    // });

    /**
     * @description post a feed
     * @param {JSON} feedData {
     *   feed_title: string,
     *   feed_content: string,
     *   feed_author: string,
     * }
     */ 
    // postFeed(client, {
    //     feed_title: "test",
    //     feed_content: "test",
    //     feed_author: "test"
    // });
};

main();