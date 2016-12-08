"use strict";

const RenderComponent = require("../render-component");
const GLMatrix = require("gl-matrix");

const MovableProto = {
    /**
     * Moves the object by the given vec3
     * @param {Float32Array} dVector
     * @returns {undefined}
     */
    move(dVector) {
        GLMatrix.vec3.add(this.position, this.position, dVector);
    }
};

/**
 * Adds the ability for an object to be movable.
 * @param {number} maxSpeed - The max speed for the object
 * @constructor
 * @returns {Object}
 */
const Movable = maxSpeed => RenderComponent({
    name: "Movable",
    validator: instance => {
        if (!instance.position) {
            throw new TypeError("The Movable update component depends on a 'position' property.");
        }
    },
    proto: MovableProto,
    properties: {
        velocity: GLMatrix.vec3.create([0, 0, 0]),
        maxSpeed: maxSpeed,
        acceleration: GLMatrix.vec3.create([0, 0, 0])
    },
    updateFn: (instance, dtSeconds) => {
        for (let i = 0; i < instance.position.length; i++) {
            instance.velocity[i] += instance.acceleration[i] * dtSeconds;
            if (instance.velocity[i] > instance.maxSpeed) {
                instance.velocity[i] = instance.maxSpeed;
            } else if (instance.velocity[i] < -instance.maxSpeed) {
                instance.velocity[i] = -instance.maxSpeed;
            }

            instance.position[i] += instance.velocity[i] * dtSeconds;
        }
    }
});

module.exports = Movable;
