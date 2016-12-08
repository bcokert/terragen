"use strict";

const TextureLoader = require("./loaders/texture-loader");

const ResourceLoaderProto = {
    /**
     * Loads all resources configured in the resource loaders. Returns a Promise that results in a map from resource type to the loaded resources,
     * in the format each individual loader returns them
     * @param {WebGLRenderingContext} gl
     * @returns {Promise}
     */
    loadResources(gl) {
        const resourceTypes = Object.keys(this.resourceConfig);
        return new Promise((resolve) => {
            Promise.all(resourceTypes.map(type => this.resourceConfig[type].loader.loadResources(gl, this.resourceConfig[type].resources))).then(resources => {
                resolve(resources.reduce((acc, next, i) => {
                    acc[resourceTypes[i]] = next;
                    return acc;
                }, {}));
            });
        });
    }
};

/**
 * A ResourceLoader loads resources of various kinds by delegating groups of resources to component loaders, like TextureLoader.
 * @param {{textures: object?}} loaders A map from resourceType to the appropriate loader for it. If any loader is not provided, a default is created.
 * @returns {ResourceLoader}
 * @constructor
 */
const ResourceLoader = (loaders) => {
    loaders = loaders || {};
    const resourceConfig = {
        textures: {
            loader: loaders.textures || TextureLoader(),
            resources: [{
                name: "testWoodTexture",
                url: require("../../../img/textures/testWoodTexture.jpg"),
                invertY: true
            }]
        }
    };

    return Object.assign(Object.create(ResourceLoaderProto), {
        resourceConfig: resourceConfig
    });
};

ResourceLoader.prototype = ResourceLoaderProto;

module.exports = ResourceLoader;
