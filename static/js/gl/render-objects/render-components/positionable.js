"use strict";

const RenderComponent = require("../render-component");
const GLMatrix = require("gl-matrix");

/**
 * Adds the ability for an object to be positionable.
 * @param {Float32Array} position
 * @constructor
 * @returns {Object}
 */
const Positionable = (position) => {
    if (Object.prototype.toString.call(position) !== "[object Float32Array]" || position.length !== 3) {
        throw new TypeError("Invalid position provided to Positionable factory.");
    }

    return RenderComponent({
        name: "Positionable",
        properties: {
            position: position
        },
        renderFn: (instance, gl, programInfo, modelViewMatrix) => {
            GLMatrix.mat4.translate(modelViewMatrix, modelViewMatrix, instance.position);
            return modelViewMatrix;
        }
    });
};

module.exports = Positionable;
