// Imports
const { Sequelize, DataTypes } = require('sequelize');

// Creating sequelize instance
const sequelize = new Sequelize({
    dialect: "sqlite",
    storage: `${__dirname}/../database.sqlite`,
    logging: false
});

// feed model to store in database
const FeedModel = sequelize.define('feed', {
    id: {
        type: DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true
    },
    title: {
        type: DataTypes.STRING,
        allowNull: false
    },
    content: {
        type: DataTypes.STRING,
        allowNull: false
    },
    author: {
        type: DataTypes.STRING,
        allowNull: false
    },
    createdAt: {
        type: DataTypes.DATE,
        allowNull: false,
        timestamps: true
    }
});

// function to create the feeds table with required columns
const FeedUp = () => {
    const queryInterface = sequelize.getQueryInterface();
    queryInterface.createTable('feeds', {
        id: {
            type: DataTypes.INTEGER,
            primaryKey: true,
            autoIncrement: true
        },
        title: {
            type: DataTypes.STRING,
            allowNull: false
        },
        content: {
            type: DataTypes.STRING,
            allowNull: false
        },
        author: {
            type: DataTypes.STRING,
            allowNull: false
        },
        createdAt: {
            type: DataTypes.DATE,
            allowNull: false
        },
        updatedAt: {
            type: DataTypes.DATE,
            allowNull: false
        }
    })
};

// function to delete the feeds table
const FeedDown = () => {
    const queryInterface = sequelize.getQueryInterface();
    queryInterface.dropTable('feeds');
};

module.exports = {
    FeedModel,
    FeedUp,
    FeedDown
};