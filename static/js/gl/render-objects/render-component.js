"use strict";

/**
 * Creates a standard RenderComponent, which is something that a RenderObject can be composed of
 * @param {Object} config
 * @param {string} config.name - The unique name for this component. Available in the final object in object.renderComponents
 * @param {Function?} config.validator - A function that will be passed the instance before composing, to ensure it is ready to be composed. It should throw an exception on failure.
 * @param {Function?} config.initializer - A function that will be passed the instance after composing, to run initialization code
 * @param {Object?} config.properties - Properties to add to the new instance
 * @param {Object?} config.proto - Properties to add to the new instances prototype
 * @param {Function?} config.updateFn - The update function that will be appended to this components updateQueue
 * @param {Function?} config.renderFn - The render function that will be appended to this components renderQueue
 * @param {Function?} config.finishRenderFn - A function that will be called once all other render functions are called. At least one renderComponent must provide this for every renderObject
 * @returns {Object}
 */
const RenderComponent = (config) => {
    config.updateFn = config.updateFn || null;
    config.renderFn = config.renderFn || null;
    config.proto = config.proto || null;
    config.properties = config.properties || null;

    const configParams = ["updateFn", "renderFn", "proto", "properties", "name", "validator", "initializer", "finishRenderFn"];
    const illegalParams = Object.keys(config).filter(v => !configParams.includes(v));
    if (illegalParams.length > 0) {
        throw new TypeError("Invalid config params given to RenderComponent: '" + illegalParams.join(", ") + "'. Accepted config params are: " + configParams.join(", "));
    }

    if (!config.name || typeof config.name !== "string" || config.name.length < 6) {
        throw new TypeError("RenderComponent was given an invalid name. It must be a string of at least 6 characters.");
    }
    if (config.validator && typeof config.validator !== "function") {
        throw new TypeError("validator in RenderComponent must be a function or null");
    }
    if (config.initializer && typeof config.initializer !== "function") {
        throw new TypeError("initializer in RenderComponent must be a function or null");
    }
    if (config.updateFn && typeof config.updateFn !== "function") {
        throw new TypeError("updateFn in RenderComponent must be a function or null");
    }
    if (config.renderFn && typeof config.renderFn !== "function") {
        throw new TypeError("renderFn in RenderComponent must be a function or null");
    }
    if (config.finishRenderFn && typeof config.finishRenderFn !== "function") {
        throw new TypeError("finishRenderFn in RenderComponent must be a function or null");
    }
    if (typeof config.proto !== "object") {
        throw new TypeError("proto in RenderComponent must be an object or null");
    }
    if (typeof config.properties !== "object") {
        throw new TypeError("properties in RenderComponent must be an object or null");
    }

    return Object.create({
        /**
         * Mutates the given RenderObject by making it composed of this RenderComponent
         * @param {Object} renderObject A RenderObject
         * @returns {Object} The same renderObject reference it was given
         */
        compose: renderObject => {
            if (!renderObject || typeof renderObject !== "object") {
                throw new TypeError("Attempted to compose with a null RenderObject");
            }
            if (!Array.isArray(renderObject.renderQueue)) {
                throw new TypeError("The given RenderObject does not have a renderQueue array");
            }
            if (!Array.isArray(renderObject.updateQueue)) {
                throw new TypeError("The given RenderObject does not have an updateQueue array");
            }
            if (!Array.isArray(renderObject.renderComponents)) {
                throw new TypeError("The given RenderObject does not have a renderComponents array");
            }

            if (config.validator) {
                config.validator(renderObject);
            }
            
            renderObject.renderComponents.push(config.name);

            if (config.proto) {
                renderObject.__proto__ = Object.assign(renderObject.__proto__, config.proto);
            }

            if (config.properties) {
                renderObject = Object.assign(renderObject, config.properties);
            }

            if (config.updateFn) {
                renderObject.updateQueue.push(config.updateFn);
            }

            if (config.renderFn) {
                renderObject.renderQueue.push(config.renderFn);
            }

            if (config.finishRenderFn) {
                renderObject.finishRender = config.finishRenderFn;
            }

            if (config.initializer) {
                config.initializer(renderObject);
            }

            return renderObject;
        }
    });
};

module.exports = RenderComponent;
